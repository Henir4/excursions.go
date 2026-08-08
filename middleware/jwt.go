package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a middleware function that checks for the presence of a valid JWT token in the Authorization header of incoming requests. If the token is valid, it allows the request to proceed; otherwise, it aborts the request and returns an unauthorized error.
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")

		// Check if the token is present in the header
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// Split the token string to extract the actual token
		// The expected format is "Bearer <token>"
		// This is a common convention for passing JWT tokens in HTTP headers
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		// Extract the token string from the header
		tokenString = parts[1]

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &schemas.Claims{}, func(token *jwt.Token) (interface{}, error) {
			
			// Ensure the signing method is HMAC
			// This is a security measure to prevent certain types of attacks
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key used for signing the token
			return schemas.JwtKey, nil

			},_)

			// Check if there was an error parsing the token or if the token is invalid
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// If the token is valid, extract the claims and set them in the context for use in subsequent handlers
		if claims, ok := token.Claims.(*schemas.Claims); ok && token.Valid {

			user := schemas.User{
				UserID: claims.UserID,
				Username: claims.Username,
			}

			// Set the user information in the context for use in subsequent handlers
			c.Set("user", user)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// If everything is valid, proceed to the next handler
		c.Next()

		
	}
}

func GenerateToken(user *schemas.User) (string, error) {
	// Set the expiration time for the token (e.g., 24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the user information and expiration time
	claims := &schemas.Claims{
		Username: user.Username,
		UserID:   user.UserID,
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