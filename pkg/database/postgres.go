package database

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

func newPostgresDriver(destination string, opts ...gorm.Option) (Driver, error) {
	db, err := gorm.Open(postgres.Open(destination), opts...)
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
