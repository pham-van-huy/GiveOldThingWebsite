package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"old.thing/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8001", n)
}
