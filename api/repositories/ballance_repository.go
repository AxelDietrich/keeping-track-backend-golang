package repositories

import (
	"database/sql"
	"errors"
	"keeping-track-backend-golang/api/models"
)

func MoveFundsToSavings(db *sql.DB, amount float64, accID int) error {

	var err error
	b := &models.Balance{}
	err = db.QueryRow("select * from keepingtrack.balances where account_id = $1;", accID).
		Scan(&b.ID, &b.AvailableAmount, &b.SavingsAmount, &b.Debt, &b.AccountID)
	if err != nil {
		return err
	}
	if b.AvailableAmount-amount < 0 {
		return errors.New("There is not enough funds to move the requested amount to savings")
	}
	b.AvailableAmount = b.AvailableAmount - amount
	b.SavingsAmount = b.SavingsAmount + amount
	err = db.QueryRow("update keepingtrack.balances set available_amount = $1, savings_amount = $2 where id = $3;",
		&b.AvailableAmount, &b.SavingsAmount, &b.ID).Err()
	if err != nil {
		return err
	}
	return nil

}

func GetBalance(db *sql.DB, accID int) (*models.Balance, error) {
	b := &models.Balance{}
	err := db.QueryRow("select * from keepingtrack.balances where account_id = $1", accID).
		Scan(&b.ID, &b.AvailableAmount, &b.SavingsAmount, &b.Debt, &b.AccountID)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func AddIncome(db *sql.DB, amount float64, accID int) error {

	var err error
	err = db.QueryRow("update keepingtrack.balances set available_amount = available_amount + $1 where account_id = $2;",
		amount, accID).Err()
	if err != nil {
		return err
	}
	return nil
}
