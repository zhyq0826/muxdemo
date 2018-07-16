package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/zhyq0826/muxdemo/router"
)

func main() {
	r := router.InitRouter()
	handler := handlers.LoggingHandler(os.Stdout, r)
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: handler,
	}
	log.Println("Listening...")
	// start run http server
	server.ListenAndServe()
}
