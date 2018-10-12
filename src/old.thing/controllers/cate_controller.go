package controllers

import (
	"net/http"

	"old.thing/models"
	"old.thing/services"
)

// PostCreate create post
func CateList(w http.ResponseWriter, r *http.Request) {
	db := services.DB_Instance()
	catgories := []models.Category{}
	db.Find(&catgories)
	ResSuccess(w, catgories)
}
