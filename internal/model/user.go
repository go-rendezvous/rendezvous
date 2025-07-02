package model

import "time"

type User struct {
	UserId    int
	Username  string
	Password  string
	Phone     string
	Email     string
	CreatedAt time.Time
	Meetings  []Meeting `gorm:"many2many:user_meeting_rule"`
}
