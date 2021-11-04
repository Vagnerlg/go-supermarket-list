package repository

import "github.com/vagnerlg/supermaketlist/src/domain"

type UserRepository interface {
	Insert(user domain.User) domain.User
	First(id string) domain.User
	FindByEmail(email string) domain.Item
	Login(email string, password string) bool
}
