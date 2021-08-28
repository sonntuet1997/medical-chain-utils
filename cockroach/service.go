package cockroach

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type CockroachDBStore struct {
	Db         *gorm.DB
	Interfaces []interface{}
}

func NewCockroachDBConnection(dsn string) (*gorm.DB , error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *CockroachDBStore) Close() error {
	d, err := c.Db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CockroachDBStore) Migrate() error {
	return c.Db.AutoMigrate(c.Interfaces...)
}
