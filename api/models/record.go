package models

import (
	"time"
)

type Record struct {
	ID            int         `gorm:"primary_key;auto_increment" json:"id"`
	Name          string      `json:"name"`
	Amount        float64     `json:"amount"`
	Subcategory   Subcategory `json:"-"`
	SubcategoryID int         `json:"-"`
	CreatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Record) TableName() string {
	return "keepingtrack.records"
}
