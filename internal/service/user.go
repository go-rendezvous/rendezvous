package service

import (
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/internal/store/dbstore"
	"github.com/go-rendezvous/rendezvous/pkg/service"
)

type User struct {
	service.Service
}

func (s User) List(req *dto.UserListRequest) ([]dto.UserListResponse, error) {
	users, err := dbstore.Factory().UserStore().List(req.Ids)
	if err != nil {
		return nil, err
	}

	var respList []dto.UserListResponse
	for _, user := range users {
		respList = append(respList, dto.GenUserListResponse(user))
	}
	return respList, nil
}

func (s User) Insert(req *dto.UserInsertRequest) error {
	var err error

	return err
}
