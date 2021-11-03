package sqlxcockroach

import (
	"github.com/jmoiron/sqlx"
)

type CommonDataService interface {
	Close() error
	Migrate() error
}

type CockroachDBStore struct {
	Db *sqlx.DB
}

func NewCockroachDBConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *CockroachDBStore) Close() error {
	return c.Db.Close()
}

func (c *CockroachDBStore) Migrate(sql string) error {
	_, err := c.Db.Exec(sql)
	return err
}
