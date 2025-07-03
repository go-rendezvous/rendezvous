package router

import (
	"github.com/go-rendezvous/rendezvous/internal/api"
	"github.com/labstack/echo/v4"
)

func init() {
	auth = append(auth, registerUserRouter)
}

func registerUserRouter(r *echo.Group) {
	user := api.User{}
	r.GET("/user", user.List)
}
