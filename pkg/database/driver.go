package database

import (
	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	postgresScheme string = "postgres"
	sqliteScheme   string = "sqlite"
	mysqlScheme    string = "mysql"
)

type Driver interface {
	GetOrm() *gorm.DB
}

func NewDBDriver(conf *config.Conf) (Driver, error) {
	var driver Driver
	var err error
	driverType := conf.DiverType
	dest := conf.Destination

	cfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	}

	switch driverType {
	case postgresScheme:
		driver, err = newPostgresDriver(dest, cfg)
	case sqliteScheme:
		driver, err = newSqliteDriver(dest, cfg)
	case mysqlScheme:
		driver, err = newMysqlDriver(dest, cfg)
	default:
		logrus.Fatal("Not supporting schema: ", driverType)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return driver, nil
}
