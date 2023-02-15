package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "infotecs-EWallet/internal/user"
)

func DatabaseConnection() (*sql.DB, error) {
	db, DBerr := sql.Open("sqlite3", "infotecs-EWallet.db")
	if DBerr != nil {
		return db, DBerr
	}

	return db, nil
}

func AddNewUser(db *sql.DB, u *user) error {
	tx, err := db.Begin()
	temp := 
}
