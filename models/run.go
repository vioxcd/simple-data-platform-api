package models

import "gorm.io/gorm"

type Run struct {
	gorm.Model
	Type        string `gorm:"not null"`
	Description string
}
