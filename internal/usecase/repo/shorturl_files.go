package repo

import (
	"bufio"
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
		file   *os.File
		writer *bufio.Writer
	}

	consumer struct {
		cfg  *config.Config
		file *os.File
		// заменяем reader на scanner
		reader *bufio.Reader
	}

	InFiles struct {
		cfg *config.Config
		r   *consumer
		w   *producer
	}
)

// NewInFiles слой взаимодействия с файловым хранилищем
func NewInFiles(cfg *config.Config) *InFiles {
	return &InFiles{
		cfg: cfg,
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
		file:   file,
		writer: bufio.NewWriter(file),
	}
}

func (i *InFiles) Post(ctx context.Context, sh *entity.Shorturl) error {
	data, err := json.Marshal(&sh)
	if err != nil {
		return err
	}
	// записываем событие в буфер
	if _, err = i.w.writer.Write(data); err != nil {
		return err
	}
	// добавляем перенос строки
	if err = i.w.writer.WriteByte('\n'); err != nil {
		return err
	}
	// записываем буфер в файл
	t := i.w.writer.Flush()
	return t
}

func (i *InFiles) Put(ctx context.Context, sh *entity.Shorturl) error {
	var shorts []entity.Shorturl
	var sh2 entity.Shorturl
	for {
		path, err := i.r.reader.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			break
		} else if err != nil {
			return err // if you return error
		}
		err = json.Unmarshal([]byte(path), &sh2)
		if err != nil {
			i.r.file.Seek(0, 0)
		}
		if sh2.Slug != sh.Slug {
			shorts = append(shorts, sh2)
		}

	}
	shorts = append(shorts, *sh)

	_, err := i.w.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	err = i.w.file.Truncate(0)
	if err != nil {
		return err
	}

	for _, short := range shorts {

		data, err := json.Marshal(&short)
		if err != nil {
			return err
		}
		// записываем событие в буфер
		if _, err = i.w.writer.Write(data); err != nil {
			return err
		}
		// добавляем перенос строки
		if err = i.w.writer.WriteByte('\n'); err != nil {
			return err
		}

	}

	fmt.Printf("TTt:: %v", shorts)
	fmt.Printf("Count: %d\n", len(shorts))
	i.r.file.Seek(0, 0)
	// записываем буфер в файл
	t := i.w.writer.Flush()
	return t
}

func (p *producer) Close() error {
	return p.file.Close()
}

// NewConsumer потребитель
func NewConsumer(cfg *config.Config) *consumer {
	file, _ := os.OpenFile(cfg.FileStorage, os.O_RDONLY|os.O_CREATE, 0777)
	return &consumer{
		file: file,
		// создаём новый scanner
		reader: bufio.NewReader(file),
	}
}

func (i *InFiles) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	sh2 := entity.Shorturl{}
	if i.r.reader.Size() < 1 {
		return nil, ErrNotFound
	}
	for {
		data, err := i.r.reader.ReadBytes('\n')
		if err != nil {
			return nil, ErrNotFound
		}
		err = json.Unmarshal(data, &sh2)
		if err != nil {
			i.r.file.Seek(0, 0)
		}
		if sh2.Slug == sh.Slug {
			sh.URL = sh2.URL
			sh.UserID = sh2.UserID
			break
		}
	}
	_, err := i.r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return sh, nil
}

func (i *InFiles) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	sh := entity.Shorturl{}
	lst := entity.List{}
	size := i.r.reader.Size()
	if size < 1 {
		return nil, ErrNotFound
	}
	for j := 0; j < size; j++ {
		data, err := i.r.reader.ReadBytes('\n')
		if err != nil {
			break
		}
		err = json.Unmarshal(data, &sh)
		if err != nil {
			i.r.file.Seek(0, 0)
		}
		if sh.UserID == u.UserID {
			lst.URL = sh.URL
			lst.Slug = scripts.GetHost(i.cfg.HTTP, sh.Slug)
			u.Urls = append(u.Urls, lst)
		}
	}
	_, err := i.r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (i *InFiles) Delete(ctx context.Context, u *entity.User) error {
	//user := entity.User{}
	return nil
}
func (c *consumer) Close() error {
	return c.file.Close()
}
