package handlers

import (
	"net/http"
	"strconv"

	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
type PostEditRequest struct {
	ID     int    `json:"id"`
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

	if req.Detail != "" {

		post.Detail = req.Detail
	}
	if req.Title != "" {

		post.Title = req.Title
	}

	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post Edited"})

}

func GetPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	var posts []models.Post
	offset := (page - 1) * 15 // Assuming 15 posts per page
	if err := database.DB.Offset(offset).Limit(15).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
