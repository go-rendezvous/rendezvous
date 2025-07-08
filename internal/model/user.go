package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    int            `gorm:"primaryKey"`
	Username  string         `gorm:"uniqueIndex"`
	Password  string         `gorm:"not null"`
	Phone     string         `gorm:"uniqueIndex"`
	Email     string         `gorm:"uniqueIndex"`
	CreatedAt time.Time      `gorm:"not null"`
	Meetings  []Meeting      `gorm:"many2many:user_meeting_rule"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
