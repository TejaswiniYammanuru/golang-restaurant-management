package controllers

import (
	"net/http"
	"time"
	"regexp"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"` 
	jwt.StandardClaims
}


var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	
	if err := models.CreateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// This function handles the login process for users.
// It expects a JSON payload containing the user's email and password.
// If the credentials are valid, it generates a JWT token and returns it along with the user's information.
func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve the user from the database using their email.
    storedUser, err := models.GetUserByEmail(database.DB, user.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Verify the provided password with the hashed password stored in the database.
    if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Set the expiration time for the JWT token.
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Email: user.Email, 
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    
    // Generate a new JWT token with the user's email and expiration time.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    
    // If there's an error during the signing process, return an HTTP 500 Internal Server Error response with an error message.
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }


    // Return the JWT token and the user's information in the response.
    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
        "user": gin.H{
            "id":         storedUser.ID,
            "first_name": storedUser.FirstName,
            "last_name":  storedUser.LastName,
            "email":      storedUser.Email, 
            "avatar":     storedUser.Avatar,
            "phone":      storedUser.Phone,
        },
		"claims":claims,
    })
}
