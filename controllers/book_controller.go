package controllers

import (
	"strconv"
	"time"

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

func FindAllBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books:" + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind json:" + err.Error(),
		})
	}

	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	var bookDatabase models.Book

	err = db.First(&bookDatabase, intId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})
		return
	}

	book.ID = uint(intId)
	book.CreatedAt = bookDatabase.CreatedAt
	book.UpdatedAt = time.Now()

	err = db.Model(&bookDatabase).Where("ID = ?", c.Param("id")).Updates(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, intId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)

}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json:" + err.Error(),
		})
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book:" + err.Error(),
		})
		return
	}

	c.JSON(201, book)

}
