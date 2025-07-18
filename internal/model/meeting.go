package model

import (
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	MeetingId    int            `gorm:"primaryKey"`
	MeetingNo    string         `gorm:"uniqueIndex"`
	BookedBy     int            `gorm:"not null"`
	BookedByName string         `gorm:"size:64;not null" comment:"who has booked this meeting"`
	MeetingState string         `gorm:"size:4;not null"`
	ScheduledAt  time.Time      `gorm:"not null"`
	EndedAt      time.Time      `gorm:"not null"`
	Users        []User         `gorm:"many2many:user_meeting_rule"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
