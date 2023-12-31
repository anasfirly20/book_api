package main

import (
	"book_api/controller"
	"book_api/database"
	"book_api/model"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Book{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))
	
	router.GET("/books", controller.GetBooks)
	router.POST("/books", controller.CreateBook)
	router.GET("/books/:id", controller.GetBookById)
	router.PATCH("books/:id", controller.UpdateBook)
	router.DELETE("books/:id", controller.DeleteBook)
	
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}