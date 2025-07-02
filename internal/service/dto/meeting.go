package dto

import (
	"time"

	"github.com/go-rendezvous/rendezvous/internal/model"
	cModel "github.com/go-rendezvous/rendezvous/pkg/model"
)

type MeetingInsertRequest struct {
	ScheduledAt    time.Time `json:"scheduled_at"`
	EndedAt        time.Time `json:"ended_at"`
	ParticipantIds []int     `json:"participant_ids"`
}

type MeetingDeleteRequest struct {
	MeetingNo string `json:"meeting_no"`
}

type MeetingUpdateRequest struct {
	MeetingNo    string              `json:"meeting_no"`
	BookedBy     int                 `json:"booked_by"`
	MeetingState cModel.MeetingState `json:"meeting_state"`
	ScheduledAt  time.Time           `json:"scheduled_at"`
	EndedAt      time.Time           `json:"ended_at"`
}

type MeetingListRequest struct {
	BookedBy int `json:"booked_by"`
}

type MeetingListResponse struct {
	MeetingNo    string
	BookedBy     string
	MeetingState string
	ScheduledAt  time.Time
	EndedAt      time.Time
	Users        []string
}

func GenMeetingListResponse(meeting *model.Meeting) MeetingListResponse {
	var users []string
	for _, user := range meeting.Users {
		users = append(users, user.Username)
	}

	return MeetingListResponse{
		MeetingNo:    meeting.MeetingNo,
		BookedBy:     meeting.BookedByName,
		MeetingState: meeting.MeetingState,
		ScheduledAt:  meeting.ScheduledAt,
		EndedAt:      meeting.EndedAt,
		Users:        users,
	}
}
