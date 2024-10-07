package controllers

import (
	"net/http"

	"github.com/pramudyaws/synapsis-restaurant/services"

	"github.com/gin-gonic/gin"
)

func CheckoutFoodCart(c *gin.Context) {
	userID, _ := c.Get("userID")

	orderService := services.NewOrderService()
	if err := orderService.CreateOrder(userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
