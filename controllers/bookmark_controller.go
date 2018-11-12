package controllers

import (
	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
	handler "bombay.com/old.thing/tigoweb"
)

type BookmarkHandler struct {
	handler.BaseHandler
}

func (bookmarkHandler *BookmarkHandler) Create() {
	db := services.DB_Instance()
	bookmarks := models.BookMark{}
	db.Find(&bookmarks)
	bookmarkHandler.ResponseAsJson(WrapSucc(bookmarks))
}
