package domain

import (
	"errors"
	"time"
)

type Item struct {
	Id          string    `json:"id"`
	Product     string    `json:"product"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CheckedAt   time.Time `json:"checked_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(i Item) (Item, error) {
	err := i.validate()
	if err != nil {
		return i, err
	}

	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()

	return i, nil
}

func (i Item) validate() error {
	if i.Product == "" {
		return errors.New("product is required")
	}

	if i.Description == "" {
		return errors.New("description is required")
	}

	if i.Amount == 0 {
		return errors.New("amount is required")
	}

	return nil
}
