package cockroach

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type CommonDataService interface {
	Close() error
	Migrate() error
	Drop() error
}

type DBInterfaces []interface{}

type CDBService struct {
	Db         *gorm.DB
	Logger     *logrus.Logger
	Context    context.Context
	Interfaces DBInterfaces
}

func NewCDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *CDBService) Close() error {
	d, err := c.Db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CDBService) Migrate() error {
	return c.Db.AutoMigrate(c.Interfaces...)
}

func (c *CDBService) Drop() error {
	return c.Db.Migrator().DropTable(c.Interfaces...)
}
