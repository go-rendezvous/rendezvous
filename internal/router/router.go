package router

import (
	"github.com/go-rendezvous/rendezvous/internal/pkg/middleware"
	"github.com/labstack/echo/v4"
)

var auth = make([]func(r *echo.Group), 0)

func InitRouter(r *echo.Echo) {
	initSysRouter(r)

	initBusinessRouter(r)
}

func initSysRouter(r *echo.Echo) {
	registerSystemRouter(r)
}

func initBusinessRouter(r *echo.Echo) {
	group := r.Group("/api/v1", middleware.JsonWebToken())

	for _, f := range auth {
		f(group)
	}
}
