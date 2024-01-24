package main

import (
	"API/controller"
	"API/initializer"

	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/showUser", controller.ShowUser)
	r.GET("/showToko", controller.ShowToko)
	r.Run()
}

// func main() {
// http.HandleFunc("/", controller.Index)
// fmt.Print("jalam mas")
// http.ListenAndServe(":8080", nil)
// }
