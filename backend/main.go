package main

import (
	"boj-garden/config"
	"boj-garden/server"
	"boj-garden/utils"
	"log"
)

func main() {
	cfg := config.Load()
	app := server.Init(cfg)
	utils.MigrateDatabase(app.DB)
	err := app.Run(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
