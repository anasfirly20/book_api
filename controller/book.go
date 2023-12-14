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