package main

import (
	"log"

	"rexlists/app"
)

func main() {
	log.Println("Running front end")
	app := app.New()
	log.Fatal(app.Start())
}
