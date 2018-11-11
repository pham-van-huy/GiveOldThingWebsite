package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"bombay.com/old.thing/services"

	m "bombay.com/old.thing/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(m.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(m.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Registry(w http.ResponseWriter, r *http.Request) {
	user := new(m.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(user)
	user.Birthday = time.Now()
	result := services.Registry(user)
	w.Header().Set("Content-Type", "application/json")

	if result {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"message": "Registry successfuly!",
		}
		message, _ := json.Marshal(data)
		w.Write(message)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		data := map[string]interface{}{
			"message": "Registry fail! Please try again!",
		}
		message, _ := json.Marshal(data)
		w.Write(message)
	}
}
