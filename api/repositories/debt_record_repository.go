package repositories

import (
	"database/sql"
	"keeping-track-backend-golang/api/models"
	"keeping-track-backend-golang/api/requests"
)

func UpdateDebRecord(db *sql.DB, recID int, r *requests.AddRecordRequest) error {
	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.DebtRecord{}
	err = tx.QueryRow("select * from keepingtrack.debt_records where id = $1", recID).
		Scan(&rec.Name, &rec.ID, &rec.Amount, &rec.Payed, &rec.CreatedAt, &rec.UpdatedAt, &rec.PayedAt, &rec.SubcategoryID)
	if err != nil {
		return err
	}
	previousAmount := rec.Amount
	err = tx.QueryRow("update keepingtrack.debt_records set name = $1, amount = $2 RETURNING *;", r.Name, r.Amount).
		Scan(&rec.Name, &rec.ID, &rec.Amount, &rec.Payed, &rec.CreatedAt, &rec.UpdatedAt, &rec.PayedAt, &rec.SubcategoryID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if previousAmount > r.Amount {
		rec.Amount = (previousAmount - rec.Amount) * -1
		err = PropagateDebtRecord(tx, rec)
		if err != nil {
			return err
		}
	} else if previousAmount < r.Amount {
		rec.Amount = rec.Amount - previousAmount
		err = PropagateDebtRecord(tx, rec)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteDebtRecord(db *sql.DB, recID int) error {
	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.DebtRecord{}
	err = tx.QueryRow("select * from keepingtrack.debt_records where id = $1", recID).
		Scan(&rec.Name, &rec.ID, &rec.Amount, &rec.Payed, &rec.CreatedAt, &rec.UpdatedAt, &rec.PayedAt, &rec.SubcategoryID)
	if err != nil {
		tx.Rollback()
		return err
	}
	rec.Amount = rec.Amount * -1
	err = PropagateDebtRecord(tx, rec)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func CreateDebtRecord(db *sql.DB, r *requests.AddRecordRequest, subID int) error {
	var err error
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	rec := &models.DebtRecord{}
	err = tx.QueryRow("insert into keepingtrack.debt_records (name, amount, subcategory_id, payed) values ($1, $2, $3, false) RETURNING *;",
		r.Name, r.Amount, subID).
		Scan(&rec.Name, &rec.ID, &rec.Amount, &rec.Payed, &rec.CreatedAt, &rec.UpdatedAt, &rec.PayedAt, &rec.SubcategoryID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = PropagateDebtRecord(tx, rec)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

func PropagateDebtRecord(tx *sql.Tx, r *models.DebtRecord) error {

	var err error
	sub := &models.Subcategory{}
	err = tx.QueryRow("update keepingtrack.subcategories set amount = amount + $1 where id = $2 RETURNING *;", r.Amount, r.SubcategoryID).
		Scan(&sub.ID, &sub.Name, &sub.Amount, &sub.CategoryID, &sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.QueryRow("update keepingtrack.balances as b set debt = debt + $1 from keepingtrack.categories as c where (c.account_id = b.account_id and c.id = $2);", r.Amount, sub.CategoryID).Err()

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
