package repo

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/app/er"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

const (
	driverName = "pgx"
)

// InSQL .-
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
func NewInSQL(db *sqlx.DB, c *config.Config) *InSQL {
	return &InSQL{
		cfg: c,
		// создаём новый потребитель
		r: NewSQLConsumer(db),
		// создаём новый производитель
		w: NewSQLProducer(db),
	}
}

// NewSQLProducer производитель
func NewSQLProducer(db *sqlx.DB) *producerSQL {
	return &producerSQL{
		db: db,
	}
}

// Post - добавляет данные в DB
func (i *InSQL) Post(ctx context.Context, sh *entity.Shorturl) error {
	if len(strings.TrimSpace(string(sh.UserID))) < 1 {
		return er.ErrBadRequest
	}
	stmt, err := i.w.db.Prepare("INSERT INTO public.shorturl (slug, url, user_id) VALUES ($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("error repo post sql: %e", err)
	}

	_, err = stmt.Exec(sh.Slug, sh.URL, sh.UserID)
	if err, ok := err.(*pgconn.PgError); ok {
		if err.Code == pgerrcode.UniqueViolation {
			return NewConflictError("old url", "http://testiki", er.ErrAlreadyExists)
		}
	}

	return nil
}

// Put - обновляет данные
func (i *InSQL) Put(ctx context.Context, sh *entity.Shorturl) error {
	return i.Post(ctx, sh)
}

// NewSQLConsumer потребитель
func NewSQLConsumer(db *sqlx.DB) *consumerSQL {
	return &consumerSQL{
		db: db,
	}
}

// Get -.
func (i *InSQL) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	var slug entity.Slug
	var url entity.URL
	var id entity.UserID
	var del bool
	rows, err := i.w.db.Query("SELECT slug, url, user_id, del FROM shorturl WHERE slug = $1 OR url = $2 ", sh.Slug, sh.URL)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("err from shorturl_repository: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&slug, &url, &id, &del)
		if err != nil {
			log.Fatal(err)
			return nil, fmt.Errorf("err from shorturl_repository: %s", err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("err from shorturl_repository: %s", err.Error())
	}
	//sh := entity.Shorturl{}
	sh.Slug = slug
	sh.URL = url
	sh.UserID = id
	sh.Del = del
	return sh, nil
}

// GetAll - получить все данные.
func (i *InSQL) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	var id string
	var slug entity.Slug
	var url entity.URL
	q := `SELECT slug, url, user_id FROM shorturl WHERE user_id=$1 AND del=$2`
	rows, err := i.r.db.Queryx(q, u.UserID, false)
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
		//l.Slug = slug
		l.ShortURL = scripts.GetHost(i.cfg.HTTP, slug)
		u.Urls = append(u.Urls, l)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return u, nil
}

// GetAllUrls - получить все url-адреса не зависимо от авторизации.
func (i *InSQL) GetAllUrls() (entity.CountURLs, error) {
	var c int
	err := i.w.db.Get(&c, "SELECT count(*) FROM shorturl")
	if err != nil {
		return 0, err
	}

	return entity.CountURLs(c), nil
}

// GetAllUsers - получить все url-адреса не зависимо от авторизации.
func (i *InSQL) GetAllUsers() (entity.CountUsers, error) {
	var c int
	err := i.w.db.Get(&c, "SELECT COUNT(DISTINCT user_id) FROM shorturl")
	if err != nil {
		return 0, err
	}

	return entity.CountUsers(c), nil
}

// Delete -.
func (i *InSQL) Delete(ctx context.Context, u *entity.User) error {
	q := `UPDATE shorturl SET del = $1
	FROM (SELECT unnest($2::text[]) AS slug) AS data_table
	WHERE shorturl.slug = data_table.slug AND shorturl.user_id=$3`

	rows, err := i.w.db.Queryx(q, true, u.DelLink, u.UserID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

// New - создает instance DB, возвращает готовое соединение.
func New(cfg *config.Config) (*sqlx.DB, error) {
	db, _ := sqlx.Open(driverName, cfg.Storage.ConnectDB)
	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
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
   user_id VARCHAR(300) NOT NULL,
   del BOOLEAN DEFAULT FALSE
);
`
	db.MustExec(schema)
	return db, nil
}

// Read -.
func (i *InSQL) Read() error {
	return nil
}

// Save -.
func (i *InSQL) Save() error {
	return nil
}
