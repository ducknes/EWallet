package user

import (
	"math/rand"
	"time"
)

type User struct {
	ID                  int    `json:"id"`
	WalletAddress       string `json:"wallet_address"`
	Balance             uint64 `json:"balance"`
	LastTransactionTime int64  `json:"last_transaction_time"`
}

type Send struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

func (u *User) NewUser() *User {
	return &User{
		WalletAddress:       u.randomWalletAddressGenerator(),
		Balance:             100,
		LastTransactionTime: int64(rand.Intn(int(time.Now().Unix()))),
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
