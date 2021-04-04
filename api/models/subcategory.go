package models

import (
	"time"
)

type Subcategory struct {
	ID	int64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
