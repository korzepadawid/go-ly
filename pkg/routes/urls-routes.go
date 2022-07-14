package routes

import (
	"github.com/gorilla/mux"
	"github.com/korzepadawid/go-ly/pkg/controllers"
)

func RegisterUrlsRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urls", controllers.CreateNewShortUrl).Methods("POST")
	r.HandleFunc("/{short_url}", controllers.RedirectToShortUrl).Methods("GET")
	return r
}
