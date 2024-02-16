package controller

import (
	"API/initializer"
	"API/models"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var cibai string = "nfjankdnfqn31oneiowe"

// var cibai string = "aaaaa"

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
	var res models.User
	if err := initializer.DB.Where("email = ?", body.Email).First(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// hash disini
			hashedPass := generatePass(body.Password, body.Email)

			// INSERT TO DATABASE
			user := models.User{ID_pembeli: id.String(), Email: body.Email, Nama_users: body.Nama_user, Password: hashedPass, Role: role}

			initializer.DB.Create(&user)
			c.JSON(http.StatusOK, gin.H{"user": user})
		} else {
			log.Fatal("error while querying:", err)
		}
	} else {
		// Data already exists in the database, handle accordingly
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email Sudah Terdaftar"})
	}

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
	err := CompareHashAndPassword(user.Password, generatePass(body.Password, body.Email))

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

func CompareHashAndPassword(hashedPass string, pass string) error {
	if hashedPass != pass {
		return errors.New("invalid Password")
	}
	return nil
}

func generatePass(pass string, aa string) string {
	var newPass string
	passIndex, aaIndex, bbIndex := 0, 0, 0
	println(len(cibai))
	// salting
	for passIndex < len(pass) || aaIndex < len(aa) || bbIndex < len(cibai) {

		if passIndex < len(pass) {
			newPass += toHex(pass[passIndex])
			passIndex++
		}
		if aaIndex < len(aa) {
			newPass += toHex(aa[aaIndex])
			aaIndex++
		}
		if bbIndex < len(cibai) {
			newPass += toHex(cibai[bbIndex])
			bbIndex++
		}
	}

	mask := "$a2$" + newPass[:38]
	hashed := sha256.Sum256([]byte(mask))
	hashedPass := fmt.Sprintf("%x", hashed)
	return hashedPass
}

func toHex(char byte) string {
	if char >= '0' && char <= '9' || char >= 'a' && char <= 'f' || char >= 'A' && char <= 'F' {
		return string(char)
	}
	return fmt.Sprintf("%x", char)
}
