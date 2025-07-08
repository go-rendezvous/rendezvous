package api

import (
	"github.com/go-rendezvous/rendezvous/internal/service"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/pkg/api"
	"github.com/labstack/echo/v4"
)

type System struct {
	api.BaseApi
}

func (a System) Login(ctx echo.Context) error {
	var err error
	var req dto.LoginRequest
	s := service.System{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	token, err := s.Login(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", token)
}

func (a System) Logout(ctx echo.Context) error {
	var err error
	var req dto.LogoutRequest
	s := service.System{}
	err = a.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return a.Error(err.Error())
	}

	err = s.Logout(&req)
	if err != nil {
		return a.Error(err.Error())
	}

	return a.OK("success", nil)
}
