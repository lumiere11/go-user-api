package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lumiere11/go-user-api/controllers"
	"github.com/lumiere11/go-user-api/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/user/profile", controllers.Profile)

			// Libros relacionados al usuario autenticado
			protected.GET("/books", controllers.GetBooks)
			protected.POST("/books", controllers.CreateBook)
		}
	}
}
