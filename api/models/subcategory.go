package models

import (
	"time"
)

type Subcategory struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	CategoryID int       `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
