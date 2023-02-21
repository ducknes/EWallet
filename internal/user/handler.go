package user

import (
	"encoding/json"
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

type handler struct {
	service *service
}

func NewHandler(service *service) handlers.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(httprouter *httprouter.Router) {
	httprouter.GET(lastTransitionsUrl, h.GetLastTransitions)
	httprouter.GET(balanceUrl, h.GetBalance)
	httprouter.POST(sendMoneyUrl, h.SendMoney)
}

func (h *handler) GetLastTransitions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("запрос последних транзакций")
	count, AtoiErr := strconv.Atoi(r.URL.Query().Get("count"))
	if AtoiErr != nil {
		log.Fatalln(AtoiErr)
	}
	jsonB := h.service.GetUsersTransactions(count)
	if _, err := w.Write(jsonB); err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("запрос на просмотр баланса")
	address := params.ByName("address")
	jsonB := h.service.WatchUserBalance(address)
	if _, err := w.Write(jsonB); err != nil {
		log.Fatalln(err)
	}
}

func (h *handler) SendMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("запрос на отправку денег")
	var send Send
	errJson := json.NewDecoder(r.Body).Decode(&send)
	if errJson != nil {
		log.Fatalln(errJson)
	}
	h.service.PostSendMoney(send)
}
