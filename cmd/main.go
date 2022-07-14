package main

import (
	"log"
	"net/http"
	"time"

	"github.com/korzepadawid/go-ly/pkg/config"
	"github.com/korzepadawid/go-ly/pkg/routes"
)

const (
	Port                  = ":8080"
	ReadTimeoutInSeconds  = 10
	WriteTimeoutInSeconds = 10
)

func main() {
	log.Println("Starting server.")

	config.InitRedis()

	router := routes.RegisterUrlsRoutes()

	srv := &http.Server{
		Addr:         Port,
		Handler:      router,
		ReadTimeout:  ReadTimeoutInSeconds * time.Second,
		WriteTimeout: WriteTimeoutInSeconds * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	log.Println("Listening at http://localhost", Port)
}
