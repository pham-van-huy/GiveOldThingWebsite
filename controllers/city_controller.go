package controllers

import (
	"net/http"

	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
)

// PostCreate create post
func CityList(w http.ResponseWriter, r *http.Request) {
	db := services.DB_Instance()
	cities := []models.City{}
	db.Find(&cities)
	ResSuccess(w, cities)
}
