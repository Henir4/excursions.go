package middleware

import (
	"time"

	"excursion.com/schemas"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *schemas.User) (string, error) {
	// Set the expiration time for the token (e.g., 24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the user information and expiration time
	claims := &schemas.Claims{
		Username: user.Username,
		UserID:   user.UserID,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   user.UserID,
			Issuer:    "RotaUnicaViagens",
		},
	}

	// Create the token using the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(schemas.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}