package models

type Address struct {
	ID      string `gorm:"primaryKey;type:TEXT;not null;default:''" json:"id"`
	UserID  string `gorm:"type:TEXT;not null;default:''" json:"userId"`
	Street  string `gorm:"not null" json:"street"`
	City    string `gorm:"not null" json:"city"`
	State   string `gorm:"not null" json:"state"`
	ZipCode string `gorm:"column:zipcode;not null" json:"zipcode"`
}
