package repositories

import (
	"errors"
	"gorm.io/gorm"
	"keeping-track-backend-golang/api/models"
)

func MoveFundsToSavings(db *gorm.DB, amount float64, accID int) error {

	var err error
	b := &models.Balance{}
	err = db.Model(&b).Where("account_id = ?", accID).Take(&b).Error
	if err != nil {
		return err
	}
	if b.AvailableAmount-amount < 0 {
		return errors.New("There is not enough funds to move the requested amount to savings")
	}
	b.AvailableAmount = b.AvailableAmount - amount
	b.SavingsAmount = b.SavingsAmount + amount
	err = db.Save(&b).Error
	if err != nil {
		return err
	}
	return nil

}
