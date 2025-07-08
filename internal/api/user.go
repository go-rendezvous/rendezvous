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

func (a User) Insert(ctx echo.Context) error {
	var err error
	var req dto.UserInsertRequest
	s := service.User{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	err = s.Insert(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", nil)
}

func (a User) Delete(ctx echo.Context) error {
	var err error
	var req dto.UserDeleteRequest
	s := service.User{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	err = s.Delete(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", nil)
}

func (a User) Update(ctx echo.Context) error {
	var err error
	var req dto.UserUpdateRequest
	s := service.User{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	err = s.Update(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", nil)
}

func (a User) List(ctx echo.Context) error {
	var err error
	var req dto.UserListRequest
	s := service.User{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	list, err := s.List(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", list)
}
