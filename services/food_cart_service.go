package services

import (
	"synapsis-restaurant/models"
	"synapsis-restaurant/config"
)

type FoodCartService struct{}

func NewFoodCartService() *FoodCartService {
	return &FoodCartService{}
}

func (s *FoodCartService) AddToCart(userID uint, foodID uint, quantity int) error {
    var foodCart models.FoodCart

    // Check if food cart has been created before
    if err := config.DB.Where("user_id = ? AND food_id = ?", userID, foodID).First(&foodCart).Error; err == nil {
        // If food cart is found, update quantity
        foodCart.Quantity += quantity
        if err := config.DB.Save(&foodCart).Error; err != nil {
            return err
        }
        return nil
    }

    // If food cart not found, create new food cart
    foodCart = models.FoodCart{
        UserID:   userID,
        FoodID:   foodID,
        Quantity: quantity,
    }

    if err := config.DB.Create(&foodCart).Error; err != nil {
        return err
    }

    return nil
}


func (s *FoodCartService) DeleteFromCart(foodCartID string) error {
	var foodCart models.FoodCart
	if err := config.DB.Where("id = ?", foodCartID).Delete(&foodCart).Error; err != nil {
		return err
	}

	return nil
}

func (s *FoodCartService) GetCartItemsByUserID(userID uint) ([]models.FoodCart, error) {
    var foodCarts []models.FoodCart
    if err := config.DB.Preload("Food").Preload("User").Where("user_id = ?", userID).Find(&foodCarts).Error; err != nil {
        return nil, err
    }

    return foodCarts, nil
}

func (s *FoodCartService) DeleteFoodCart(userID uint) error {
	return config.DB.Where("user_id = ?", userID).Delete(&models.FoodCart{}).Error
}
