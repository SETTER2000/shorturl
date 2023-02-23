package repo

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

const (
	driverName = "pgx"
)

type (
	producerSQL struct {
		db *pgxpool.Pool
		//writer *bufio.Writer
	}

	consumerSQL struct {
		db *pgxpool.Pool
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
	connect := Connect(cfg)
	return &producerSQL{
		db: connect,
		//writer: bufio.NewWriter(file),
	}
}

func (i *InSQL) Post(ctx context.Context, sh *entity.Shorturl) error {
	var slug string
	q := `INSERT INTO public.shorturl (slug, url, user_id) VALUES ($1,$2,$3) RETURNING slug`
	if err := i.w.db.QueryRow(ctx, q, sh.Slug, sh.URL, sh.UserID).Scan(&slug); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Sprintf("SQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState())
			fmt.Println(newErr)
			return nil
		}
		fmt.Printf("%s", err)
		return err
	}
	//fmt.Printf("new slug: %s\n", slug)
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
		//reader: bufio.NewReader(file),
	}
}

func (i *InSQL) Get(ctx context.Context, key string) (*entity.Shorturl, error) {
	var slug, url, id string
	q := `SELECT * FROM shorturl WHERE slug=$1`
	if err := i.w.db.QueryRow(ctx, q, key).Scan(&slug, &url, &id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Sprintf("SQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState())
			fmt.Println(newErr)
			return nil, err
		}
		fmt.Printf("%s", err)
		return nil, err
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
	rows, err := i.w.db.Query(ctx, q, u.UserID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Sprintf("SQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState())
			fmt.Println(newErr)
			return nil, err
		}
		fmt.Printf("%s", err)
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
	return u, nil
}

func Connect(cfg *config.Config) (db *pgxpool.Pool) {
	dbpool, err := pgxpool.New(context.Background(), cfg.ConnectDB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()
	return dbpool
}
