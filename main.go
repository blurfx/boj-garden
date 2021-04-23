package main

import (
	"boj-garden/config"
	"boj-garden/server"
	"boj-garden/server/routes"
	"boj-garden/utils"
	"log"
)

func main() {
	cfg := config.Load()
	app := server.Init(cfg)

	routes.InitRouter(app)
	utils.MigrateDatabase(app.DB)
	utils.RunCrontab(app.DB)

	if err := app.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
