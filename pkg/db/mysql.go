package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDriver struct {
	db *gorm.DB
}

func (d *mysqlDriver) GetOrm() *gorm.DB {
	return d.db
}

func newMysqlDriver(destination string) (Driver, error) {
	db, err := gorm.Open(mysql.Open(destination))
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
