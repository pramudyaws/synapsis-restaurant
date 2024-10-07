package models

import (
	"gorm.io/gorm"
)

type FoodCart struct {
    gorm.Model
    UserID   	uint
    User     	User	`gorm:"foreignKey:UserID"`
    FoodID 		uint
    Food		Food	`gorm:"foreignKey:FoodID"`
    Quantity 	int
}
