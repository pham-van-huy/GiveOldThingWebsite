package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"old.thing/routers"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8001", n)
}
