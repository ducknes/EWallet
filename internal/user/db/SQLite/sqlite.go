package sqlite

import (
	"database/sql"
	"infotecs-EWallet/internal/user"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnection() (*sql.DB, error) {
	db, DBerr := sql.Open("sqlite3", "/Users/ilyaantonov/Downloads/ВАЖНОЕ/golang/infotecs-EWallet/internal/user/db/sqlite/infotecs-EWallet.db")
	if DBerr != nil {
		return db, DBerr
	}

	return db, nil
}

func AddNewUser(db *sql.DB, u *user.User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users (wallet_address, balance, last_transaction_time) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.WalletAddress, u.Balance, u.LastTransactionTime)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

// how to сount row in database table
