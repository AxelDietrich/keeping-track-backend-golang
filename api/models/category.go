package models

type Category struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Income    bool    `json:"income"`
	AccountID int     `json:"-"`
	Account   Account `json:"-"`
}
