package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahmali/api_denemeleri/models"
	"github.com/mahmali/api_denemeleri/veritabani"
)

func Guncelle(c *gin.Context) {
	var model models.Kullanici
	c.ShouldBindJSON(&model)

	fmt.Println(model)
	var user models.Kullanici
	err := veritabani.DB().Model(&user).Where("id = ?", model.ID).Save(
		models.Kullanici{
			Isim:  model.Isim,
			Mail:  model.Mail,
			Sifre: model.Sifre,
			ID:    model.ID,
		}).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, "işlem başarılı")
}
