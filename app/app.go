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
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) GetBookByID(c *gin.Context) {
	BookID := c.Param("id")

	var books models.Book

	if h.DB.Find(&books, "id=?", BookID).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "Book.html", gin.H{
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
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var book models.Book

	bookID := c.Param("id")
	if h.DB.Find(&book, "id=?", bookID).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Book not found",
		})
	}

	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title":   "Update Books",
		"payload": book,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) PutUpdateBook(c *gin.Context) {
	var book models.Book

	bookID := c.Param("id")
	if h.DB.Find(&book, "id=?", bookID).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Book not found",
		})
	}

	var reqBook = book
	c.Bind(&reqBook)

	h.DB.Model(&book).Where("id=?", bookID).Update(&reqBook)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/book/%s?auth=%s", bookID, c.PostForm("auth")))
}

func (h *Handler) DeleteBook(c *gin.Context) {
	var book models.Book

	bookID := c.Param("id")
	h.DB.Delete(&book, "id=?", bookID)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}
