package routers

import (
	"github.com/gorilla/mux"
	"old.thing/controllers"
)

func handleApi(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Home).Methods("GET")

	return router
}
