package models

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name           string 		`gorm:"unique"`
	Price          float64
	FoodCategoryID uint
	FoodCategory   FoodCategory `gorm:"foreignKey:FoodCategoryID"`
}
