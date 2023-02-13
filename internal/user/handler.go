package user

import (
	"infotecs-EWallet/internal/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	lastTransitionsUrl = "/api/transactions"
	balanceUrl         = "/api/wallet/{address}/balance"
	sendMoneyUrl       = "/api/send"
)

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
	_, err := w.Write([]byte("get last transitions\n"))
	if err != nil {
		log.Fatalln(err)
	}

	_, err = w.Write([]byte(r.URL.Query().Get("count")))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("get balance"))
}

func (h *handler) SendMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("send money"))
}
