package domain

import "time"

type Item struct {
	Id          string    `json:"id"`
	Product     string    `json:"product"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CheckedAt   time.Time `json:"checked_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
