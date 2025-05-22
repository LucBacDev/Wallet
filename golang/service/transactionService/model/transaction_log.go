package model

import "time"

type TransactionLog struct {
	ID              uint   `gorm:"primaryKey"`
	SenderID        string `gorm:"column:sender_id"`
	ReceiverID      string `gorm:"column:receiver_id"`
	Amount          int32  `gorm:"column:amount"`
	TransactionType string `gorm:"column:transaction_type"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
