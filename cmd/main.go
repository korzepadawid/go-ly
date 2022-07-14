package main

import (
	"log"
	"net/http"
	"time"

	"github.com/korzepadawid/go-ly/pkg/routes"
	"github.com/korzepadawid/go-ly/pkg/util"
)

func main() {
	log.Println("Starting server.")

	util.InitRedis()

	router := routes.RegisterUrlsRoutes()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
