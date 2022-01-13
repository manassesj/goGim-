package main

import (
	"github.com/gin-gonic/gin"
	"goGim/controllers"
	"goGim/models"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
