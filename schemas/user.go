package schemas

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
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

type UserResponse struct {
  ID       uint   `json:"id"`
  UserID   string `json:"user_id"`
  Username string `json:"username"`
  Email    string `json:"email"`
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
