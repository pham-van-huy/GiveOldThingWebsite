package controllers

import (
	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
	handler "bombay.com/old.thing/tigoweb"
)

type CateHandler struct {
	handler.BaseHandler
}

func (cateHandler *CateHandler) Get() {
	db := services.DB_Instance()
	categories := []models.Category{}
	db.Find(&categories)
	cateHandler.ResponseAsJson(WrapSucc(categories))
}
