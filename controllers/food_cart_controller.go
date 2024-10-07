package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pramudyaws/synapsis-restaurant/services"
)

type AddToFoodCartRequest struct {
	FoodID   uint `json:"food_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}

func AddToFoodCart(c *gin.Context) {
	var request AddToFoodCartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	foodCartService := services.NewFoodCartService()
	if err := foodCartService.AddToCart(userID.(uint), request.FoodID, request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food added to cart successfully"})
}

func DeleteFromFoodCart(c *gin.Context) {
	foodCartID := c.Param("id")

	foodCartService := services.NewFoodCartService()
	if err := foodCartService.DeleteFromCart(foodCartID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food removed from cart successfully"})
}

func GetFoodCartList(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	foodCartService := services.NewFoodCartService()
	foodCartItems, err := foodCartService.GetCartItemsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foodCartItems)
}
