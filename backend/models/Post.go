package models

import "time"

type Post struct {
	ID        string    `gorm:"primaryKey;type:TEXT;not null;default:''" json:"id"`
	UserID    string    `gorm:"type:TEXT;not null;default:''" json:"userId"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}
