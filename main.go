package main

import (
	"codetest/database"
	"codetest/rest/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	domain := url + ":" + port
	database.GormSetting()

	api.RestApi(domain)
}
