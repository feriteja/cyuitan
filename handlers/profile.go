package handlers

import (
	"net/http"

	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
)

type ProfileEditRequest struct {
	Name           string `json:"email"`
	ProfilePicture string `json:"password"`
}

func EditProfile(c *gin.Context) {
	var req ProfileEditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface := c.MustGet("userID").(float64)
	userID := uint(userIDInterface)

	var profile models.Profile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name != "" {
		profile.Name = req.Name
	}
	if req.ProfilePicture != "" {
		profile.ProfilePicture = req.ProfilePicture
	}

	// Save the updated profile to the database
	if err := database.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile Updated"})

}
