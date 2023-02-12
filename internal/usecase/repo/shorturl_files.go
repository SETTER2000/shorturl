package repo

import (
	"bufio"
	"encoding/json"
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

	InFiles struct {
		r *consumer
		w *producer
	}
)

// NewInFiles слой взаимодействия с файловым хранилищем
func NewInFiles(file *os.File) *InFiles {
	return &InFiles{
		// создаём новый потребитель
		r: NewConsumer(file),
		// создаём новый производитель
		w: NewProducer(file),
	}
}

// NewConsumer потребитель
func NewConsumer(file *os.File) *consumer {
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

	i.r.file.Seek(0, 0)
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

//func (p *producer) Close() error {
//	return p.file.Close()
//}
