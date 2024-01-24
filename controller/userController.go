package controller

import (
	"API/initializer"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {

	var user []models.User

	initializer.DB.Raw("select * from users").Scan(&user)

	// initializer.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}

// var user models.User

//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//			"message": "Error bang",
//			"error":   err.Error(),
//		})
//		return
//	}
//
// initializer.DB.Create(&user)
// c.JSON(http.StatusOK, gin.H{"user": user})
func SignUp(c *gin.Context) {

	// get nama, nama, pass
	var body struct {
		Nama_user string
		Email     string
		Password  string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to read body",
		})
		return
	}

	// create user
	user := models.User{Email: body.Email, Nama_user: body.Nama_user, Password: body.Password}
	initializer.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error bang",
			"error":   err.Error(),
		})
		return
	}
	initializer.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}
