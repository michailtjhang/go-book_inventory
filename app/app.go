package app

import (
	"book_inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) GetBooks(c *gin.Context) {
	var books []models.Book

	h.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": books,
	})
}

func (h *Handler) GetBookByID(c *gin.Context) {
	BookID := c.Param("id")

	var books models.Book

	if h.DB.Find(&books, "id=?", BookID).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "book.html", gin.H{
		"title":   books.Title,
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) AddBook(c *gin.Context) {
	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title": "Add Books",
		"auth":  c.Query("auth"),
	})
}

func (h *Handler) SaveBook(c *gin.Context) {
	var book models.Book

	c.Bind(&book)
	h.DB.Create(&book)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.Query("auth")))
}
