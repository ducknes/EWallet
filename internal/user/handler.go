package user

import (
	"infotecs-EWallet/internal/handlers"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	lastTransitionsUrl = "/api/transactions"
	balanceUrl         = "/api/wallet/:address/balance"
	sendMoneyUrl       = "/api/send"
)

var s *service

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(httprouter *httprouter.Router) {
	httprouter.GET(lastTransitionsUrl, h.GetLastTransitions)
	httprouter.GET(balanceUrl, h.GetBalance)
	httprouter.POST(sendMoneyUrl, h.SendMoney)
}

func (h *handler) GetLastTransitions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	count, AtoiErr := strconv.Atoi(r.URL.Query().Get("count"))
	if AtoiErr != nil {
		log.Fatalln(AtoiErr)
	}
	jsonB := s.GetUsersTransactions(count)
	if _, err := w.Write(jsonB); err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	address := params.ByName("address")
	jsonB := s.WatchUserBalance(address)
	if _, err := w.Write(jsonB); err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) SendMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("send money"))
}
