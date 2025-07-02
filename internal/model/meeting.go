package model

import (
	"time"

	"github.com/go-rendezvous/rendezvous/pkg/model"
)

type Meeting struct {
	MeetingId    int
	MeetingNo    string
	MeetingState model.MeetingState `gorm:"size:4"`
	ScheduledAt  time.Time
	EndedAt      time.Time
	Users        []User `gorm:"many2many:user_meeting_rule"`
}
