package models

import (
	"time"
)

type Snapshot struct {
	Id     uint      `json:"id"`
	Date   time.Time `json:"date"`
	Amount uint      `json:"amount"`
}
