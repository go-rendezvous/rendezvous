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

		db := driver.GetOrm().Debug()
		factory = &datastore{
			db: db,
		}

	})
	return factory
}

type datastore struct {
	db *gorm.DB
}

func (s *datastore) Migrate() error {
	return s.db.AutoMigrate(&model.Meeting{}, &model.User{})
}

func (s *datastore) MeetingStore() store.MeetingStore {
	return newMeetingStore(s)
}

func (s *datastore) UserStore() store.UserStore {
	return newUserStore(s)
}
