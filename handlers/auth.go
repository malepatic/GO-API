package handlers

import (
	"net/http"

	"go-api/database"
	"go-api/models"
	"go-api/utils"

	"github.com/gin-gonic/gin"
)

// Register handler
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the username is already taken
	for _, u := range database.Users {
		if u.Username == user.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Save the user
	user.ID = len(database.Users) + 1
	user.Password = hashedPassword
	database.Users = append(database.Users, user)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handler
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find user by username
	var foundUser models.User
	for _, u := range database.Users {
		if u.Username == input.Username {
			foundUser = u
			break
		}
	}

	if foundUser.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Comparing passwords
	if !utils.ComparePasswords(foundUser.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
