package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	paginator "bombay.com/old.thing/controllers/pagination"
	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	validator "gopkg.in/validator.v2"
)

type NewPostRequest struct {
	Title       string `validate:"min=3,max=100"`
	Description string `validate:"min=3,max=1000"`
	UserId      int    `validate:nonzero`
	CategoryId  int    `validate:nonzero`
}

var decoder *form.Decoder

// PostCreate create post
func PostCreate(w http.ResponseWriter, r *http.Request) {
	decoder = form.NewDecoder()
	r.ParseMultipartForm(32 << 20)
	post := models.Post{}
	err := decoder.Decode(&post, r.PostForm)
	if err == nil {
		// handle error
		nur := NewPostRequest{Title: post.Title, Description: post.Description, UserId: post.UserId, CategoryId: post.CategoryId}

		if errs := validator.Validate(nur); errs != nil {
			ResErrors(w, errs)
		}
		db := services.DB_Instance()
		db.Create(&post)

		images := r.MultipartForm.File["Images[]"]
		if images != nil {
			for _, fileHeader := range images {
				file, _ := fileHeader.Open()
				arrFilename := strings.Split(fileHeader.Filename, ".")
				extension := arrFilename[len(arrFilename)-1]
				fileName := RandSeq(20) + "." + extension
				f, err := os.OpenFile("uploads/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer f.Close()

				io.Copy(f, file)

				mImage := models.Image{}
				mImage.Name = fileHeader.Filename
				mImage.Description = ""
				mImage.Link = fileName
				mImage.OwnerId = post.ID
				mImage.OwnerType = models.TYPE_POST
				db.Create(&mImage)
			}
		}
		ResSuccess(w, post)
	}
}

// PostUpdate update post
func PostUpdate(w http.ResponseWriter, r *http.Request) {
	decoder = form.NewDecoder()
	db := services.DB_Instance()
	r.ParseMultipartForm(32 << 20)
	post := models.Post{}
	err := decoder.Decode(&post, r.PostForm)
	if err != nil {
		panic(err)
	}
	images := r.MultipartForm.File["Images[]"]
	if images != nil {
		for _, fileHeader := range images {
			file, _ := fileHeader.Open()
			arrFilename := strings.Split(fileHeader.Filename, ".")
			extension := arrFilename[len(arrFilename)-1]
			fileName := RandSeq(20) + "." + extension
			f, err := os.OpenFile("uploads/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			io.Copy(f, file)
			mImage := models.Image{}
			mImage.Name = fileHeader.Filename
			mImage.Description = ""
			mImage.Link = fileName
			mImage.OwnerId = post.ID
			mImage.OwnerType = models.TYPE_POST
			db.Create(&mImage)
		}
	}
	removeLinks := r.Form["RemoveImages[]"]

	if len(removeLinks) != 0 {
		for _, link := range removeLinks {
			path := "uploads/" + link
			os.Remove(path)
			// if err != nil {
			// 	return
			// }
			fmt.Println("==> done deleting file")
			db.Where("link = ?", link).Delete(models.Image{})
		}
	}

	request := NewPostRequest{Title: post.Title, Description: post.Description, UserId: post.UserId, CategoryId: post.CategoryId}

	if errs := validator.Validate(request); errs != nil {
		ResErrors(w, errs)
	}

	id := mux.Vars(r)["id"]
	inpost := models.Post{}
	db.First(&inpost, id)

	db.Model(&inpost).Updates(post)

	ResSuccess(w, inpost)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	limit, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		page = 1
	}
	if err1 == nil {
		limit = 10
	}
	db := services.DB_Instance()
	posts := []models.Post{}
	paginator := paginator.Pagging(&paginator.Param{
		DB:      db.Preload("Category").Preload("Images", "owner_type = ?", models.TYPE_POST).Preload("City"),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"ID desc"},
		ShowSQL: true,
	}, &posts)

	ResSuccess(w, paginator)
}

func PostEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	post := models.Post{}
	db := services.DB_Instance()
	db.Preload("Category").Preload("Images", "owner_type = ?", models.TYPE_POST).Preload("City").First(&post, id)
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
