package store

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
)

// store for meeting
type MeetingStore interface {
	Insert(meeting *model.Meeting) error
	Delete(meetingNo string) error
	Update(meeting *model.Meeting) error
	List(userId int) ([]*model.Meeting, error)
}
