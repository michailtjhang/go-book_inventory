package main

import (
	apps "book_inventory/app"
	"book_inventory/auth"
	"book_inventory/db"
	"book_inventory/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := apps.New(conn)

	// Home
	router.GET("/", auth.HomeHandler)

	// Login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// Get All Books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/books/:id", middleware.AuthValid, handler.GetBookByID)

	// Add Books
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.SaveBook)

	// Update Books
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutUpdateBook)

	// Delete Books
	router.GET("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
