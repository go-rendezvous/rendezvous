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
	r.POST("/meeting", meeting.Insert)
	r.DELETE("/meeting/:meeting_no", meeting.Delete)
	r.PUT("/meeting", meeting.Update)
	r.GET("/meeting", meeting.List)
}
