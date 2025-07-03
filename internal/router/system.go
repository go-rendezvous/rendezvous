package router

import (
	"github.com/go-rendezvous/rendezvous/internal/service"
	"github.com/labstack/echo/v4"
)

func registerSystemRouter(r *echo.Echo) {
	r.GET("/login", service.Login)
	r.POST("/logout", service.Logout)
}
