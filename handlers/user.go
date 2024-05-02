// handlers/user.go
package handlers

import (
	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
)

// Handler to get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

// Handler to get a single user by ID
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	c.JSON(200, user)
}

// Handler to create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(201, user)
}

// Handler to update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, user)
}

// Handler to delete a user
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
