package repo

import (
	"bufio"
	"encoding/json"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
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
	//if err != nil {
	//	return nil, err
	//}
	return &producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}
}

func (i *InFiles) Post(sh *entity.Shorturl) error {
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

func (i *InFiles) Put(sh *entity.Shorturl) error {
	return i.Post(sh)
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

func (i *InFiles) Get(key string) (*entity.Shorturl, error) {
	sh := entity.Shorturl{}
	if i.r.reader.Size() < 1 {
		return nil, ErrNotFound
	}
	for {
		data, err := i.r.reader.ReadBytes('\n')
		if err != nil {
			return nil, ErrNotFound
		}

		err = json.Unmarshal(data, &sh)
		if err != nil {
			i.r.file.Seek(0, 0)
		}

		if sh.Slug == key {
			break
		}
	}
	_, err := i.r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return &sh, nil
}

func (i *InFiles) GetAll(u *entity.User) (*entity.User, error) {
	sh := entity.Shorturl{}
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
		if sh.UserId == u.UserId {
			sh.UserId = ""
			sh.Slug = scripts.GetHost(i.cfg.HTTP, sh.Slug)
			u.Urls = append(u.Urls, sh)
		}
	}
	_, err := i.r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *consumer) Close() error {
	return c.file.Close()
}
