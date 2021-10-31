package repository

import "github.com/vagnerlg/supermaketlist/src/domain"

type ItemRepository interface {
	Insert(item domain.Item) domain.Item
	All() []domain.Item
}
