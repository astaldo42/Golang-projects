package client

import (
	"github.com/Jateq/ass3/database"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
)

func GreetMe(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to LitRes")
}
func AllBooks(context *gin.Context) {
	db := database.GetDB()
	var books []database.Books
	db.Find(&books)
	context.JSON(http.StatusOK, gin.H{"books": books})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter

	// Retrieve the book from the database
	db := database.GetDB()
	var book database.Books
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Bind the updated fields from JSON input to the book struct
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Return the updated book as JSON response
	c.JSON(http.StatusOK, book)
}

func Search(context *gin.Context) {
	id := context.Param("id")
	db := database.GetDB()
	var book database.Books
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, book)
}

func DeleteBook(context *gin.Context) {
	db := database.GetDB()
	id := context.Param("id")
	var book database.Books
	db.Where("id = ?", id).Find(&book)
	if book.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&book)
	context.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func Sorter(context *gin.Context) {
	db := database.GetDB()
	order := context.Query("order")

	var books []database.Books
	if order == "asc" {
		db.Order("cost asc").Find(&books)
	} else if order == "desc" {
		db.Order("cost desc").Find(&books)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort order"})
		return
	}

	context.JSON(http.StatusOK, books)
}

func AddBook(context *gin.Context) {
	db := database.GetDB()

	var book database.Books
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book record
	result := db.Create(&book)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Book added successfully", "data": book})
}
