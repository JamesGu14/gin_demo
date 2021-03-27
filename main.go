package main

import (
	"gin_demo/controllers"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func setupRouter() *gin.Engine {
	//gin.DisableConsoleColor()
	router := gin.Default()

	router.GET("/ping", controllers.Home)
	router.GET("/user/:name", controllers.UserValue)
	// Student related
	router.GET("/student", controllers.FindStudents)
	router.GET("/student/:id", controllers.FindStudent)
	router.POST("/student", controllers.CreateStudent)
	router.PUT("/student/:id", controllers.UpdateStudent)

	return router
}

func main() {
	// Get Server Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Prepare DB connection
	db := util.InitDB()
	util.Migrate(db)

	// Setup router
	router := setupRouter()
	router.Run(":" + port)
}
