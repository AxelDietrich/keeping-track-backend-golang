package repositories

import (
	"gorm.io/gorm"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/requests"
)

func CreateRecord(db *gorm.DB, r *requests.AddRecordRequest, subID int) error {

	var err error
	tx := db.Begin()
	rec := &models.Record{}
	rec.SubcategoryID = subID
	rec.Name = r.Name
	rec.Amount = r.Amount
	err = tx.Create(&rec).Error
	if err != nil {
		return err
		tx.Rollback()
	}
	err = PropagateRecordChanges(db, subID, r)
	if err != nil {
		return err
	}
	return nil
}

func PropagateRecordChanges(tx *gorm.DB, subID int, r *requests.AddRecordRequest) error {

	sub, err := GetSubcategory(tx, subID)
	if err != nil {
		return err
	}
	sub.Amount = sub.Amount + r.Amount
	err = tx.Save(&sub).Error
	if err != nil {
		return err
	}
	categ := &models.Category{}
	err = tx.Model(&models.Category{}).Where("id = ?", sub.CategoryID).Take(&categ).Error
	accID := categ.AccountID
	balance := &models.Balance{}
	err = tx.Model(&models.Balance{}).Where("account_id = ?", accID).Take(&balance).Error
	if err != nil {
		return err
	}
	currentAmount := balance.AvailableAmount
	if sub.Category.Income {
		balance.AvailableAmount = currentAmount + r.Amount
	} else {
		if currentAmount-r.Amount >= 0 {
			if currentAmount == 0 && r.Amount < 0 {
				debt := balance.Debt
				if debt+r.Amount < 0 {
					balance.Debt = 0
					balance.AvailableAmount = (debt + r.Amount) * -1
				}
			} else {
				balance.AvailableAmount = currentAmount - r.Amount
			}
		} else {
			balance.AvailableAmount = 0
			balance.Debt = balance.Debt + (r.Amount - currentAmount)
			recordDebt := &models.Record{}
			recordDebt.Name = r.Name
			recordDebt.Amount = r.Amount - currentAmount
			subAutoDebt := &models.Subcategory{}
			err = tx.Model(&models.Subcategory{}).Joins("INNER JOIN keepingtrack.categories c ON keepingtrack.subcategories.category_id = c.id").
				Where("keepingtrack.subcategories.name = ?", "Auto-generated").Take(&subAutoDebt).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			recordDebt.Subcategory = *subAutoDebt
			recordDebt.SubcategoryID = subAutoDebt.ID
			err = tx.Create(&recordDebt).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Save(&balance).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
