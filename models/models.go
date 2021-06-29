package models

type Kullanici struct {
	ID    int64  `gorm:"primarykey" json:"id,omitempty"`
	Isim  string `json:"isim" gorm:"column:isim"`
	Mail  string `json:"mail" gorm:"column:mail"`
	Sifre string `json:"sifre" gorm:"column:sifre"`
}
