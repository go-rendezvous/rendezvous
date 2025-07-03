package dbstore

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type meetingStore struct {
	db *gorm.DB
}

func (m *meetingStore) List(userId int) ([]model.Meeting, error) {
	var err error
	var meetingList []model.Meeting
	err = m.db.Preload("Users").
		Where("booked_by = ?", userId).
		Find(&meetingList).
		Error

	if err != nil {
		return nil, err
	}

	return meetingList, nil
}

func (m *meetingStore) Insert(meeting *model.Meeting) error {
	var err error
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return m.db.Create(meeting).Error
}

func (m *meetingStore) Delete(meetingNo string) error {
	return nil
}

func (m *meetingStore) Update(meeting *model.Meeting) error {
	return nil
}

func newMeetingStore(s *datastore) store.MeetingStore {
	return &meetingStore{db: s.db}
}
