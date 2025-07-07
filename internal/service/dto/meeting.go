package dto

import (
	"time"

	"github.com/go-rendezvous/rendezvous/internal/model"
)

type MeetingInsertRequest struct {
	ScheduledAt    time.Time `json:"scheduled_at"`
	EndedAt        time.Time `json:"ended_at"`
	ParticipantIds []int     `json:"participant_ids"`
}

type MeetingDeleteRequest struct {
	MeetingNo string `param:"meeting_no"`
}

type MeetingUpdateRequest struct {
	MeetingNo      string    `json:"meeting_no"`
	MeetingState   string    `json:"meeting_state"`
	ScheduledAt    time.Time `json:"scheduled_at"`
	EndedAt        time.Time `json:"ended_at"`
	ParticipantIds []int     `json:"participant_ids"`
}

type MeetingListRequest struct {
	BookedBy int `query:"booked_by"`
}

type MeetingListResponse struct {
	MeetingNo    string    `json:"meeting_no"`
	BookedBy     string    `json:"booked_by"`
	MeetingState string    `json:"meeting_state"`
	ScheduledAt  time.Time `json:"scheduled_at"`
	EndedAt      time.Time `json:"ended_at"`
	Users        []string  `json:"users"`
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
