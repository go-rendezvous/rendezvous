package api

import (
	"errors"
	"net/http"

	"github.com/go-rendezvous/rendezvous/pkg/model"
	"github.com/go-rendezvous/rendezvous/pkg/service"
	"github.com/labstack/echo/v4"
)

type BaseApi struct {
	Ctx    echo.Context
	Errors error
}

func (a *BaseApi) BindContext(ctx echo.Context) *BaseApi {
	a.Ctx = ctx
	return a
}

func (a *BaseApi) Bind(value interface{}) *BaseApi {
	err := a.Ctx.Bind(value)
	a.addError(err)
	return a
}

func (a *BaseApi) MakeService(s *service.Service) *BaseApi {
	s.Ctx = a.Ctx
	return a
}

func (a *BaseApi) Error(message string) error {
	resp := model.Response{
		Message: message,
	}
	return a.Ctx.JSON(http.StatusBadRequest, resp)
}

func (a *BaseApi) OK(message string, value interface{}) error {
	resp := model.Response{
		Message: message,
		Value:   value,
	}
	return a.Ctx.JSON(http.StatusOK, resp)
}

func (a *BaseApi) addError(err error) {
	a.Errors = errors.Join(a.Errors, err)
}
