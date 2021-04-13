package models

import "time"

type DebtRecord struct {
	ID            int
	Name          string
	Amount        float64
	Payed         bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PayedAt       time.Time
	SubcategoryID int
}
