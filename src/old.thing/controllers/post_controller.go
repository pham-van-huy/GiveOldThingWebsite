package controllers

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/validator.v2"
	"old.thing/models"
	"old.thing/services"
)

type NewPostRequest struct {
	Title       string `validate:"min=3,max=100"`
	Description string `validate:"min=3,max=1000"`
	UserId      int    `validate:nonzero`
	CategoryId  int    `validate:nonzero`
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		panic(err)
	}

	nur := NewPostRequest{Title: post.Title, Description: post.Description, UserId: post.UserId, CategoryId: post.CategoryId}

	if errs := validator.Validate(nur); errs != nil {
		ResErrors(w, errs)
	}
	db := services.DB_Instance()
	db.Create(&post)
	ResSuccess(w, post)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	db := services.DB_Instance()
	posts := []models.Post{}
	db.Preload("Category").Limit(10).Find(&posts)

	ResSuccess(w, posts)
}
