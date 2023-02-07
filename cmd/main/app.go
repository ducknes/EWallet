package main

import (
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

	log.Println("listening server on localhost")
	log.Fatalln(server.Serve(listener))
}
