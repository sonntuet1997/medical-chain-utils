package cockroach

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type CommonDataService interface {
	Close() error
	Migrate() error
}

type CDBService struct {
	Db         *ExtendGorm
	Interfaces []interface{}
}

func NewCDBConnection(dsn string) (*ExtendGorm, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, err
	}
	return &ExtendGorm{DB: *db}, nil
}

func (c *CDBService) Close() error {
	d, err := c.Db.DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CDBService) Migrate() error {
	return c.Db.AutoMigrate(c.Interfaces...)
}
