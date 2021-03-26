package main

import (
	"boj-garden/server"
	"log"
)

func main() {
	app := server.Init()

	err := app.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
