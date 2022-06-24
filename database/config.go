package database

import (
	"codetest/board/entity"
	entity2 "codetest/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func GormSetting() {
	var err error
	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost port=15432 user=postgres dbname=postgres password=postgres sslmode=disable",
	}))
	if err != nil {
		log.Panicf("gorm setting err : %v", err)
		return
	}

	err = Db.AutoMigrate(&entity.Board{})
	if err != nil {
		log.Panicf("board ddl err : %v", err)
		return
	}

	err = Db.AutoMigrate(&entity.Reply{})
	if err != nil {
		log.Panicf("err : %v", err)
		return
	}

	err = Db.AutoMigrate(&entity2.User{})
	if err != nil {
		log.Panicf("err : %v", err)
		return
	}
}
