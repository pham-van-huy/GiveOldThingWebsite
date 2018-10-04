package controllers

import (
	"encoding/json"
	"net/http"
)

type payload interface {
}

type response struct {
	Success bool    `json:"success"`
	Data    payload `json:"data"`
}

func ResSuccess(w http.ResponseWriter, data payload) {
	res := response{}

	res.Success = true
	res.Data = data

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func ResErrors(w http.ResponseWriter, data payload) {
	res := response{}

	res.Success = false
	res.Data = data

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}
