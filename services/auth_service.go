package services

import (
	"errors"
	"os"
	"time"

	"github.com/pramudyaws/synapsis-restaurant/config"
	"github.com/pramudyaws/synapsis-restaurant/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) GenerateJWT(userId uint) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("email not found")
	}

	if err := s.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("incorrect password")
	}

	// Update user.LastLoginAt
	now := time.Now()
	user.LastLoginAt = &now
	if err := config.DB.Save(&user).Error; err != nil {
		return "", errors.New("failed to update last login time")
	}

	token, err := s.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
