package database

import (
	"sync"

	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

var once sync.Once
var factory store.DBFactory

func Factory(db *gorm.DB) store.DBFactory {
	once.Do(func() {
		factory = &datastore{
			db: db,
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
