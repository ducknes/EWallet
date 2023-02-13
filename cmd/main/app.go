package main

import (
	"fmt"
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

	fmt.Println(database)

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
