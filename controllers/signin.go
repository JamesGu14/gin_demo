package controllers

import (
	"gin_demo/dao"
	services "gin_demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	var input dao.SignInInput
	// Validate input username and password
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	stringToken, ok := services.StudentSignIn(input.UserName, input.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": stringToken})
}
