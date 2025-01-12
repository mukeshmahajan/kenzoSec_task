package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" binding:"required" json:"title"`
	Description string    `gorm:"not null" binding:"required" json:"description"`
	Status      string    `gorm:"type:text;default:'pending'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
