package repository

import (
	"github.com/vagnerlg/supermaketlist/src/domain"
)

type RepositoryItem interface {
	Insert(item domain.Item) domain.Item
	All() []domain.Item
	First(id string) (domain.Item, error)
}

type RepositoryUser interface {
	Insert(user domain.User) domain.User
	First(id string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Login(email string, password string) bool
}
