package routes

import (
	"full-crud/models"
	"full-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Bind the incoming JSON to the user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	// Save the user to the database
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authanticate user.", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authanticate user.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
