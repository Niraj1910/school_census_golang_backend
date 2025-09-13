package model

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" binding:"required" gorm:"not null"`
	Address   string    `json:"address" binding:"required" gorm:"not null"`
	City      string    `json:"city" binding:"required" gorm:"not null"`
	State     string    `json:"state" binding:"required" gorm:"not null"`
	Contact   string    `json:"contact" binding:"required" gorm:"not null"`
	Image     string    `json:"image"`
	Email     string    `json:"email" binding:"required,email" gorm:"not null;uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (School) TableName() string {
	return "schools"
}
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&School{}, &User{})
}
