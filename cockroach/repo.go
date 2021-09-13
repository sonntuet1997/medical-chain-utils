package cockroach

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommonDataRepo interface {
	Close() error
	Migrate() error
	Drop() error
	RawSql(string) error
}

type CDBRepo struct {
	Db         *gorm.DB
	Logger     *logrus.Logger
	Context    context.Context
	Interfaces DBInterfaces `wire:"-"`
}

func (c *CDBRepo) Close() error {
	d, err := c.Db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CDBRepo) Migrate() error {
	return c.Db.AutoMigrate(c.Interfaces...)
}

func (c *CDBRepo) Drop() error {
	return c.Db.Migrator().DropTable(c.Interfaces...)
}

func (c *CDBRepo) RawSql(a string) error {
	return c.Db.Exec(a).Error
}
