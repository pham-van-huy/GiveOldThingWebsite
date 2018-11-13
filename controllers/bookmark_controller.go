package controllers

import (
	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
	handler "bombay.com/old.thing/tigoweb"
)

type BookmarkHandler struct {
	handler.BaseHandler
}

func (bookmarkHandler *BookmarkHandler) Post() {
	db := services.DB_Instance()
	bookmark := models.BookMark{}
	bookmark.UserId = bookmarkHandler.GetCtxVal("UserId").ToInt()
	bookmark.PostId = bookmarkHandler.GetParameter("PostId").ToInt()
	db.Create(&bookmark)

	bookmarkHandler.ResponseAsJson(WrapSucc(bookmark))
}
