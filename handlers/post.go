package handlers

import (
	"net/http"

	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
type PostEditRequest struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func SendPost(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Accessing fields of req

	userIDInterface := c.MustGet("userID").(float64)

	userID := uint(userIDInterface)

	// Create a new auth record
	post := models.Post{
		Title:  req.Title,
		Detail: req.Detail,
		UserID: userID,
	}

	// Save the auth record to the database
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create the Post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post Created"})

}

func EditPost(c *gin.Context) {
	var req PostEditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface := c.MustGet("userID").(float64)
	userID := uint(userIDInterface)

	var post models.Post
	if err := database.DB.Where("id = ?", req.ID).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not qualified to Edit this post"})

	}

	post.Detail = req.Detail
	post.Title = req.Title

	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post Created"})

}
