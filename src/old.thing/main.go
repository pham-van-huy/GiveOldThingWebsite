package main

import (
	"fmt"
	"net/http"

	"old.thing/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/negroni"
)

func main() {
	router := routers.InitRoutes()

	router.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		fmt.Fprintf(w, "Welcome to the home page!")

	})

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8001", n)

}
