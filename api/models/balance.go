package models

type Balance struct {

	ID	int64 `gorm:"primary_key;auto_increment" json:"id"`
	AvailableAmount float64 `json:"available_amount"`
	SavingsAmount float64 `json:"savings_amount"`
}