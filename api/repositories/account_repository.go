package repositories

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/stdlib"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/requests"
)

func CreateAccount(db *sql.DB, account *models.Account) (*models.Account, error) {

	var err error
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return &models.Account{}, err
	}
	var isEmailUsed int
	err = tx.QueryRow("select count(*) from keepingtrack.accounts where email=$1;", account.Email).Scan(&isEmailUsed)
	if err != nil {
		return &models.Account{}, err
	}
	if isEmailUsed == 1 {
		return &models.Account{}, errors.New("Email already in use")
	}
	err = tx.QueryRow("insert into keepingtrack.accounts (username, email, password) values ($1, $2, $3) RETURNING *;",
		account.Username, account.Email, account.Password).Scan(&account.ID, &account.Username, &account.CreatedAt, &account.Email, &account.Password)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	_, err = tx.Exec("insert into keepingtrack.categories (name, income, account_id) values ($1, $2, $3);",
		"Savings", true, account.ID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	var categoryDebtID int

	err = tx.QueryRow("insert into keepingtrack.categories (name, income, account_id) values ($1, $2, $3) RETURNING id;",
		"Debt", false, account.ID).Scan(&categoryDebtID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	_, err = tx.Exec("insert into keepingtrack.categories (name, income, account_id) values ($1, $2, $3);",
		"Expenses", false, account.ID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	_, err = tx.Exec("insert into keepingtrack.categories (name, income, account_id) values ($1, $2, $3);",
		"Income", true, account.ID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	_, err = tx.Exec("insert into keepingtrack.subcategories (name, amount, category_id) values ($1, $2, $3);",
		"Auto-generated", 0, categoryDebtID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	_, err = tx.Exec("insert into keepingtrack.balances (available_amount, savings_amount, debt, account_id) values (0, 0, 0, $1);",
		account.ID)
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	tx.Commit()

	return account, nil

}

func GetAccount(db *sql.DB, login *requests.Login) (*models.Account, error) {
	var err error
	acc := &models.Account{}
	err = db.QueryRow("select * from keepingtrack.accounts where email = $1;", login.Email).
		Scan(&acc.ID, &acc.Username, &acc.CreatedAt, &acc.Email, &acc.Password)
	if err != nil {
		return acc, errors.New("Invalid account")
	}
	return acc, nil
}
