package main

import (
	"log"
	"net/http"

	"bombay.com/old.thing/routers"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
)

func main() {
	// app := cli.NewApp()
	// app.Name = "migrate"
	// app.Usage = "migrate database !"
	// app.Action = func(c *cli.Context) error {
	// 	m.Init()
	// 	fmt.Println("Done !!!")
	// 	return nil
	// }

	// err1 := app.Run(os.Args)
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8001", n)
}
