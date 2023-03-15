package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"io"
	"os"
	"sync"
)

const (
	secretSecret = "RtsynerpoGIYdab_s234r"
	cookieName   = "access_token"
)

type (
	producer struct {
		lock sync.Mutex // <-- этот мьютекс защищает
		file *os.File
		//encoder *bufio.Writer
		encoder *json.Encoder
	}

	consumer struct {
		lock    sync.Mutex // <-- этот мьютекс защищает
		file    *os.File
		cfg     *config.Config
		decoder *json.Decoder
		//decoder *bufio.Reader
	}

	InFiles struct {
		cfg *config.Config
		m   map[string][]byte
		r   *consumer
		w   *producer
	}
)

// NewInFiles слой взаимодействия с файловым хранилищем
func NewInFiles(cfg *config.Config) *InFiles {
	return &InFiles{
		cfg: cfg,
		m:   make(map[string][]byte),
		// создаём новый потребитель
		r: NewConsumer(cfg),
		// создаём новый производитель
		w: NewProducer(cfg),
	}
}

// NewProducer производитель
func NewProducer(cfg *config.Config) *producer {
	file, _ := os.OpenFile(cfg.FileStorage, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
		//encoder: bufio.NewWriter(file),
	}
}

func (i *InFiles) post(sh *entity.Shorturl) error {
	return i.w.encoder.Encode(&sh)
}
func (i *InFiles) Post(ctx context.Context, sh *entity.Shorturl) error {
	i.w.lock.Lock()
	defer i.w.lock.Unlock()
	return i.post(sh)
}

func (i *InFiles) Put(ctx context.Context, sh *entity.Shorturl) error {
	i.w.lock.Lock()
	defer i.w.lock.Unlock()
	return i.post(sh)
}

func (p *producer) Close() error {
	return p.file.Close()
}

// NewConsumer потребитель
func NewConsumer(cfg *config.Config) *consumer {
	file, _ := os.OpenFile(cfg.FileStorage, os.O_RDONLY|os.O_CREATE, 0777)
	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}
}

func (i *InFiles) Get2(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	//i.r.lock.Lock()
	//defer i.r.lock.Unlock()
	sh2, err := i.getSlag(ctx, sh)
	if err != nil {
		return nil, err
	}
	return sh2, nil
}
func (i *InFiles) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	short := entity.Shorturl{}
	var shorts entity.Shorturls
	//var bufferRead bytes.Buffer
	defer i.r.file.Seek(0, 0)

	for {
		if err := i.r.decoder.Decode(&short); err != nil {
			//io.TeeReader(i.r.decoder.Buffered(), &bufferRead)
			//fmt.Printf("\nBufferRead: %s\n", &bufferRead)
			//n, _ := i.r.file.Seek(0, 0)
			//fmt.Printf("NUM START : %v\n", n)
			//fmt.Printf("End FILE : %v\n", err)
			if err == io.EOF {
				//i.r.decoder.Buffered()
				break
			}
			return nil, err
		}
		fmt.Printf("APPEND : %v\n", short.URL)

		shorts = append(shorts, short)
	}

	for _, s := range shorts {
		if s.Slug == sh.Slug {
			sh.URL = s.URL
			sh.UserID = s.UserID
			sh.Del = s.Del
			break
		}
	}
	i.r.file.Seek(0, io.SeekStart)
	return sh, nil
}

func (i *InFiles) getSlag(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	shorts, err := i.getAll()
	if err != nil {
		return nil, err
	}
	for _, short := range shorts {
		if short.Slug == sh.Slug {
			sh.URL = short.URL
			sh.UserID = short.UserID
			sh.Del = short.Del
			break
		}
	}
	return sh, nil
}

func (i *InFiles) getAll() ([]entity.Shorturl, error) {
	sh := &entity.Shorturl{}
	var shorts []entity.Shorturl
	for {
		if err := i.r.decoder.Decode(&sh); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		shorts = append(shorts, *sh)
	}
	return shorts, nil
}
func (i *InFiles) getAllUserID(u *entity.User) (*entity.User, error) {
	lst := entity.List{}
	shorts, err := i.getAll()
	if err != nil {
		return nil, err
	}
	for _, short := range shorts {
		if short.UserID == u.UserID {
			lst.URL = short.URL
			lst.Slug = scripts.GetHost(i.cfg.HTTP, short.Slug)
			u.Urls = append(u.Urls, lst)
		}
	}
	return u, nil
}
func (i *InFiles) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	i.r.lock.Lock()
	defer i.r.lock.Unlock()
	return i.getAllUserID(u)
}

func (i *InFiles) delete(shorts []entity.Shorturl, u *entity.User) ([]entity.Shorturl, error) {
	var d []entity.Shorturl
	for _, v := range shorts {
		for _, g := range u.DelLink {
			if v.Slug == g && v.UserID == u.UserID {
				// изменяет флаг del на true, в результате url становиться недоступным для пользователя
				v.Del = true
			}
		}
		// обновлённый слайс данных, с флагом del=true
		d = append(d, v)
	}
	return d, nil
}

// rewriteFile перезаписать файл с новыми данными
func (i *InFiles) rewriteFile(shorts []entity.Shorturl) error {
	//переводит курсор в начало файла
	_, err := i.w.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	// очищает файл
	err = i.w.file.Truncate(0)
	if err != nil {
		return err
	}
	for _, sh := range shorts {
		err = i.post(&sh)
		if err != nil {
			return err
		}
	}
	i.r.file.Seek(0, 0)

	return nil
}
func (i *InFiles) Delete(ctx context.Context, u *entity.User) error {
	//i.w.lock.Lock()
	shorts, err := i.getAll()
	//i.w.lock.Unlock()
	if err != nil {
		return err
	}
	//// изменяет флаг del на true, в результате url становиться недоступным для пользователя
	shorts, _ = i.delete(shorts, u)
	// перезаписать файл с новыми значениями
	err = i.rewriteFile(shorts)
	return err
}
func (c *consumer) Close() error {
	return c.file.Close()
}
