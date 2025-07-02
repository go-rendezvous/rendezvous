package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDriver struct {
	db *gorm.DB
}

func (d *sqliteDriver) GetOrm() *gorm.DB {
	return d.db
}

func newSqliteDriver(destination string, opts ...gorm.Option) (Driver, error) {
	db, err := gorm.Open(sqlite.Open(destination), opts...)
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
