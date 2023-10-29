package models

import "gorm.io/gorm"

type UserLog struct {
	gorm.Model
	UserID     int    `json:"userId"`
	User       User   `json:"user"`
	RunID      int    `json:"runId"`
	Run        Run    `json:"run"`
	Parameters string `json:"parameters" gorm:"not null"`
}
