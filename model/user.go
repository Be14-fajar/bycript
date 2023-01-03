package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Email    string `json:"email" form:"email"`
	Password string `json:"pwd" form:"password"`
}
type UserModel struct {
	DB *gorm.DB
}

func (um *UserModel) Insert(Newuser User) (User, error) {
	err := um.DB.Create(&Newuser).Error

	if err != nil {
		return User{}, err
	}
	return Newuser, nil
}

func Bcript(y string) string {
	password := []byte(y)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)

}
