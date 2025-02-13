package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"not null" json:"firstname"`
	LastName  string    `gorm:"not null" json:"lastname"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Address   Address   `gorm:"foreignKey:UserID" json:"address"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
