package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("Start running authentication middleware")

		stringToken := c.Request.Header.Get("Authorization")
		if stringToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is missing",
			})
			c.Abort()
			return
		}

		userInfo, ok := ParseToken(stringToken)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Set("StudentId", userInfo.StudentId)

		c.Next()
		t2 := time.Since(t)
		fmt.Println("Execute in:", t2)
	}
}
