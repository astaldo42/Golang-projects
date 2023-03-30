package main

import (
	"github.com/Jateq/oop2/base"
	"github.com/Jateq/oop2/client"
	"github.com/Jateq/oop2/mydatabase"
	"github.com/gin-gonic/gin"
)

func main() {
	mydatabase.LoadDataFromJSON("users.json", &base.Users)
	mydatabase.LoadDataFromJSON("items.json", &base.Items)

	router := gin.Default()
	//public := router.Group("/")
	router.GET("/", client.GreetMe)
	router.POST("/register", client.Register)
	router.Use(client.Authorize(base.Users))
	router.GET("/list", client.Display)
	router.GET("/fprice", client.SortByPrice)
	router.GET("/frating", client.SortByRating)
	router.GET("/search/:name", client.Search)
	router.PATCH("/change/:name/:rating", client.ChangeRating)
	router.Run("localhost:9090")
	mydatabase.SaveDataToJSON("users.json", &base.Users)
	mydatabase.SaveDataToJSON("items.json", &base.Items)
	//filters.FilterItemsByRating()
}
