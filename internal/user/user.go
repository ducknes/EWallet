package user

import (
	"math/rand"
	"time"
)

type User struct {
	WalletAddress       string    `json:"wallet_address"`
	Balance             uint64    `json:"balance"`
	LastTransactionTime time.Time `json:"last_transaction_time"`
}

func (u *User) NewUser() *User {
	return &User{
		WalletAddress:       u.randomWalletAddressGenerator(),
		Balance:             100,
		LastTransactionTime: time.Now(),
	}
}

func (u *User) randomWalletAddressGenerator() string {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newAddress := make([]byte, 64)
	for i := range newAddress {
		newAddress[i] = letters[rand.Intn(len(letters))]
	}
	return string(newAddress)
}
