package main

import (
	"log"
	"net/http"

	"bombay.com/old.thing/routers"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
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
