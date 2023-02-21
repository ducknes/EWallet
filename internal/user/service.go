package user

import (
	"encoding/json"
	"log"
	"time"
)

type service struct {
	Repository *Repository
}

func NewService(repository *Repository) *service {
	return &service{
		Repository: repository,
	}
}

func (s *service) WatchUserBalance(params string) []byte {
	query := `SELECT * FROM users where wallet_address = ?`

	var user User

	if err := s.Repository.DB.QueryRow(query, params).Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime); err != nil {
		log.Printf("пользователь не найден. ошибка: %s\n", err)
	}

	jsonB, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		log.Printf("ошибка во время сериализации json. ошибка: %s\n", errMarshal)
	}

	return jsonB
}

func (s *service) GetUsersTransactions(count int) []byte {
	query := `SELECT * FROM users WHERE last_transaction_time != 0 ORDER BY last_transaction_time DESC LIMIT ?`
	var users []User

	rows, err := s.Repository.DB.Query(query, count)
	if err != nil {
		log.Printf("пользователь не найден. ошибка: %s\n", err)
	}
	for rows.Next() {
		var user User
		if errScan := rows.Scan(&user.ID, &user.WalletAddress, &user.Balance, &user.LastTransactionTime); errScan != nil {
			log.Printf("ошибка сканирования данных из бд. ошибка: %s\n", errScan)
		}
		users = append(users, user)
	}
	jsonB, errJson := json.Marshal(users)
	if errJson != nil {
		log.Printf("ошибка во время сериализации json. ошибка: %s\n", errJson)
	}
	return jsonB
}

func (s *service) PostSendMoney(sd Send) {
	userFrom := s.Repository.GetUserByWalletAddress(sd.From)
	userTo := s.Repository.GetUserByWalletAddress(sd.To)
	queryUpdate := `UPDATE users SET balance = ?, last_transaction_time = ? WHERE wallet_address = ?`

	if userFrom.Balance < sd.Amount {
		log.Printf("на счету позльзователя %s недостаточно средств\n", sd.From)
		return
	}

	userFrom.Balance -= sd.Amount
	userFrom.LastTransactionTime = time.Now().Unix()
	if _, err := s.Repository.DB.Exec(queryUpdate, userFrom.Balance, userFrom.LastTransactionTime, userFrom.WalletAddress); err != nil {
		log.Printf("ошибка при выполнении запроса к бд. ошибка: %s\n", err)
	}

	userTo.Balance += sd.Amount
	userTo.LastTransactionTime = time.Now().Unix()
	if _, err := s.Repository.DB.Exec(queryUpdate, userTo.Balance, userTo.LastTransactionTime, userTo.WalletAddress); err != nil {
		log.Printf("ошибка при выполнении запроса к бд. ошибка: %s\n", err)
	}
	log.Printf("%d у.е. отправлено от %s к %s\n", sd.Amount, sd.From, sd.To)
}
