package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahmali/api_denemeleri/models"
	"github.com/mahmali/api_denemeleri/veritabani"
)

func KayitGetir(c *gin.Context) {

	var TümKayitModel []models.Kullanici
	if err := veritabani.DB().Find(&TümKayitModel).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(200, TümKayitModel)
}
