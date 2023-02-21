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

	repos := user.NewRepo(database)
	service := user.NewService(repos)
	handler := user.NewHandler(service)
	handler.Register(httprouter)

	var count int
	if err := database.QueryRow("SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		log.Println(err)
	}
	if count == 0 {
		var u user.User
		for i := 0; i < 10; i++ {
			AddErr := repos.AddNewUser(u.NewUser())
			if AddErr != nil {
				log.Println(AddErr)
			}
		}
	}

	log.Println("listening server on localhost")
	log.Fatalln(server.Serve(listener))
}
