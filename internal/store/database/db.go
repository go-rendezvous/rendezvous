package database

import (
	"sync"

	"github.com/go-rendezvous/rendezvous/internal/store"
	"github.com/go-rendezvous/rendezvous/pkg/db"
	"gorm.io/gorm"
)

var once sync.Once
var factory store.DBFactory

func Factory(dbDriver db.Driver) store.DBFactory {
	once.Do(func() {
		factory = &datastore{
			db: dbDriver.GetOrm(),
		}
	})
	return factory
}

type datastore struct {
	db *gorm.DB
}

func (s *datastore) MeetingStore() store.MeetingStore {
	return newMeetingStore(s)
}
