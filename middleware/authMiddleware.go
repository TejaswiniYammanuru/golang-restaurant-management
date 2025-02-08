// Package middleware contains custom middleware functions for the Gin framework.
package middleware

import (
    "net/http"
    "strings"

    // Import the jwt-go package for JWT handling.
    "github.com/dgrijalva/jwt-go"
    // Import the Gin framework for creating HTTP web servers and middleware.
    "github.com/gin-gonic/gin"
)

// jwtKey is a byte slice used as the secret key for signing and verifying JWT tokens.
var jwtKey = []byte("my_secret_key")

// Claims represents the custom claims structure for JWT tokens.
type Claims struct {
    Email string `json:"email"` // Email field to store the user's email.
    jwt.StandardClaims           // Embed the jwt.StandardClaims struct to include standard JWT claims.
}

// Authentication is a Gin middleware function that authenticates incoming requests using JWT tokens.
func Authentication() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the "Authorization" header from the request.
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            // If the "Authorization" header is missing, return an unauthorized response.
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // Trim the "Bearer " prefix from the token string.
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims := &Claims{}

        // Parse the JWT token with the custom claims structure.
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            // If the token is invalid or parsing fails, return an unauthorized response.
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set the "email" value in the Gin context for further use in the request handler.
        c.Set("email", claims.Email) 
        // Move on to the next middleware or request handler.
        c.Next()
    }
}
