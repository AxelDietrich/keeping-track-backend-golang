package models

import (
	"database/sql"
	"time"
)

type DebtRecord struct {
	ID            int
	Name          string
	Amount        float64
	Payed         bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PayedAt       sql.NullTime
	SubcategoryID int
}
