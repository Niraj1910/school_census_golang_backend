package model

import "time"

type School struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
	City      string    `json:"city" binding:"required"`
	State     string    `json:"state" binding:"required"`
	Contact   string    `json:"contact" binding:"required"`
	Image     string    `json:"image"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
