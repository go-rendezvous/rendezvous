package database

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type meeting struct {
	db *gorm.DB
}

func (m *meeting) List(req *dto.MeetingListRequest) ([]dto.MeetingListResponse, error) {
	var err error
	var meetingList []model.Meeting
	var respList []dto.MeetingListResponse
	if err = m.db.Find(&meetingList).Error; err != nil {
		return nil, err
	}

	for _, meeting := range meetingList {
		respList = append(respList, dto.GenMeetingListResponse(&meeting))
	}

	return respList, nil
}

func (m *meeting) Create(req *dto.MeetingInsertRequest) error {
	return nil
}
func (m *meeting) Delete(req *dto.MeetingDeleteRequest) error {
	return nil
}
func (m *meeting) Update(req *dto.MeetingUpdateRequest) error {
	return nil
}

func newMeetingStore(s *datastore) store.MeetingStore {
	return &meeting{db: s.db}
}
