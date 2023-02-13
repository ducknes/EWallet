package sqlite

import (
	"database/sql"
)

func DatabaseConnection() (*sql.DB, error) {
	db, DBerr := sql.Open("sqlite3", "infotecs-EWallet.db")
	if DBerr != nil {
		return db, DBerr
	}

	return db, nil
}
