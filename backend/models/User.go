package models

type User struct {
	ID       string  `gorm:"primaryKey;type:TEXT;not null;default:''" json:"id"`
	Name     string  `gorm:"type:TEXT;not null;default:''" json:"name"`
	Username string  `gorm:"type:TEXT;not null;default:''" json:"username"`
	Email    string  `gorm:"type:TEXT;not null;unique;default:''" json:"email"`
	Phone    string  `gorm:"type:TEXT;not null;default:''" json:"phone"`
	Address  Address `gorm:"foreignKey:UserID" json:"address"`
	Posts    []Post  `gorm:"foreignKey:UserID" json:"posts"`
}
