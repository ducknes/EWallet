package user

import "math/rand"

type user struct {
	WalletAddress       string `json:"wallet_address"`
	Balance             uint64 `json:"balance"`
	LastTransactionTime string `json:"last_transaction_time"`
}

func (u *user) NewUser() *user {
	return &user{
		WalletAddress:       u.randomWalletAddressGenerator(),
		Balance:             100,
		LastTransactionTime: "",
	}
}

func (u *user) randomWalletAddressGenerator() string {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newAdress := make([]byte, 25)
	for i := range newAdress {
		newAdress[i] = letters[rand.Intn(len(letters))]
	}
	return string(newAdress)
}
