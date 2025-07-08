package dbstore

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type meetingStore struct {
	db *gorm.DB
}

func (s *meetingStore) GetMeeting(meetingNo string) (model.Meeting, error) {
	meeting := model.Meeting{}
	err := s.db.Where("meeting_no = ?", meetingNo).Find(&meeting).Error
	return meeting, err
}

func (s *meetingStore) Insert(meeting *model.Meeting) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(meeting).Error
	})
}

func (s *meetingStore) Delete(meetingNo string) error {
	var err error
	meeting := &model.Meeting{}
	err = s.db.Find(meeting, "meeting_no = ?", meetingNo).Error
	if err != nil {
		return err
	}

	err = s.db.Model(meeting).Association("Users").Clear()
	if err != nil {
		return err
	}

	return s.db.Where("meeting_no = ?", meetingNo).Delete(meeting).Error
}

func (s *meetingStore) Update(meeting *model.Meeting) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.db.Model(meeting).Association("Users").Replace(meeting.Users)
		if err != nil {
			return err
		}

		return s.db.Save(meeting).Error
	})
}

func (s *meetingStore) List(userId int) ([]model.Meeting, error) {
	var err error
	var meetingList []model.Meeting
	err = s.db.Preload("Users").
		Where("booked_by = ?", userId).
		Find(&meetingList).
		Error

	if err != nil {
		return nil, err
	}

	return meetingList, nil
}

func newMeetingStore(s *datastore) store.MeetingStore {
	return &meetingStore{db: s.db}
}
