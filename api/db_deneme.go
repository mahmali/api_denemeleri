package api

import (
	"fmt"
	"github.com/mahmali/api_denemeleri/models"
	"github.com/mahmali/api_denemeleri/veritabani"
)

func DBdeneme() {
	//var TümKayitModel []models.Kullanici
	//if err:= veritabani.DB().Find(&TümKayitModel).Error; err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("DBdeneme",TümKayitModel)

	//var user models.Kullanici
	//id:=5
	//veritabani.DB().Where("id = ?", id).Delete(&user)

	//eklenecek:=models.Kullanici{
	//	Isim:"db debene",
	//	Mail: "db deneme",
	//	Sifre: "db deneme",
	//}
	//fmt.Println(eklenecek)
	//err:=veritabani.DB().Create(&eklenecek).Error
	//if err!= nil{
	//	fmt.Println(err.Error())
	//}

	var user models.Kullanici
	err := veritabani.DB().Model(&user).Where("id = ?", "56").Save(models.Kullanici{
		Isim:  "deneme2",
		Mail:  "deeneme2",
		Sifre: "deneme2",
		ID:    56,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
