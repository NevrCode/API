package main

import (
	"API/controller"
	"API/initializer"
	"API/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()

}

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.SessionMiddleware())
	r.GET("/showUser", controller.ShowUser)
	r.GET("/showToko", controller.ShowToko)
	r.POST("/showAduh", controller.ShowAduh)
	r.POST("/createUser", controller.SignUp)
	r.POST("/login", controller.Login)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// r.Run("0.0.0.0:" + port)
	r.Run()

}
