package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(u *User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	passWord, err := GenerateHashSalt(u.Password)
	if err != nil {
		return errors.New(err.Error())
	}

	u.Password = passWord

	return nil
}

func (u User) ComparePassword(secret string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(secret),
	)

	if err == nil {
		return true
	}

	return false
}
