package models

import "gorm.io/gorm"

type Run struct {
	gorm.Model
	Type        string `json:"type" gorm:"not null"`
	Description string `json:"description"`
}
