package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lumiere11/go-user-api/config"
	"github.com/lumiere11/go-user-api/models"
	"github.com/lumiere11/go-user-api/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Book{})
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
