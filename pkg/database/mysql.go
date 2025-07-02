package database

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

func newMysqlDriver(destination string, opts ...gorm.Option) (Driver, error) {
	db, err := gorm.Open(mysql.Open(destination), opts...)
	if err != nil {
		return nil, err
	}
	return &sqliteDriver{db: db}, nil
}
