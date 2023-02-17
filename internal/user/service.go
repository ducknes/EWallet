package user

import (
	"database/sql"
	"fmt"
)

type service struct {
	DB *sql.DB
}

func (s *service) WatchUserBalance(params string) {
	database, err := DatabaseConnection()
	if err != nil {
		panic(err)
	}
	s.DB = database
	walletAddress := params
	fmt.Println(walletAddress)
}

// TODO service layer
