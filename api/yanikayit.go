package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahmali/api_denemeleri/models"
	"github.com/mahmali/api_denemeleri/veritabani"
)

func YeniKayit(c *gin.Context) {
	var model models.Kullanici
	c.ShouldBindJSON(&model) //gelen verileri birleştirir.

	eklenecek := models.Kullanici{
		Isim:  model.Isim,
		Mail:  model.Mail,
		Sifre: model.Sifre,
	}
	fmt.Println(eklenecek)
	err := veritabani.DB().Create(&eklenecek).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, "yeni kayit eklendi")
}
