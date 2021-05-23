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
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/assets", "./assets")

	router.GET("/", controllers.Dashboard)
	router.GET("/ping", controllers.Home)
	router.GET("/user/:name", controllers.UserValue)
	// Student related
	router.POST("/signin", controllers.SignIn)
	router.GET("/student", controllers.StudentPage)
	router.GET("/student/create", controllers.StudentCreatePage)
	router.DELETE("/student/:id", util.Authentication(), controllers.DeleteStudent)
	//router.GET("/student/:id", controllers.FindStudent)
	router.POST("/student", controllers.CreateStudent)
	//router.PUT("/student/:id", controllers.UpdateStudent)

	// Teacher Related

	// Class Related

	// Course Related

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
