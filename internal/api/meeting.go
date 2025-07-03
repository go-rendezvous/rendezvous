package api

import (
	"github.com/go-rendezvous/rendezvous/internal/service"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/pkg/api"
	"github.com/labstack/echo/v4"
)

type Meeting struct {
	api.BaseApi
}

func (e Meeting) List(ctx echo.Context) error {
	var err error
	var req dto.MeetingListRequest
	s := service.Meeting{}
	err = e.BindContext(ctx).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		return e.Error(err.Error())
	}

	list, err := s.List(&req)
	if err != nil {
		return e.Error(err.Error())
	}

	return e.OK("query successful", list)
}

func (e Meeting) Insert(ctx echo.Context) error {
	var err error
	var req dto.MeetingInsertRequest
	s := service.Meeting{}
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

	return e.OK("insert successful", nil)
}
