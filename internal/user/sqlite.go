package user

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	databaseDriver = "sqlite3"
	databaseSource = "/Users/ilyaantonov/Downloads/ВАЖНОЕ/golang/infotecs-EWallet/internal/user/infotecs-EWallet.db"
)

func DatabaseConnection() (*sql.DB, error) {
	db, DBerr := sql.Open(databaseDriver, databaseSource)
	if DBerr != nil {
		return db, DBerr
	}

	return db, nil
}
