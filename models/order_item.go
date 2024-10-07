package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
    gorm.Model
    OrderID   	uint
    Order     	Order   `gorm:"foreignKey:OrderID"`
    FoodID 		uint
    Food   		Food 	`gorm:"foreignKey:FoodID"`
    Price     	float64
    Quantity  	int
}
