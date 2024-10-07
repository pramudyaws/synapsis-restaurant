package services

import (
	"github.com/pramudyaws/synapsis-restaurant/config"
	"github.com/pramudyaws/synapsis-restaurant/models"
)

type FoodService struct{}

func NewFoodService() *FoodService {
	return &FoodService{}
}

func (s *FoodService) GetFoodList(categoryName string) ([]models.Food, error) {
	var foods []models.Food
	var err error

	if categoryName == "" {
		err = config.DB.Preload("FoodCategory").Find(&foods).Error
	} else {
		err = config.DB.
			Preload("FoodCategory").
			Joins("JOIN food_categories ON food_categories.id = foods.food_category_id").
			Where("food_categories.name = ?", categoryName).
			Find(&foods).Error
	}

	if err != nil {
		return nil, err
	}
	return foods, nil
}
