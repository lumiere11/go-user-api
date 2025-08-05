package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lumiere11/go-user-api/config"
	"github.com/lumiere11/go-user-api/models"
	"github.com/lumiere11/go-user-api/utils"
)

type BookInput struct {
	Title  string `json:"title" validate:"required,min=2"`
	Author string `json:"author" validate:"required"`
}

func CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validar campos
	if err := utils.Validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		UserID: user.ID,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el libro"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)

	var books []models.Book
	if err := config.DB.Where("user_id = ?", user.ID).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros"})
		return
	}

	c.JSON(http.StatusOK, books)
}
