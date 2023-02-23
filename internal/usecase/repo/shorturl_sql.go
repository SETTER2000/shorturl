package repo

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
)

const (
	driverName = "pgx"
)

type (
	producerSQL struct {
		ctx *context.Context
		db  *sql.DB
		//writer *bufio.Writer
	}

	consumerSQL struct {
		ctx *context.Context
		db  *sql.DB
		// заменяем reader на scanner
		//reader *bufio.Reader
		//sql.DB(QueryContext)
	}

	InSQL struct {
		cfg *config.Config
		r   *consumerSQL
		w   *producerSQL
	}
)

// NewInSQL слой взаимодействия с db в данном случаи с postgresql,
// хотя наверно можно объединить под эгидой всех db-sql-ориентированных (время покажет)
func NewInSQL(cfg *config.Config) *InSQL {
	return &InSQL{
		cfg: cfg,
		// создаём новый потребитель
		r: NewSQLConsumer(cfg),
		// создаём новый производитель
		w: NewSQLProducer(cfg),
	}
}

// NewSQLProducer производитель
func NewSQLProducer(cfg *config.Config) *producerSQL {
	context, connect := Connect(cfg)
	return &producerSQL{
		ctx: &context,
		db:  connect,
		//writer: bufio.NewWriter(file),
	}
}

func (i *InSQL) Post(sh *entity.Shorturl) error {
	//data, err := json.Marshal(&sh)
	//if err != nil {
	//	return err
	//}
	//// записываем событие в буфер
	//if _, err = i.w.writer.Write(data); err != nil {
	//	return err
	//}
	//// добавляем перенос строки
	//if err = i.w.writer.WriteByte('\n'); err != nil {
	//	return err
	//}
	//// записываем буфер в файл
	//t := i.w.writer.Flush()
	//return t
	// TODO POST producerSQL
	return nil
}

func (i *InSQL) Put(sh *entity.Shorturl) error {
	return i.Post(sh)
}

// NewSQLConsumer потребитель
func NewSQLConsumer(cfg *config.Config) *consumerSQL {
	context, connect := Connect(cfg)
	return &consumerSQL{
		ctx: &context,
		db:  connect,
		//reader: bufio.NewReader(file),
	}
}

func (i *InSQL) Get(key string) (*entity.Shorturl, error) {
	//sh := entity.Shorturl{}
	//if i.r.reader.Size() < 1 {
	//	return nil, ErrNotFound
	//}
	//for {
	//	data, err := i.r.reader.ReadBytes('\n')
	//	if err != nil {
	//		return nil, ErrNotFound
	//	}
	//
	//	err = json.Unmarshal(data, &sh)
	//	if err != nil {
	//		i.r.file.Seek(0, 0)
	//	}
	//
	//	if sh.Slug == key {
	//		break
	//	}
	//}
	//_, err := i.r.file.Seek(0, 0)
	//if err != nil {
	//	return nil, err
	//}
	//return &sh, nil
	// TODO GET /user/{user_id}/url/{url_id} consumerSQL
	return nil, nil
}

func (i *InSQL) GetAll(u *entity.User) (*entity.User, error) {
	//sh := entity.Shorturl{}
	//lst := entity.List{}
	//size := i.r.reader.Size()
	//if size < 1 {
	//	return nil, ErrNotFound
	//}
	//for j := 0; j < size; j++ {
	//	data, err := i.r.reader.ReadBytes('\n')
	//	if err != nil {
	//		break
	//	}
	//	err = json.Unmarshal(data, &sh)
	//	if err != nil {
	//		i.r.file.Seek(0, 0)
	//	}
	//	if sh.UserID == u.UserID {
	//		lst.URL = sh.URL
	//		lst.Slug = scripts.GetHost(i.cfg.HTTP, sh.Slug)
	//		u.Urls = append(u.Urls, lst)
	//	}
	//}
	//_, err := i.r.file.Seek(0, 0)
	//if err != nil {
	//	return nil, err
	//}
	//return u, nil
	// TODO GET /user/{user_id}/url consumerSQL
	return nil, nil
}

func Connect(cfg *config.Config) (ctx context.Context, db *sql.DB) {
	// Контекст позволяет ограничить по времени или прервать слишком долгие или уже не
	// нужные операции с базой данных, назначить для них дедлайн или тайм-аут.
	// Вот пример использования контекста:
	// конструируем контекст с 5-секундным тайм-аутом
	// после 5 секунд затянувшаяся операция с БД будет прервана
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// не забываем освободить ресурс
	defer cancel()
	//db, err := sql.Open("sqlite3", "db.db")
	db, err := sql.Open(driverName, cfg.ConnectDB)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return ctx, db
}
