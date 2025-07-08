package service

import (
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/internal/store/dbstore"
	"github.com/go-rendezvous/rendezvous/pkg/service"
)

type User struct {
	service.Service
}

func (s User) Insert(req *dto.UserInsertRequest) error {
	user := req.GenModel()

	err := dbstore.Factory().UserStore().Insert(user)

	return err
}

func (s User) Delete(req *dto.UserDeleteRequest) error {
	return dbstore.Factory().UserStore().Delete(req.UserId)
}

func (s User) Update(req *dto.UserUpdateRequest) error {
	user := req.GenModel()

	return dbstore.Factory().UserStore().Update(user)
}

func (s User) List(req *dto.UserListRequest) ([]dto.UserListResponse, error) {
	users, err := dbstore.Factory().UserStore().List(req.Ids)
	if err != nil {
		return nil, err
	}

	var respList []dto.UserListResponse
	for _, user := range users {
		respList = append(respList, dto.GenUserListResponse(&user))
	}
	return respList, nil
}
