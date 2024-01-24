package controller

import (
	"API/initializer"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowVarian(c *gin.Context) {
	var varians []models.Varian

	initializer.DB.Raw("select * from varian").Scan(&varians)
	c.JSON(http.StatusOK, gin.H{
		"varian": varians,
	})

}
