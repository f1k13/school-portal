package userHandlers

import (
	"net/http"

	"github.com/f1k13/school-portal/internal/handlers/dto"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/services"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user dto.UserDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Error("Error binding JSON", err)
		return
	}
	u, err := services.SignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Error("Error creating user", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": u, "message": "User created successfully"})
}

func SignIn(c *gin.Context) {}
