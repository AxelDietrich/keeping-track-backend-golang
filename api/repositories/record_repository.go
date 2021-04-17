package repositories

import (
	"database/sql"
	"errors"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/requests"
)

func CreateRecord(db *sql.DB, r *requests.AddRecordRequest, subID int) error {

	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.Record{}
	rec.SubcategoryID = subID
	rec.Name = r.Name
	rec.Amount = r.Amount
	err = tx.QueryRow("insert into keepingtrack.records (name, amount, subcategory_id) values ($1, $2, $3);",
		rec.Name, rec.Amount, rec.SubcategoryID).Err()
	if err != nil {
		return err
		tx.Rollback()
	}
	err = PropagateRecordChanges(tx, rec)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(db *sql.DB, recID int, r *requests.AddRecordRequest) error {

	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.Record{}
	err = tx.QueryRow("select * from keepingtrack.records where id = $1;", recID).
		Scan(&rec.ID, &rec.Name, &rec.Amount, &rec.CreatedAt, &rec.UpdatedAt, &rec.SubcategoryID)
	if err != nil {
		return err
	}
	previousAmount := rec.Amount
	err = tx.QueryRow("update keepingtrack.records (name, amount) values ($1, $2) RETURNING *;", r.Name, r.Amount).
		Scan(&rec.ID, &rec.Name, &rec.Amount, &rec.CreatedAt, &rec.UpdatedAt, &rec.SubcategoryID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if previousAmount > r.Amount {
		rec.Amount = (previousAmount - rec.Amount) * -1
		err = PropagateRecordChanges(tx, rec)
		if err != nil {
			return err
		}
	} else if previousAmount < r.Amount {
		rec.Amount = rec.Amount - previousAmount
		err = PropagateRecordChanges(tx, rec)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteRecord(db *sql.DB, recID int) error {

	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.Record{}
	err = tx.QueryRow("select * from keepingtrack.records where id = $1;", recID).
		Scan(&rec.ID, &rec.Name, &rec.Amount, &rec.CreatedAt, &rec.UpdatedAt, &rec.SubcategoryID)
	if err != nil {
		return err
	}
	err = tx.QueryRow("delete from keepingtrack.records where id = $1;", recID).Err()
	if err != nil {
		return err
	}
	rec.Amount = rec.Amount * -1
	err = PropagateRecordChanges(tx, rec)
	if err != nil {
		return err
	}
	return nil

}

func PropagateRecordChanges(tx *sql.Tx, r *models.Record) error {

	var err error
	sub := &models.Subcategory{}
	err = tx.QueryRow("select * from keepingtrack.subcategories where id = $1;", r.SubcategoryID).
		Scan(&sub.ID, &sub.Name, &sub.Amount, &sub.CategoryID, &sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		return err
	}
	sub.Amount += r.Amount
	err = tx.QueryRow("update keepingtrack.subcategories set amount = $1 where id = $2;",
		sub.Amount, sub.ID).Err()
	if err != nil {
		return err
	}
	categ := &models.Category{}
	err = tx.QueryRow("select * from keepingtrack.categories where id = $1;", sub.CategoryID).
		Scan(&categ.ID, &categ.Name, &categ.AccountID, &categ.Income)
	balance := &models.Balance{}
	err = tx.QueryRow("select * from keepingtrack.balances where account_id = $1;", categ.AccountID).
		Scan(&balance.ID, &balance.AvailableAmount, &balance.SavingsAmount, &balance.Debt, &balance.AccountID)
	if err != nil {
		return err
	}
	currentAmount := balance.AvailableAmount
	if categ.Income {
		if categ.Name == "Income" {
			if currentAmount+r.Amount >= 0 { //if true, then the value of r.Amount is negative and the deleted income corresponds to spent funds
				balance.AvailableAmount = currentAmount + r.Amount
			} else {
				balance.AvailableAmount = 0
				balance.Debt = balance.Debt + r.Amount
				debtRecord := &models.DebtRecord{}
				debtRecord.Name = r.Name
				debtRecord.Amount = r.Amount - currentAmount
				subAutoDebt := &models.Subcategory{}
				err = tx.QueryRow("select s.id, s.name, s.amount, s.category_id, s.created_at, s.updated_at from keepingtrack.subcategories s inner join keepingtrack.categories c on s.category_id = c.id where s.name = 'Auto-generated' and c.account_id = $1;", balance.AccountID).
					Scan(&subAutoDebt.ID, &subAutoDebt.Name, &subAutoDebt.Amount, &subAutoDebt.CategoryID, &subAutoDebt.CreatedAt, &subAutoDebt.UpdatedAt)
				if err != nil {
					tx.Rollback()
					return err
				}
				subAutoDebt.Amount += r.Amount
				err = tx.QueryRow("update keepingtrack.subcategories set amount = $1 where id = $2;",
					subAutoDebt.Amount, subAutoDebt.ID).Err()
				if err != nil {
					tx.Rollback()
					return err
				}
				debtRecord.SubcategoryID = subAutoDebt.ID
				err = tx.QueryRow("insert into keepingtrack.records (name, amount, subcategory_id) values ($1, $2, $3);",
					debtRecord.Name, debtRecord.Amount, debtRecord.SubcategoryID).Err()
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			if balance.AvailableAmount-r.Amount >= 0 {
				balance.AvailableAmount -= r.Amount
				balance.SavingsAmount += r.Amount
			} else {
				return errors.New("There is not enough funds to move the requested amount to savings")
			}
		}
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
			debtRecord := &models.Record{}
			debtRecord.Name = r.Name
			debtRecord.Amount = r.Amount - currentAmount
			subAutoDebt := &models.Subcategory{}
			err = tx.QueryRow("select s.id, s.name, s.amount, s.category_id, s.created_at, s.updated_at from keepingtrack.subcategories s inner join keepingtrack.categories c on s.category_id = c.id where s.name = 'Auto-generated' and c.account_id = $1;", balance.AccountID).
				Scan(&subAutoDebt.ID, &subAutoDebt.Name, &subAutoDebt.Amount, &subAutoDebt.CategoryID, &subAutoDebt.CreatedAt, &subAutoDebt.UpdatedAt)
			if err != nil {
				tx.Rollback()
				return err
			}
			subAutoDebt.Amount += r.Amount
			err = tx.QueryRow("update keepingtrack.subcategories set amount = $1 where id = $2;",
				subAutoDebt.Amount, subAutoDebt.ID).Err()
			if err != nil {
				tx.Rollback()
				return err
			}
			debtRecord.SubcategoryID = subAutoDebt.ID
			err = tx.QueryRow("insert into keepingtrack.debt_records (name, amount, subcategory_id, payed) values ($1, $2, $3, false);",
				debtRecord.Name, debtRecord.Amount, debtRecord.SubcategoryID).Err()
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.QueryRow("update keepingtrack.balances set available_amount = $1, savings_amount = $2, debt = $3;",
		balance.AvailableAmount, balance.SavingsAmount, balance.Debt).Err()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
