// handlers/auth.go
package handlers

import (
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtInfo struct {
	UserId int
	Status int
	Email  string
}

// Login handler
func Login(c *gin.Context) {
	// Parse request body
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	var auth models.Auth
	if err := database.DB.Where("email = ?", req.Email).First(&auth).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	var user models.User
	if err := database.DB.Where("auth_id = ?", auth.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtInfo := JwtInfo{
		UserId: int(user.ID),
		Status: auth.Status,
		Email:  auth.Email,
	}

	// Generate JWT token
	token, err := generateToken(jwtInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return success response with JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	// Parse request body
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email format
	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Validate password format
	if !isValidPassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password at least 8 character"})
		return
	}

	if req.ConfirmPassword != req.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password doesn't match"})
		return
	}
	// Check if the email is already registered
	if isEmailRegistered(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new auth record
	auth := models.Auth{
		Email:    req.Email,
		Password: hashedPassword,
		Status:   2,
	}

	// Save the auth record to the database
	if err := database.DB.Create(&auth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Create a new user record
	user := models.User{
		AuthID:    auth.ID,
		ProfileID: nil,
	}

	// Save the user record to the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user record"})
		return
	}

	auth.UserID = user.ID

	if err := database.DB.Save(&auth).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtInfo := JwtInfo{
		UserId: int(user.ID),
		Status: auth.Status,
		Email:  auth.Email,
	}

	token, err := generateToken(jwtInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Return success response
	c.JSON(http.StatusCreated, gin.H{"token": token, "message": "Account Created"})
}

// Function to check if the email is already registered
func isEmailRegistered(email string) bool {
	var auth models.Auth
	database.DB.Where("email = ?", email).First(&auth)
	return auth.ID != 0
}

// Function to hash the password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateToken(jwtInfo JwtInfo) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = jwtInfo.UserId
	claims["status"] = jwtInfo.Status
	claims["email"] = jwtInfo.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiry time (e.g., 24 hours)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to validate email format
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Function to validate password format
func isValidPassword(password string) bool {
	return len(password) >= 8

}
