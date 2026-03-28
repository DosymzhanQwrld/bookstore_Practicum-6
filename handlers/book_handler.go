package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = make(map[int]models.Book)
var nextID = 1

func GetBooks(c *gin.Context) {

	var bookList []models.Book

	category := c.Query("category")

	for _, book := range books {
		if category != "" {
			if strconv.Itoa(book.CategoryID) == category {
				continue
			}
		}
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, bookList)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	book.ID = nextID
	nextID++
	books[book.ID] = book
	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	book, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	updatedBook.ID = id
	books[id] = updatedBook
	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	delete(books, id)
	c.Status(http.StatusNoContent)
}
