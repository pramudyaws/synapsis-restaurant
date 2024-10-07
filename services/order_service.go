package services

import (
	"errors"

	"github.com/pramudyaws/synapsis-restaurant/config"
	"github.com/pramudyaws/synapsis-restaurant/models"
)

type OrderService struct {
	foodCartService 				*FoodCartService
	paymentTransactionService    	*PaymentTransactionService
}

func NewOrderService() *OrderService {
	return &OrderService{
		foodCartService: 			NewFoodCartService(),
		paymentTransactionService: 	NewPaymentTransactionService(),
	}
}

func (s *OrderService) CreateOrder(userID uint) error {
	var foodCarts []models.FoodCart
	if err := config.DB.Where("user_id = ?", userID).Find(&foodCarts).Error; err != nil {
		return err
	}

	// If no food cart, return error
	if len(foodCarts) == 0 {
		return errors.New("no items in food cart")
	}

	// Calculate total amount
	totalAmount := 0.0
	var orderItems []models.OrderItem

	for _, cart := range foodCarts {
		var food models.Food
		if err := config.DB.First(&food, cart.FoodID).Error; err != nil {
			return err
		}

		totalAmount += food.Price * float64(cart.Quantity)

		orderItems = append(orderItems, models.OrderItem{
			FoodID:   food.ID,
			Price:    food.Price,
			Quantity: cart.Quantity,
		})
	}

	// Save order
	order := models.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return err
	}

	// Save order items
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}

	if err := config.DB.Create(&orderItems).Error; err != nil {
		return err
	}

	// Save payment transaction
	if err := s.paymentTransactionService.SavePaymentTransaction(order.ID, totalAmount); err != nil {
		return err
	}

	// Delete user's food carts
	if err := s.foodCartService.DeleteFoodCart(userID); err != nil {
		return err
	}

	return nil
}
