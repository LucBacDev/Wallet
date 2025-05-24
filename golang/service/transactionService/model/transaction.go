package model

import "time"

type Transaction struct {
	ID         string `gorm:"primaryKey" json:"id"`
	SenderID   int32
	ReceiverID int32
	Amount     int32
	Type       string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}