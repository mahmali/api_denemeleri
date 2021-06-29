package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahmali/api_denemeleri/models"
	"github.com/mahmali/api_denemeleri/veritabani"
	"strconv"
)

func KayitSil(c *gin.Context) {
	var user models.Kullanici

	idSitring := c.Query("id")
	id, _ := strconv.Atoi(idSitring)

	veritabani.DB().Where("id = ?", id).Delete(&user)
	c.JSON(200, "silme işlemi başarılı")
}
