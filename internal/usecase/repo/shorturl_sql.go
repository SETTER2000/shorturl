package repo

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

const (
	driverName = "pgx"
)

type (
	producerSQL struct {
		db *sqlx.DB
	}

	consumerSQL struct {
		db *sqlx.DB
	}

	InSQL struct {
		cfg *config.Config
		r   *consumerSQL
		w   *producerSQL
	}
)

// NewInSQL слой взаимодействия с db в данном случаи с postgresql
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
	connect := Connect(cfg)
	return &producerSQL{
		db: connect,
	}
}

func (i *InSQL) Post(ctx context.Context, sh *entity.Shorturl) error {
	stmt, err := i.w.db.Prepare("INSERT INTO public.shorturl (slug, url, user_id) VALUES ($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(sh.Slug, sh.URL, sh.UserID)
	if err, ok := err.(*pgconn.PgError); ok {
		if err.Code == pgerrcode.UniqueViolation {
			return NewConflictError("old url", "http://testiki", ErrAlreadyExists)
		}
	}
	//if sh.CorrelationOrigin == nil {
	//	_, err := stmt.Exec(sh.Slug, sh.URL, sh.UserID)
	//	if err, ok := err.(*pgconn.PgError); ok {
	//		if err.Code == pgerrcode.UniqueViolation {
	//			return NewConflictError("old url", "http://testiki", ErrAlreadyExists)
	//		}
	//	}
	//} else {
	//	for _, j := range *sh.CorrelationOrigin {
	//		res, err := stmt.Exec(j.Slug, j.URL, sh.UserID)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		r, err := res.RowsAffected()
	//		if err != nil {
	//			fmt.Printf("Error Affected: %e\n", err)
	//		}
	//		fmt.Printf("RowsAffected: %v\n", r)
	//	}
	//}
	return nil
}

func (i *InSQL) Put(ctx context.Context, sh *entity.Shorturl) error {
	return i.Post(ctx, sh)
}

// NewSQLConsumer потребитель
func NewSQLConsumer(cfg *config.Config) *consumerSQL {
	connect := Connect(cfg)
	return &consumerSQL{
		db: connect,
	}
}

func (i *InSQL) Get(ctx context.Context, key string) (*entity.Shorturl, error) {
	var slug, url, id string
	rows, err := i.w.db.Query("SELECT slug, url, user_id FROM shorturl WHERE slug = $1", key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&slug, &url, &id)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	sh := entity.Shorturl{}
	sh.Slug = slug
	sh.URL = url
	sh.UserID = id
	return &sh, nil
}

func (i *InSQL) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	var slug, url, id string
	q := `SELECT * FROM shorturl WHERE user_id=$1`
	rows, err := i.w.db.Queryx(q, u.UserID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	l := entity.List{}
	for rows.Next() {
		err = rows.Scan(&slug, &url, &id)
		if err != nil {
			return nil, err
		}
		l.URL = url
		l.Slug = scripts.GetHost(i.cfg.HTTP, slug)
		u.Urls = append(u.Urls, l)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return u, nil
}

func Connect(cfg *config.Config) (db *sqlx.DB) {
	db, _ = sqlx.Open(driverName, cfg.ConnectDB)
	err := db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	n := 100
	db.SetMaxIdleConns(n)
	db.SetMaxOpenConns(n)
	schema := `CREATE TABLE IF NOT EXISTS public.user
(
   id   VARCHAR(300) NOT NULL
);
CREATE TABLE IF NOT EXISTS public.shorturl
(
   slug    VARCHAR(300) NOT NULL,
   url     VARCHAR NOT NULL UNIQUE,
   user_id VARCHAR(300) NOT NULL
);
`
	db.MustExec(schema)
	if err != nil {
		panic(err)
	}
	return db
}
