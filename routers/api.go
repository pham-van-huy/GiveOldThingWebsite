package routers

import (
	"bombay.com/old.thing/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleApi(router *mux.Router) *mux.Router {
	router.HandleFunc("/web/cities", controllers.CityList)
	router.HandleFunc("/web/categories", controllers.CateList)

	apiRouter := mux.NewRouter()
	subRouter := apiRouter.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/posts", controllers.PostList)
	subRouter.HandleFunc("/posts/create", controllers.PostCreate).Methods("POST")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostEdit).Methods("GET")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostUpdate).Methods("PUT")
	subRouter.HandleFunc("/posts/{id:[0-9]+}", controllers.PostDelete).Methods("DELETE")
	router.PathPrefix("/api").Handler(negroni.New(negroni.HandlerFunc(RequireTokenAuthentication), negroni.Wrap(apiRouter)))

	return router
}
