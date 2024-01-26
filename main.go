package main

import (
	"API/controller"
	"API/initializer"
	"net/http"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/showUser", controller.ShowUser)
	r.GET("/showToko", controller.ShowToko)
	r.POST("/showAduh", controller.ShowAduh)
	r.POST("/createUser", controller.SignUp)
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// r.Run("0.0.0.0:" + port)
	r.Run()
}

// func main() {
// http.HandleFunc("/", controller.Index)
// fmt.Print("jalam mas")
// http.ListenAndServe(":8080", nil)
// }
