package service

import (
	"errors"

	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/internal/store/dbstore"
	"github.com/go-rendezvous/rendezvous/pkg/consts"
	"github.com/go-rendezvous/rendezvous/pkg/service"
	"github.com/go-rendezvous/rendezvous/pkg/util"
)

type Meeting struct {
	service.Service
}

func (s Meeting) List(req *dto.MeetingListRequest) ([]dto.MeetingListResponse, error) {
	meetings, err := dbstore.Factory().MeetingStore().List(req.BookedBy)
	if err != nil {
		return nil, err
	}

	var respList []dto.MeetingListResponse
	for _, meeting := range meetings {
		respList = append(respList, dto.GenMeetingListResponse(meeting))
	}
	return respList, nil
}

func (s Meeting) Insert(req *dto.MeetingInsertRequest) error {
	var err error
	claims, err := s.ExtractClaims()
	if err != nil {
		return err
	}

	if len(req.ParticipantIds) == 0 {
		req.ParticipantIds = append(req.ParticipantIds, claims.UserId)
	}

	// TODO: needs optimization
	users, err := dbstore.Factory().UserStore().List(req.ParticipantIds)
	if err != nil {
		return err
	}

	if len(users) != len(req.ParticipantIds) {
		return errors.New("invalid participants")
	}

	meeting := model.Meeting{
		MeetingNo:    util.GenerateCode(config.GetConf().MeetingNoLength),
		BookedBy:     claims.UserId,
		BookedByName: claims.Username,
		MeetingState: consts.Booked,
		ScheduledAt:  req.ScheduledAt,
		EndedAt:      req.EndedAt,
		Users:        users,
	}

	err = dbstore.Factory().MeetingStore().Insert(&meeting)

	return err
}
