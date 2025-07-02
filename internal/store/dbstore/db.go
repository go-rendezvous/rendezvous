package dbstore

import (
	"sync"

	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"github.com/go-rendezvous/rendezvous/pkg/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var once sync.Once
var factory store.DBFactory

func Factory() store.DBFactory {
	once.Do(func() {
		driver, err := database.NewDBDriver(config.GetConf())
		if err != nil {
			logrus.Fatal(err)
		}

		db := driver.GetOrm()
		factory = &datastore{
			db: db,
		}

		driver.GetOrm().AutoMigrate(&model.Meeting{}, &model.User{})
	})
	return factory
}

type datastore struct {
	db *gorm.DB
}

func (s *datastore) MeetingStore() store.MeetingStore {
	return newMeetingStore(s)
}
