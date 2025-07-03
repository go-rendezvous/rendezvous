package router

import (
	"github.com/go-rendezvous/rendezvous/internal/api"
	"github.com/labstack/echo/v4"
)

func init() {
	auth = append(auth, registerMeetingRouter)
}

func registerMeetingRouter(r *echo.Group) {
	meeting := api.Meeting{}
	r.GET("/meeting", meeting.List)
	r.POST("/meeting", meeting.Insert)
}
