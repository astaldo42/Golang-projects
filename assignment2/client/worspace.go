package client

import (
	"github.com/Jateq/oop2/base"
	"github.com/Jateq/oop2/filters"
	"github.com/Jateq/oop2/mydatabase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GreetMe(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to Jateq's Market")
}
func Authorize(users []base.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		for _, user := range users {
			if user.Username == username && user.Password == password {
				return
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

//func registration(username, password string, usersFile string) error {
//	user := database.User{Username: username}
//	user.SetPassword(password)
//	database.Users = append(database.Users, user)
//
//	// Write the updated user data to the JSON file
//	file, err := os.OpenFile(usersFile, os.O_RDWR|os.O_CREATE, 0644)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	encoder := json.NewEncoder(file)
//	encoder.SetIndent("", "    ")
//	if err := encoder.Encode(database.Users); err != nil {
//		return err
//	}
//
//	return nil
//}

func Register(c *gin.Context) {
	var user base.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, u := range base.Users {
		if u.Username == user.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already taken"})
			return
		}
	}

	base.Users = append(base.Users, user)

	if err := mydatabase.SaveDataToJSON("users.json", &base.Users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func ChangeRating(c *gin.Context) {

	name := c.Param("name")
	rating, err := strconv.ParseFloat(c.Param("rating"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	base.GiveRating(name, rating)

	for i := range base.Items {
		if base.Items[i].Name == name {
			base.Items[i].Rating = rating
			c.JSON(http.StatusOK, gin.H{"message": "rating updated successfully"})
			if err := mydatabase.SaveDataToJSON("items.json", &base.Items); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save data"})
				return
			}
			return
		}
	}

	// If the item was not found, return an error
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
	if err := mydatabase.SaveDataToJSON("items.json", &base.Items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save data"})
		return
	}
}

func Search(c *gin.Context) {
	var itemme []base.Item
	name := c.Param("name")
	itemme = base.SearchItemsByName(name)
	c.IndentedJSON(http.StatusOK, itemme)
}

func Display(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, base.Items)
}

func SortByPrice(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, filters.FilterItemsByPrice())
}

func SortByRating(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, filters.FilterItemsByRating())
}
