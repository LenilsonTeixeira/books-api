package controllers

import (
	"strconv"

	"github.com/LenilsonTeixeira/books-api/database"
	"github.com/LenilsonTeixeira/books-api/database/models"
	"github.com/gin-gonic/gin"
)

func FindBookById(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, intId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}
