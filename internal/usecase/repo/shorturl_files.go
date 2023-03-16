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
)

const (
	secretSecret = "RtsynerpoGIYdab_s234r"
	cookieName   = "access_token"
)

type (
	producer struct {
		file    *os.File
		encoder *json.Encoder
	}

	consumer struct {
		file    *os.File
		cfg     *config.Config
		decoder *json.Decoder
		//decoder *bufio.Reader
	}

	InFiles struct {
		//lock sync.Mutex // <-- этот мьютекс защищает
		m   map[string]entity.Shorturls
		cfg *config.Config
		r   *consumer
		w   *producer
	}
)

// NewInFiles слой взаимодействия с файловым хранилищем
func NewInFiles(cfg *config.Config) *InFiles {
	return &InFiles{
		cfg: cfg,
		m:   make(map[string]entity.Shorturls),
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
	}
}

func (i *InFiles) post(sh *entity.Shorturl) error {
	i.m[sh.UserID] = append(i.m[sh.UserID], *sh)
	return nil
}
func (i *InFiles) Post(ctx context.Context, sh *entity.Shorturl) error {
	//i.lock.Lock()
	//defer i.lock.Unlock()
	return i.post(sh)
}

func (i *InFiles) Put(ctx context.Context, sh *entity.Shorturl) error {
	ln := len(i.m[sh.UserID])
	if ln < 1 {
		i.Post(ctx, sh)
		return nil
	}
	for j := 0; j < ln; j++ {
		if i.m[sh.UserID][j].Slug == sh.Slug {
			i.m[sh.UserID][j].URL = sh.URL
		}
	}
	return i.Post(ctx, sh)
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

func (i *InFiles) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	return i.searchBySlug(sh)
}

func (i *InFiles) searchUID(sh *entity.Shorturl) (*entity.Shorturl, error) {
	for _, short := range i.m[sh.UserID] {
		if short.Slug == sh.Slug {
			sh.URL = short.URL
			sh.UserID = short.UserID
			sh.Del = short.Del
			fmt.Println("НАШЁЛ URL: ", sh.URL)
			break
		}
	}
	return sh, nil
}

// search by slug
func (i *InFiles) searchBySlug(sh *entity.Shorturl) (*entity.Shorturl, error) {
	shorts := entity.Shorturls{}
	for _, uid := range i.m {
		for j := 0; j < len(uid); j++ {
			shorts = append(shorts, uid[j])
		}
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
func (i *InFiles) getAll() error {
	sh := &entity.Shorturl{}
	for {
		if err := i.r.decoder.Decode(&sh); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		i.m[sh.UserID] = append(i.m[sh.UserID], *sh)
	}
	return nil
}
func (i *InFiles) getAllUserID(u *entity.User) (*entity.User, error) {
	lst := entity.List{}
	shorts := i.m[u.UserID]
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
	//i.lock.Lock()
	//defer i.lock.Unlock()
	return i.getAllUserID(u)
}

// Save перезаписать файл с новыми данными
func (i *InFiles) Read() error {
	for {
		if err := i.r.decoder.Decode(&i.m); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	//fmt.Printf("Read data for file: %v", i.m)
	return nil
}

// Save перезаписать файл с новыми данными
func (i *InFiles) Save() error {
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
	// пишем из памяти в файл
	return i.w.encoder.Encode(i.m)
}
func (i *InFiles) Delete(ctx context.Context, u *entity.User) error {
	//i.lock.Lock()
	//defer i.lock.Unlock()
	return i.delete(u)
}
func (i *InFiles) delete(u *entity.User) error {
	for j := 0; j < len(i.m[u.UserID]); j++ {
		for _, slug := range u.DelLink {
			if i.m[u.UserID][j].Slug == slug {
				// изменяет флаг del на true, в результате url становиться недоступным для пользователя
				i.m[u.UserID][j].Del = true
			}
		}
	}
	return nil
}
func (c *consumer) Close() error {
	return c.file.Close()
}
