package controllers

import (
	"net/http"
	"synapsis-restaurant/models"
	"synapsis-restaurant/services"

	"github.com/gin-gonic/gin"
)

var userService = services.UserService{}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
