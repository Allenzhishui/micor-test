package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordCost = 12
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}

func (user *User) SetPassword(password string) error {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(fromPassword)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
