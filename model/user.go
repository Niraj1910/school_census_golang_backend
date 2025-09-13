package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:VARCHAR(255);not null;uniqueIndex"`
	Password  string    `json:"-" gorm:"not null" binding:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
