package database

import (
	"codetest/board/entity"
	entity2 "codetest/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

func GormSetting() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost port=15432 user=postgres dbname=postgres password=postgres sslmode=disable",
	}))
	Db = db
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Board{})
	db.AutoMigrate(&entity.Reply{})
	db.AutoMigrate(&entity2.User{})

	db.Create(&entity2.User{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Name:      "teddy",
	})

}
