package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = make(map[int]models.Category)
var nextCategoryID = 1

func GetCategories(c *gin.Context) {
	var categoryList []models.Category
	for _, category := range categories {
		categoryList = append(categoryList, category)
	}
	c.JSON(http.StatusOK, categoryList)
}

func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.ID = nextCategoryID
	nextCategoryID++
	categories[category.ID] = category
	c.JSON(http.StatusCreated, category)
}
