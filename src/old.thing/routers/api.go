package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"old.thing/controllers"
)

func handleApi(router *mux.Router) *mux.Router {
	apiRouter := mux.NewRouter()
	subRouter := apiRouter.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/cities", controllers.CityList)

	subRouter.HandleFunc("/posts", controllers.PostList)
	subRouter.HandleFunc("/posts/create", controllers.PostCreate).Methods("POST")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostUpdate).Methods("PUT")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostDelete).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(negroni.HandlerFunc(RequireTokenAuthentication), negroni.Wrap(apiRouter)))

	return router
}
