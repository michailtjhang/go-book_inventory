package main

import (
	apps "book_inventory/app"
	"book_inventory/auth"
	"book_inventory/db"

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
	router.GET("/books", handler.GetBooks)

	router.Run()
}
