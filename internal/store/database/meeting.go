package database

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"gorm.io/gorm"
)

type meeting struct {
	db *gorm.DB
}

func (m *meeting) List(req *dto.MeetingListRequest) ([]dto.MeetingListResponse, error) {
	var err error
	var meetingList []model.Meeting
	var respList []dto.MeetingListResponse
	if err = m.db.Find(meetingList).Error; err != nil {
		return nil, err
	}

	for _, meeting := range meetingList {
		respList = append(respList, dto.GenMeetingListResponse(&meeting))
	}

	return respList, nil
}
