package api

import (
	"github.com/go-rendezvous/rendezvous/internal/service"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/pkg/api"
	"github.com/labstack/echo/v4"
)

type User struct {
	api.BaseApi
}

func (e User) List(ctx echo.Context) error {
	var err error
	var req dto.UserListRequest
	s := service.User{}
	err = e.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return e.Error(err)
	}

	list, err := s.List(&req)
	if err != nil {
		return e.Error(err)
	}

	return e.OK(list)
}

func (e User) Insert(ctx echo.Context) error {
	var err error
	var req dto.UserInsertRequest
	s := service.User{}
	err = e.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return e.Error(err.Error())
	}

	err = s.Insert(&req)
	if err != nil {
		return e.Error(err.Error())
	}

	return e.OK("insert successful")
}
