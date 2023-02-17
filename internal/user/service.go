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

	stmt, err := database.Prepare("SELECT * FROM users where wallet_address = ?")
	if err != nil {
		panic(err)
	}

	rows, errQuery := stmt.Query(params)
	if errQuery != nil {
		panic(errQuery)
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime)
		if err != nil {
			panic(err)
		}
	}

	jsonB, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		panic(err)
	}

	return jsonB
}

// TODO service layer
