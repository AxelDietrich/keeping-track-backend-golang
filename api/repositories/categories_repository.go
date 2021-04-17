package repositories

import "database/sql"

func CheckIfUserCategory(db *sql.DB, catID int, accID int) error {
	var a int
	err := db.QueryRow("select id from keepingtrack.categories where id = $1 and account_id = $2", catID, accID).Scan(&a)
	if err != nil && a != 0 {
		return err
	}
	return nil
}
