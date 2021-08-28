package cockroach

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type CockroachDBStore struct {
	db         *gorm.DB
	interfaces []interface{}
}

func NewCockroachDB(dsn string) (*CockroachDBStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, err
	}
	return &CockroachDBStore{db: db}, nil
}
func (c *CockroachDBStore) Close() error {
	d, err := c.db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CockroachDBStore) Migrate() error {
	return c.db.AutoMigrate(c.interfaces...)
}
