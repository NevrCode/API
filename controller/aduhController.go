package controller

import (
	"API/initializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Aduh struct {
	A string
	N string
}

func ShowAduh(c *gin.Context) {
	var aduh []Aduh

	initializer.DB.Raw("select * from aduh").Scan(&aduh)

	c.JSON(http.StatusOK, gin.H{
		"aduh": &aduh,
	})
}
