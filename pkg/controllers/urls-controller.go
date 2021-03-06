package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/korzepadawid/go-ly/pkg/config"
	"github.com/korzepadawid/go-ly/pkg/dto"
	"github.com/korzepadawid/go-ly/pkg/model"
	"github.com/korzepadawid/go-ly/pkg/util"
)

func CreateNewShortUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new short url")

	requestDTO := dto.URLRequestDTO{}
	json.NewDecoder(r.Body).Decode(&requestDTO)
	defer r.Body.Close()

	url := requestDTO.Map()
	url.AuthorIPv4 = r.RemoteAddr

	url = url.SaveURL()
	base62, err := util.Base62().Encode(int64(url.ID - 1)) // -1 cuz database id starts from 1, base62 starts from 0

	if err != nil {
		log.Printf("can't convert %d\n", url.ID)
	}

	redis := config.Redis
	redis.Set(base62, url.Url, 1*time.Hour)

	var responseDTO dto.URLResponseDTO
	responseDTO.Map(url, base62)
	body, err := json.Marshal(responseDTO)

	if err != nil {
		log.Printf("can't convert %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func RedirectToShortUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting to short url from ", r.RemoteAddr)

	vars := mux.Vars(r)
	redis := config.Redis

	value := redis.Get(vars["base62"]).Val()

	if len(value) > 0 {
		http.Redirect(w, r, value, http.StatusPermanentRedirect)
		return
	}

	ID, err := util.Base62().Decode(value)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, tx := model.GetURLByID(ID + 1)

	if tx.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusPermanentRedirect)
}
