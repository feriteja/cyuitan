// models/user.go
package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	AuthID    uint      `json:"auth_id"`
	ProfileID *uint     `json:"profile_id"`
	Auth      Auth      `json:"auth" gorm:"foreignkey:AuthID"`
	Profile   Profile   `json:"profile" gorm:"foreignkey:ProfileID"`
	Posts     []Post    `json:"posts"`
	Comments  []Comment `json:"comments"`
}

// Auth model
type Auth struct {
	gorm.Model
	Email           string `json:"email"`
	Password        string `json:"password"`
	ActivationToken string `json:"activation_token"`
	Status          int    `json:"status"`
	UserID          uint   `json:"user_id"`
}

// Post model
type Post struct {
	gorm.Model
	Title        string    `json:"title"`
	Detail       string    `json:"detail"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	UserID       uint      `json:"user_id"`
	Comments     []Comment `json:"comments"`
}

// Comment model
type Comment struct {
	gorm.Model
	Detail    string    `json:"detail"`
	PostID    uint      `json:"post_id"`
	CommentID uint      `json:"comment_id"` // Assuming this is the parent comment ID, change to uint if it's the comment ID itself
	UserID    uint      `json:"user_id"`
	Comments  []Comment `json:"comments"`
}

// Profile model
type Profile struct {
	gorm.Model

	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         uint   `json:"user_id"`
}
