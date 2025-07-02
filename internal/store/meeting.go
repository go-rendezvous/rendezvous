package store

import "github.com/go-rendezvous/rendezvous/internal/service/dto"

// store for meeting
type MeetingStore interface {
	Create(req *dto.MeetingInsertRequest) error
	Delete(req *dto.MeetingDeleteRequest) error
	Update(req *dto.MeetingUpdateRequest) error
	List(req *dto.MeetingListRequest) ([]dto.MeetingListResponse, error)
}
