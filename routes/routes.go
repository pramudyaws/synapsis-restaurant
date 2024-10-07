package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pramudyaws/synapsis-restaurant/controllers"
	"github.com/pramudyaws/synapsis-restaurant/middlewares"
)

// Public routes
func setupPublicRoutes(apiV1 *gin.RouterGroup) {
	public := apiV1.Group("/public")
	{
		public.POST("/register", controllers.CreateUser)
		public.POST("/login", controllers.Login)
		public.GET("/foods", controllers.GetFoodList)
	}
}

// Private routes
func setupPrivateRoutes(apiV1 *gin.RouterGroup) {
	private := apiV1.Group("/private")
	private.Use(middlewares.JWTAuthMiddleware())
	{
		private.POST("/food-carts", controllers.AddToFoodCart)
		private.DELETE("/food-carts/:id", controllers.DeleteFromFoodCart)
		private.GET("/food-carts", controllers.GetFoodCartList)
		private.POST("/food-carts/checkout", controllers.CheckoutFoodCart)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API V1
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to Synapsis Restaurant API",
			})
		})
		setupPublicRoutes(apiV1)
		setupPrivateRoutes(apiV1)
	}

	return r
}
