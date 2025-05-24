package entity

import "time"

type Wallet struct {
	ID            int `gorm:"primaryKey"`
	UserID        int32
	AccountNumber string
	Balance       int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type WalletInput struct {
	UserID int32
	Amount int32
}
