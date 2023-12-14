package controller

import (
	"book_api/database"
	"book_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []model.Book
	database.Database.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input model.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := model.Book{Title: input.Title, Author: input.Author}
	database.Database.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}