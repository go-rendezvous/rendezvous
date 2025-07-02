package api

import (
	"errors"
	"net/http"

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

func (a *BaseApi) Error(value interface{}) error {
	return a.Ctx.JSON(http.StatusBadRequest, value)
}

func (a *BaseApi) OK(value interface{}) error {
	return a.Ctx.JSON(http.StatusOK, value)
}

func (a *BaseApi) addError(err error) {
	a.Errors = errors.Join(a.Errors, err)
}
