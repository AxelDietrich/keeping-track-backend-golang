package models

type Balance struct {
	ID              int     `json:"id"`
	AvailableAmount float64 `json:"available_amount"`
	SavingsAmount   float64 `json:"savings_amount"`
	Debt            float64 `json:"debt"`
	AccountID       int     `json:"-"`
}
