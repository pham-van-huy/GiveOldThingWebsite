package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	m "old.thing/databases"
)

func main() {

	app := cli.NewApp()
	app.Name = "migrate"
	app.Usage = "migrate database !"
	app.Action = func(c *cli.Context) error {
		m.Init()
		fmt.Println("Done !!!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
