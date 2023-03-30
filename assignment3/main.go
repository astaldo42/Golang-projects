package main

import (
	"github.com/Jateq/ass3/client"
	"github.com/Jateq/ass3/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.Init()
	router := gin.Default()
	router.GET("/", client.GreetMe)
	router.GET("/all", client.AllBooks)
	router.GET("/all/sorted", client.Sorter)
	router.POST("book", client.AddBook)
	router.PUT("book/:id", client.UpdateBook)
	router.DELETE("book/:id", client.DeleteBook)
	router.GET("/book/:id", client.Search)
	router.Run("localhost:9090")
}
