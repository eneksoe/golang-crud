package entities

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model

	FirstName string     `gorm:"type:varchar(100);not null"`
	LastName  string     `gorm:"type:varchar(100);not null"`
	BirthDate *time.Time `gorm:"type:date;not null"`
	Gender    string     `gorm:"not null"`
	Email     string     `gorm:"type:varchar;not null"`
	Address   string     `gorm:"type:varchar(200)"`
}
