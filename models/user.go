package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
