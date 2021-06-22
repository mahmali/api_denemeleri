package main

import (
	"fmt"
	"github.com/mahmali/api_denemeleri/api"
	"github.com/mahmali/api_denemeleri/veritabani"
)

func main() {
	fmt.Println("bismillah")
	veritabani.Baglan()
	veritabani.AutoMigrate()
	veritabani.KullaniciOlustur()
	api.Login()

}
