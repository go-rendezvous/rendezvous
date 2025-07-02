package db

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

func newSqliteDriver(destination string) (Driver, error) {
	db, err := gorm.Open(sqlite.Open(destination))
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
