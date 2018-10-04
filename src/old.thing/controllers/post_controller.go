package controllers

import (
	"encoding/json"
	"net/http"

	"old.thing/models"
	"old.thing/services"
)

type response struct {
	Success bool          `json:"success"`
	Data    []models.Post `json:"data"`
}

func PostList(w http.ResponseWriter, r *http.Request) {
	db := services.DB_Open()
	posts := []models.Post{}
	db.Preload("Category").Limit(10).Find(&posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	res := response{}

	res.Success = true
	res.Data = posts

	result, _ := json.Marshal(res)

	w.Write(result)
}
