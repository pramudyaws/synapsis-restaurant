package routes

import (
	"synapsis-restaurant/controllers"
	"synapsis-restaurant/middlewares"
	"github.com/gin-gonic/gin"
)

func setupPublicRoutes(apiV1 *gin.RouterGroup) {
	public := apiV1.Group("/public")
	{
		// Public routes
		public.POST("/register", controllers.CreateUser)
		public.POST("/login", controllers.Login)
	}
}

func setupPrivateRoutes(apiV1 *gin.RouterGroup) {
	private := apiV1.Group("/private")
	private.Use(middlewares.JWTAuthMiddleware())
	{
		// Private routes
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API V1
	apiV1 := r.Group("/api/v1")
	{
		setupPublicRoutes(apiV1)
		setupPrivateRoutes(apiV1)
	}

	return r
}
