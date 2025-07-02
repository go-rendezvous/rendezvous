package api

import (
	"github.com/go-rendezvous/rendezvous/pkg/api"
	"github.com/labstack/echo/v4"
)

type Meeting struct {
	api.BaseApi
}

func (e Meeting) Insert(ctx echo.Context) {

}
