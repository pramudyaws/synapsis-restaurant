package models

import (
	"time"
	"gorm.io/gorm"
)

type PaymentTransaction struct {
    gorm.Model
    OrderID    uint
    Order      Order 		`gorm:"foreignKey:OrderID"`
    AmountPaid float64
    PaidAt     *time.Time
}
