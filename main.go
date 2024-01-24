package main

import (
	"API/controller"
	"API/initializer"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Ganti dengan origin aplikasi klien Anda
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/showUser", controller.ShowUser)
	r.GET("/showToko", controller.ShowToko)
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Run("0.0.0.0:" + port)
}

// func main() {
// http.HandleFunc("/", controller.Index)
// fmt.Print("jalam mas")
// http.ListenAndServe(":8080", nil)
// }
