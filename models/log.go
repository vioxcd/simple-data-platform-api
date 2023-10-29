package models

import "gorm.io/gorm"

type UserLog struct {
	gorm.Model
	UserID     uint   `json:"userId"`
	User       User   `json:"user"`
	RunID      uint   `json:"runId"`
	Run        Run    `json:"run"`
	Parameters string `json:"parameters" gorm:"not null"`
}
