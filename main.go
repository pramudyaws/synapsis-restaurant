package main

import (
	"synapsis-restaurant/config"
	"synapsis-restaurant/models"
	"synapsis-restaurant/routes"
)

func main() {
    // Connect to DB
	config.ConnectDatabase()
    // Auto-migrate all models
	config.DB.AutoMigrate(&models.User{})
    // Setup routes
	r := routes.SetupRouter()
    // Run server
	r.Run()
}
