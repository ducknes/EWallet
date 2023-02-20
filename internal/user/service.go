package user

import (
	"database/sql"
	"encoding/json"
)

type service struct {
	DB *sql.DB
}

func (s *service) WatchUserBalance(params string) []byte {
	database, err := DatabaseConnection()
	if err != nil {
		panic(err)
	}

	query := `SELECT * FROM users where wallet_address = ?`

	var user User

	if err := database.QueryRow(query, params).Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime); err != nil {
		panic(err)
	}

	jsonB, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		panic(err)
	}

	return jsonB
}

func (s *service) GetUsersTransactions(count int) []byte {
	database, err := DatabaseConnection()
	if err != nil {
		panic(err)
	}

	query := `SELECT * FROM users WHERE last_transaction_time != 0 ORDER BY last_transaction_time DESC LIMIT ?`
	var users []User

	rows, err := database.Query(query, count)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var user User
		if errScan := rows.Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime); errScan != nil {
			panic(errScan)
		}
		users = append(users, user)
	}
	jsonB, errJson := json.Marshal(users)
	if errJson != nil {
		panic(errJson)
	}
	return jsonB
}

func (s *service) PostSendMoney() {}
