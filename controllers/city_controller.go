package controllers

import (
	"bombay.com/old.thing/models"
	"bombay.com/old.thing/services"
	handler "bombay.com/old.thing/tigoweb"
)

type CityHandler struct {
	handler.BaseHandler
}

func (cityHandler *CityHandler) Get() {
	db := services.DB_Instance()
	cities := []models.City{}
	db.Find(&cities)
	cityHandler.ResponseAsJson(WrapSucc(cities))
}
