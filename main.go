package main

import (
	"gin_demo/common"
	"gin_demo/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func SetupRouter(router *gin.Engine) {
	//gin.DisableConsoleColor()

	router.GET("/ping", controllers.Home)
	router.GET("/user/:name", controllers.UserValue)
	// Student related
	router.GET("/student", controllers.FindStudents)
	router.GET("/student/:id", controllers.FindStudent)
	router.POST("/student", controllers.CreateStudent)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("/student/:id", controllers.DeleteStudent)
}

func main() {

	// Prepare DB connection
	//db := common.ConnectMySQL()
	//common.Migrate(db)
	//common.ConnectRedis()
	router := gin.Default()
	router.Use(common.LoggerMiddleware())
	// Get Server Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	// Setup router
	SetupRouter(router)

	router.Run(":" + port)
}
