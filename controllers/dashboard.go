package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func init() {
	db["james"] = "hi James"
	db["kara"] = "hello Kara"
}

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard/dashboard.tmpl", gin.H{
		"title": "Dashboard",
	})
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func UserValue(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
