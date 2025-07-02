package model

import (
	"time"

	"github.com/go-rendezvous/rendezvous/pkg/model"
)

type Meeting struct {
	MeetingId    int                `gorm:"primaryKey"`
	MeetingNo    string             `gorm:"unique;not null"`
	MeetingState model.MeetingState `gorm:"size:4;not null"`
	ScheduledAt  time.Time          `gorm:"not null"`
	EndedAt      time.Time          `gorm:"not null"`
	Users        []User             `gorm:"many2many:user_meeting_rule"`
}
