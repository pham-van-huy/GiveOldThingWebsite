package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"old.thing/routers"
)

func main() {

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8001", n)
}
