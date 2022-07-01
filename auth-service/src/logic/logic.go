package logic

import "database/sql"

func GetUserSHA256(db *sql.DB, userID int64) (string, error) {
	var (
		SHA256 string
		err    error
	)
	row := db.QueryRow("select SHA256 from users where UserID = ?", userID)
	err = row.Scan(&SHA256)
	return SHA256, err
}
