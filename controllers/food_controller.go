package controllers

import (
	"net/http"

	"github.com/pramudyaws/synapsis-restaurant/services"

	"github.com/gin-gonic/gin"
)

func GetFoodList(c *gin.Context) {
	categoryName := c.Query("category")

	foodService := services.NewFoodService()
	foods, err := foodService.GetFoodList(categoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}
