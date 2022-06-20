package main

import (
	"codetest/database"
	"codetest/rest/gin"
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

	gin.RestApi(domain)
}
