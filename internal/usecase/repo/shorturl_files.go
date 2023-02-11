package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/entity"
	"os"
)

type (
	producer struct {
		file   *os.File
		writer *bufio.Writer
	}

	consumer struct {
		file *os.File
		// заменяем reader на scanner
		reader *bufio.Reader
	}
)

// NewConsumer потребитель
func NewConsumer(file *os.File) *consumer {
	//if cfg.FileStorage == " " {
	//	return nil
	//}
	//file, _ := os.OpenFile(cfg.FileStorage, os.O_RDONLY|os.O_CREATE, 0777)

	return &consumer{
		file: file,
		// создаём новый scanner
		reader: bufio.NewReader(file),
	}
}

type Shorturl struct {
	Slug string `json:"slug" example:"1674872720465761244B_5"`
	URL  string `json:"url" example:"https://example.com/go/to/home.html"`
}

func (c *consumer) Get(key string) (*entity.Shorturl, error) {
	sh := entity.Shorturl{}
	fmt.Println(c.reader.Size())
	if c.reader.Size() < 1 {
		return nil, ErrNotFound
	}
	for {
		data, err := c.reader.ReadBytes('\n')
		if err != nil {
			return nil, ErrNotFound
		}

		err = json.Unmarshal(data, &sh)
		if err != nil {
			c.file.Seek(0, 0)
		}

		if sh.Slug == key {
			break
		}
	}

	c.file.Seek(0, 0)
	return &sh, nil
}

func (c *consumer) Close() error {
	return c.file.Close()
}

// NewProducer производитель
func NewProducer(file *os.File) *producer {
	return &producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}
}

func (p *producer) Post(sh *entity.Shorturl) error {
	data, err := json.Marshal(&sh)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return p.writer.Flush()
}

func (p *producer) Put(sh *entity.Shorturl) error {
	return p.Post(sh)
}

func (p *producer) Close() error {
	return p.file.Close()
}
