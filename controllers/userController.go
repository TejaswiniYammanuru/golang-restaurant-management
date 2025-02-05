package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc { // Defines a function named GetUsers that returns a gin.HandlerFunc, which is a type for handling HTTP requests in the Gin framework.
	return func(c *gin.Context) { // Returns an anonymous function that takes a gin.Context as a parameter, used to handle the HTTP request context.

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {

}
