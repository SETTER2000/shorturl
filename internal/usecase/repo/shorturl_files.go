package repo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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
		scanner *bufio.Scanner
	}
)

// NewConsumer потребитель
func NewConsumer() *consumer {
	link := fmt.Sprintf("%s", os.Getenv("FILE_STORAGE_PATH"))
	file, err := os.OpenFile(link, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return &consumer{
		file: file,
		// создаём новый scanner
		scanner: bufio.NewScanner(file),
	}
}

type Event struct{}

func (c *consumer) Get(key string) (string, error) {
	// одиночное сканирование до следующей строки
	if !c.scanner.Scan() {
		return "", c.scanner.Err()
	}
	// читаем данные из scanner
	data := c.scanner.Bytes()

	event := Event{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return "", err
	}

	fmt.Sprintf("%v", event)

	return "", nil
}

func (c *consumer) Close() error {
	return c.file.Close()
}

// NewProducer производитель
func NewProducer() *producer {

	// путь к файлу
	link := fmt.Sprintf("%s", os.Getenv("FILE_STORAGE_PATH"))

	file, err := os.OpenFile(link, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return &producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}
}

func (p *producer) Get(key string) (string, error) {
	//s.lock.Lock()
	//defer s.lock.Unlock()
	//
	//if v, ok := s.m[key]; ok {
	//	return v, nil
	//}
	return "", ErrNotFound
}

func (p *producer) Put(key string, value string) error {
	//s.lock.Lock()
	//defer s.lock.Unlock()
	//
	//if _, ok := s.m[key]; ok {
	//	return ErrAlreadyExists
	//}
	//s.m[key] = value
	return nil
}
func (p *producer) Post(key string, value string) error {

	data, err := json.Marshal(fmt.Sprintf("%s;%s", key, value))
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
func (p *producer) Close() error {
	return p.file.Close()
}

//func NewProducer(fileName string) (*producer, error) {
//	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
//	if err != nil {
//		return nil, err
//	}
//	return &producer{
//		file:    file,
//		encoder: json.NewEncoder(file),
//	}, nil
//}

//
//func (s *producer) WriteRepo(key string) (string, error) {
//	s.lock.Lock()
//	defer s.lock.Unlock()
//
//	if v, ok := s.m[key]; ok {
//		return v, nil
//	}
//	return "", ErrNotFound
//}
//
//func (s *producer) ReadRepo(key string, value string) error {
//	s.lock.Lock()
//	defer s.lock.Unlock()
//
//	if _, ok := s.m[key]; ok {
//		return ErrAlreadyExists
//	}
//	s.m[key] = value
//	return nil
//}

//func (s *producer) Post(key string, value string) error {
//	s.lock.Lock()
//	defer s.lock.Unlock()
//
//	if _, ok := s.m[key]; ok {
//		return ErrAlreadyExists
//	}
//	s.m[key] = value
//	return nil
//}
