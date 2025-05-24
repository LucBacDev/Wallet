package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int32    `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex" `
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Token		string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashPassword)

	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
