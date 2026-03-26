package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authors = make(map[int]models.Author)
var nextAuthorID = 1

func GetAuthors(c *gin.Context) {
	var authorList []models.Author
	for _, author := range authors {
		authorList = append(authorList, author)
	}
	c.JSON(http.StatusOK, authorList)
}

func AddAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.ID = nextAuthorID
	nextAuthorID++
	authors[author.ID] = author
	c.JSON(http.StatusCreated, author)
}
