package sqlxcockroach

import (
	"github.com/jmoiron/sqlx"
)

type CommonDataService interface {
	Close() error
	Migrate(sql string) error
}

type CockroachCDBRepo struct {
	Db *sqlx.DB
}

func NewCockroachDBConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *CockroachCDBRepo) Close() error {
	return c.Db.Close()
}

func (c *CockroachCDBRepo) Migrate(sql string) error {
	_, err := c.Db.Exec(sql)
	return err
}
