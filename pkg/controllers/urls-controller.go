package controllers

import (
	"log"
	"net/http"
)

func CreateNewShortUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new short url")
}

func RedirectToShortUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting to short url from ", r.RemoteAddr)
}
