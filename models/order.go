package models

import (
	"gorm.io/gorm"
)

type Order struct {
    gorm.Model
    UserID      uint
    User        User   `gorm:"foreignKey:UserID"`
    TotalAmount float64
}
