package controller

import (
	"API/initializer"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	id := uuid.New()
	role := "User"

	var body struct {
		Nama_user string
		Email     string
		Password  string
	}
	// GET THE VALUE FROM BODY
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to read body",
		})
		return
	}
	// hash disini
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to hash pass",
		})
		return
	}

	// INSERT TO DATABASE
	user := models.User{ID_pembeli: id.String(), Email: body.Email, Nama_users: body.Nama_user, Password: string(hashedPass), Role: role}
	initializer.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func Login(c *gin.Context) {

	var user models.User
	var body struct {
		Email    string
		Password string
	}
	// GET THE VALUE FROM BODY
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to read body",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error bang",
			"error":   err.Error(),
		})
		return
	}
	initializer.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

	// SEARCH FOR AN USER
	initializer.DB.First(&user, "email = ?", body.Email)

	if user.ID_pembeli == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

}
