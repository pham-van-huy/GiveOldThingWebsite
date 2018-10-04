package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"old.thing/controllers"
)

func handleApi(router *mux.Router) *mux.Router {
	apiRouter := mux.NewRouter()
	subRouter := apiRouter.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/posts", controllers.PostList)
	router.PathPrefix("/api").Handler(negroni.New(negroni.HandlerFunc(RequireTokenAuthentication), negroni.Wrap(apiRouter)))

	return router
}
