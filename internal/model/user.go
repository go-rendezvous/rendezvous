package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    int            `gorm:"primaryKey"`
	Username  string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Phone     string         `gorm:"unique;not null"`
	Email     string         `gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"not null"`
	Meetings  []Meeting      `gorm:"many2many:user_meeting_rule"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
