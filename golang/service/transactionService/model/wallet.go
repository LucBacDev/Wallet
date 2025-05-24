package model

type DebitInfo struct {
	UserID int32
	Amount int32
}

type CreditInfo struct {
	UserID int32
	Amount int32
}
type WalletResult struct {
	Success bool
	Message string
}