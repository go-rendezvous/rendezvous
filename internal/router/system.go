package router

import (
	"github.com/go-rendezvous/rendezvous/internal/api"
	"github.com/labstack/echo/v4"
)

func registerSystemRouter(r *echo.Echo) {
	system := api.System{}
	r.GET("/login", system.Login)
	r.POST("/logout", system.Logout)
}
