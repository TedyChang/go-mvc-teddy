package entity

import (
	"gorm.io/gorm"
	"time"
)

type Rows struct {
	Rows []User
}

type User struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}
