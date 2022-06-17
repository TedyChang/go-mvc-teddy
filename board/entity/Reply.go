package entity

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BoardID   int64
	Board     Board `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content   string
}
