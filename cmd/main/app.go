package main

import (
	"infotecs-EWallet/internal/user"
	"infotecs-EWallet/internal/user/db/sqlite"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("creating router...")
	router := httprouter.New()

	database, DBerr := sqlite.DatabaseConnection()
	if DBerr != nil {
		panic(DBerr)
	}

	var u user.User

	for i := 0; i < 10; i++ {
		AddErr := sqlite.AddNewUser(database, u.NewUser())
		if AddErr != nil {
			panic(AddErr)
		}

	}

	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(httprouter *httprouter.Router) {
	log.Println("starting application...")

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatalln(err)
	}

	server := &http.Server{
		Handler:      httprouter,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("listening server on localhost")
	log.Fatalln(server.Serve(listener))
}
