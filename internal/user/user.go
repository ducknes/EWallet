package user

import "time"

type user struct {
	WalletAddress       string    `json:"wallet_address"`
	Balance             uint64    `json:"balance"`
	LastTransactionTime time.Time `json:"last_transaction_time"`
}

func (u *user) NewUser() *user {
	return &user{
		WalletAddress:       "",
		Balance:             0,
		LastTransactionTime: time.Now(),
	}
}
