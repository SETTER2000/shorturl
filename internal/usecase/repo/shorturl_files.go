package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/entity"
	"os"
)

type (
	Producer struct {
		file   *os.File
		writer *bufio.Writer
	}

	Consumer struct {
		file *os.File
		// заменяем reader на scanner
		reader *bufio.Reader
	}

	//InFiles struct {
	//	cns *Consumer
	//	prd *Producer
	//}
	InFiles struct {
		file *os.File
		// заменяем reader на scanner
		reader *bufio.Reader // consumer
		writer *bufio.Writer // producer
	}
)

func NewInFiles(file *os.File) *InFiles {
	return &InFiles{
		file: file,
		// создаём новый scanner
		reader: bufio.NewReader(file),
		writer: bufio.NewWriter(file),
	}
}

// NewConsumer потребитель
func NewConsumer(file *os.File) *Consumer {
	return &Consumer{
		file: file,
		// создаём новый scanner
		reader: bufio.NewReader(file),
	}
}

type Shorturl struct {
	Slug string `json:"slug" example:"1674872720465761244B_5"`
	URL  string `json:"url" example:"https://example.com/go/to/home.html"`
}

func (i *InFiles) Get(key string) (*entity.Shorturl, error) {
	sh := entity.Shorturl{}
	fmt.Println(i.reader.Size())
	if i.reader.Size() < 1 {
		return nil, ErrNotFound
	}
	for {
		data, err := i.reader.ReadBytes('\n')
		if err != nil {
			return nil, ErrNotFound
		}

		err = json.Unmarshal(data, &sh)
		if err != nil {
			i.file.Seek(0, 0)
		}

		if sh.Slug == key {
			break
		}
	}

	i.file.Seek(0, 0)
	return &sh, nil
}

func (c *Consumer) Close() error {
	return c.file.Close()
}

// NewProducer производитель
func NewProducer(file *os.File) *Producer {
	return &Producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}
}

func (i *InFiles) Post(sh *entity.Shorturl) error {
	data, err := json.Marshal(sh)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err = i.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err = i.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	t := i.writer.Flush()
	return t
}

func (i *InFiles) Put(sh *entity.Shorturl) error {
	return i.Post(sh)
}

func (p *Producer) Close() error {
	return p.file.Close()
}
