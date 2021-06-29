package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
