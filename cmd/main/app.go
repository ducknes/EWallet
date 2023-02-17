package main

import (
	"infotecs-EWallet/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("creating router...")
	router := httprouter.New()

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

	database, DBerr := user.DatabaseConnection()
	if DBerr != nil {
		panic(DBerr)
	}

	var count int
	if err := database.QueryRow("SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		panic(err)
	}
	if count == 0 {
		var u user.User
		for i := 0; i < 10; i++ {
			AddErr := user.AddNewUser(database, u.NewUser())
			if AddErr != nil {
				panic(AddErr)
			}
		}
	}

	handler := user.NewHandler()
	handler.Register(httprouter)

	log.Println("listening server on localhost")
	log.Fatalln(server.Serve(listener))
}
