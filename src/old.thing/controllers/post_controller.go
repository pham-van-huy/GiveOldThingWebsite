package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// PostUpdate update post
func PostUpdate(w http.ResponseWriter, r *http.Request) {

	postParams := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&postParams)
	if err != nil {
		panic(err)
	}

	request := NewPostRequest{Title: postParams.Title, Description: postParams.Description, UserId: postParams.UserId, CategoryId: postParams.CategoryId}

	if errs := validator.Validate(request); errs != nil {
		ResErrors(w, errs)
	}

	id := mux.Vars(r)["id"]
	post := models.Post{}
	db := services.DB_Instance()
	db.First(&post, id)

	db.Model(&post).Updates(postParams)

	ResSuccess(w, post)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	db := services.DB_Instance()
	posts := []models.Post{}
	db.Preload("Category").Limit(10).Find(&posts)

	ResSuccess(w, posts)
}

func PostEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	post := models.Post{}
	db := services.DB_Instance()
	db.First(&post, id)
	ResSuccess(w, post)
}

func PostDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	post := models.Post{}
	db := services.DB_Instance()
	db.First(&post, id)
	db.Delete(&post)

	ResSuccess(w, post)
}
