package entity

type User struct {
	ID   int    `gorm:"primaryKey"`
	Name string 
}
type UserResponse struct {
	UserID	int32
	Status   string  
	Name string 
	AccountNumber string 
}