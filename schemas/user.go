package schemas

import (
	"bytes"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JwtKey []byte

// User represents a user in the system with fields for username, user ID, password, email, and admin status.
type User struct {
	gorm.Model

	Username string
	UserID   string `gorm:"uniqueIndex"`
	Password string
	Email    string
	Role     string
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// Check if the password is empty, if so, skip hashing
	if u.Password == "" {
		return nil
	}

	// Check if the password is already hashed
	passwordBytes := []byte(u.Password)
	if bytes.HasPrefix(passwordBytes, []byte("$2a$")) ||
		bytes.HasPrefix(passwordBytes, []byte("$2b$")) ||
		bytes.HasPrefix(passwordBytes, []byte("$2y$")) {
		return nil
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password back to the User struct
	u.Password = string(hashedPassword)
	return nil
}

// Claims represents the JWT claims for a user, including username, user ID, admin status, and standard registered claims.
type UserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

// RegisterRequest represents the request payload for user registration, including username, email, and password.
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest represents the request payload for user login, including username and password.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Claims represents the JWT claims for a user, including username, user ID, admin status, and standard registered claims.
type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
