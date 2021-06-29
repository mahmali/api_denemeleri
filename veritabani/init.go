package veritabani

import (
	"fmt"
	"github.com/mahmali/api_denemeleri/config"
	"github.com/mahmali/api_denemeleri/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm/schema"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func Baglan() {
	var err error
	constr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_NAME,
		config.DB_PASS)
	db, err = gorm.Open(postgres.Open(constr)) //&gormConfig
	if err != nil {
		fmt.Println(err, constr)
		panic("Veritabanına bağlanılamıyor.")
	}
}
func AutoMigrate() {
	db.AutoMigrate(&models.Kullanici{})
	//db.AutoMigrate(&models.Esya{})
	//db.AutoMigrate(&models.Depo{})
	//db.AutoMigrate(&models.Kullanici{})
}
func KullaniciOlustur() {
	kullaniciSayisi := int64(0)
	DB().Model(&models.Kullanici{}).Count(&kullaniciSayisi)
	if kullaniciSayisi == 0 {
		eklenecek := models.Kullanici{
			Isim:  "muhammed ali",
			Mail:  "muh@gmail.com",
			Sifre: "123",
		}
		err := db.Create(&eklenecek).Error
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
