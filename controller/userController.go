package controller

import (
	"API/initializer"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
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
	// SEARCH FOR AN USER
	initializer.DB.First(&user, "email = ?", body.Email)
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email",
		})
		return
	}
	// hash to pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Logged in",
	// })

	session := c.MustGet("session").(*sessions.Session)
	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)
	session.Values["loggedIn"] = true
	session.Values["user"] = user
	c.JSON(http.StatusOK, gin.H{
		"user": session.Values["user"],
	})

}

// func userHandler(c *gin.Context) {
// 	session := c.MustGet("session").(*sessions.Session)
// 	authenticated := session.Values["authenticated"]
// 	if authenticated != nil && authenticated.(bool) {
// 		c.String(http.StatusOK, "User is authenticated")
// 	} else {
// 		c.String(http.StatusUnauthorized, "User is not authenticated")
// 	}
// }
