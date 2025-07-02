package dbstore

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type meeting struct {
	db *gorm.DB
}

func (m *meeting) List(userId int) ([]model.Meeting, error) {
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

func (m *meeting) Insert(meeting *model.Meeting) error {
	return m.db.Create(meeting).Error
}

func (m *meeting) Delete(meetingNo string) error {
	return nil
}

func (m *meeting) Update(meeting *model.Meeting) error {
	return nil
}

func newMeetingStore(s *datastore) store.MeetingStore {
	return &meeting{db: s.db}
}
