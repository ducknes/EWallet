package user

import (
	"database/sql"
	"log"
)

type Repository struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) AddNewUser(u *User) error {
	tx, err := r.DB.Begin()
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

func (r *Repository) GetUserByWalletAddress(wa string) *User {
	query := `SELECT * FROM users WHERE wallet_address = ?`
	var user User

	if err := r.DB.QueryRow(query, wa).Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime); err != nil {
		log.Printf("user not found. err: %s\n", err)
	}

	return &user
}
