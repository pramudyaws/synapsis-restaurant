package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email       string    `gorm:"uniqueIndex"`
    Password    string
    Name        string
    LastLoginAt *time.Time
}
