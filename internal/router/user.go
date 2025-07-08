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
	r.POST("/user", user.Insert)
	r.DELETE("/user/:user_id", user.Delete)
	r.PUT("/user", user.Update)
	r.GET("/user", user.List)
}
