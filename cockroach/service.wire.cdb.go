//go:build wireinject
// +build wireinject

package cockroach

import (
	"context"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitCDBService(ctx context.Context, logger *logrus.Logger, db *gorm.DB, interfaces DBInterfaces) (CDBService, error) {
	wire.Build(wire.Struct(new(CDBService), "Db", "Logger", "Context", "Interfaces"))
	return CDBService{}, nil
}
