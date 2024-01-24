package controller

import (
	"API/initializer"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowToko(c *gin.Context) {

	var toko []models.Toko

	initializer.DB.Raw("select * from users").Scan(&toko)

	c.JSON(http.StatusOK, gin.H{
		"toko": toko,
	})
}

func CreateToko(c *gin.Context) {

	var toko models.Toko

	if err := c.ShouldBindJSON(&toko); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error bang",
			"error":   err.Error(),
		})
		return
	}
	initializer.DB.Create(&toko)
	c.JSON(http.StatusOK, gin.H{"toko": toko})

}
