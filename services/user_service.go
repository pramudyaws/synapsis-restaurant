package services

import (
    "synapsis-restaurant/config"
    "synapsis-restaurant/models"
    "log"

    "golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func (s *UserService) CreateUser(input models.User) (models.User, error) {
    hashedPassword, err := hashPassword(input.Password)
    if err != nil {
        log.Println("Error hashing password:", err)
        return models.User{}, err
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: hashedPassword,
    }

    if err := config.DB.Create(&user).Error; err != nil {
        log.Println("Error creating user:", err)
        return models.User{}, err
    }

    return user, nil
}
