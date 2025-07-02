package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDriver struct {
	db *gorm.DB
}

func (d *postgresDriver) GetOrm() *gorm.DB {
	return d.db
}

func newPostgresDriver(destination string) (Driver, error) {
	db, err := gorm.Open(postgres.Open(destination))
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
