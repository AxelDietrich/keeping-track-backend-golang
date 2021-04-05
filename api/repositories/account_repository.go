package repositories

import (
	"errors"
	"gorm.io/gorm"
	"keeping-track-backend-golang/api/models"
	"time"
)

func CreateAccount(db *gorm.DB, account *models.Account) (*models.Account, error) {

	tx := db.Begin()
	var err error
	var acc *models.Account
	err = tx.Where("email = ?", account.Email).First(&acc).Error
	if err == nil {
		return acc, errors.New("Email already used")
	}
	err = tx.Create(&account).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	var balance models.Balance
	balance.Account = *account
	balance.AvailableAmount = 0
	balance.Debt = 0
	balance.SavingsAmount = 0
	balance.AccountID = account.Id

	err = tx.Create(&balance).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	var categorySavings models.Category
	categorySavings.Account = *account
	categorySavings.AccountID = account.Id
	categorySavings.Income = true
	categorySavings.Name = "Savings"

	err = tx.Create(&categorySavings).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}
	var categoryDebt models.Category
	categoryDebt.Account = *account
	categoryDebt.AccountID = account.Id
	categoryDebt.Income = false
	categoryDebt.Name = "Debt"

	err = tx.Create(&categoryDebt).Take(&categoryDebt).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}
	var categoryExpenses models.Category
	categoryExpenses.Account = *account
	categoryExpenses.AccountID = account.Id
	categoryExpenses.Income = false
	categoryExpenses.Name = "Expenses"

	err = tx.Create(&categoryExpenses).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	var categoryIncome models.Category
	categoryIncome.Account = *account
	categoryIncome.AccountID = account.Id
	categoryIncome.Income = true
	categoryIncome.Name = "Income"

	var subcategoryAutoGenerated models.Subcategory
	subcategoryAutoGenerated.Category = categoryDebt
	subcategoryAutoGenerated.CategoryID = categoryDebt.ID
	subcategoryAutoGenerated.Name = "Auto-generated"
	subcategoryAutoGenerated.Amount = 0
	subcategoryAutoGenerated.CreatedAt = time.Now()
	subcategoryAutoGenerated.UpdatedAt = time.Now()
	err = tx.Create(&subcategoryAutoGenerated).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	err = tx.Create(&categoryIncome).Error
	if err != nil {
		tx.Rollback()
		return &models.Account{}, err
	}

	tx.Commit()

	return account, nil

}

func GetAccount(db *gorm.DB, login *models.Account) (*models.Account, error) {
	var err error
	var acc *models.Account
	err = db.Where("email = ?", login.Email).First(&acc).Error
	if err != nil {
		return acc, errors.New("Invalid account")
	}
	return acc, nil
}
