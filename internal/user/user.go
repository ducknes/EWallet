package user

import "time"

type user struct {
	WalletAddress       string
	Balance             uint64
	LastTransactionTime time.Time
}

func (u *user) NewUser() *user {
	return &user{
		WalletAddress:       "",
		Balance:             0,
		LastTransactionTime: time.Now(),
	}
}
