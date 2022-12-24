package main

import (
	"access-management-system/config"
	"access-management-system/models"
	"access-management-system/routers"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	r := routers.Router()

	models.ConnectDB(&config)

	models.MigrateTables()

	models.SeedData()

	r.Run() // listen and serve on 0.0.0.0:8080
}
