package controller


import (
	"net/http"
	"os"

	"github.com/nethsaraPrabash/potral-backend/src/service"
	"github.com/nethsaraPrabash/potral-backend/src/model"
	"github.com/nethsaraPrabash/potral-backend/src/helper"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := helper.GenerateJWT(user.ID, user.Role, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token})

}

func LoginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := helper.GenerateJWT(user.ID, user.Role, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}