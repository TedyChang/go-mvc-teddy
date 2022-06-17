package entity

import (
	"gorm.io/gorm"
	"time"
)

type Rows struct {
	BoardRows []Board
	ReplyRows []Reply
}

type Board struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string
	Contents  string
}
