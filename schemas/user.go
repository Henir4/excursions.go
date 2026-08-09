package schemas

import (
	"bytes"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JwtKey = []byte(os.Getenv("JWT"))

type User struct {
	gorm.Model

	Username string
	UserID   string
	Password string
	Email    string
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

type UserResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	jwt.RegisteredClaims
}
