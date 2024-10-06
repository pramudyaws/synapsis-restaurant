package main

import (
	"os"

	"github.com/pramudyaws/synapsis-restaurant/config"
	"github.com/pramudyaws/synapsis-restaurant/models"
	"github.com/pramudyaws/synapsis-restaurant/routes"
)

func main() {
	// Connect to DB
	config.ConnectDatabase()
	// Auto-migrate all models
	config.DB.AutoMigrate(
		&models.User{},
		&models.FoodCategory{}, &models.Food{}, &models.FoodCart{},
		&models.Order{}, &models.OrderItem{},
		&models.PaymentTransaction{},
	)
	// Setup routes
	r := routes.SetupRouter()
	// Run server
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
