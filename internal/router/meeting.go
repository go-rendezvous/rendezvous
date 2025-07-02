package router

import (
	"github.com/go-rendezvous/rendezvous/internal/api"
	"github.com/labstack/echo/v4"
)

func init() {
	auth = append(auth, meetingRouter)
}

func meetingRouter(r *echo.Group) {
	r.Group("/meeting")

	meeting := api.Meeting{}
	r.GET("", meeting.List)
	r.POST("", meeting.Insert)
}
