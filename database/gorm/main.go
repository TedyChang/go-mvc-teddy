package main

import (
	"codetest/board/entity"
	entity2 "codetest/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost port=15432 user=postgres dbname=postgres password=postgres sslmode=disable",
	}))

	if err != nil {
		log.Panicf("err : %v", err)
		return
	}

	// Migrate the schema
	err = db.AutoMigrate(&entity.Board{})
	if err != nil {
		log.Panicf("err : %v", err)
		return
	}
	err = db.AutoMigrate(&entity.Reply{})
	if err != nil {
		log.Panicf("err : %v", err)
		return
	}
	err = db.AutoMigrate(&entity2.User{})
	if err != nil {
		log.Panicf("err : %v", err)
		return
	}
}
