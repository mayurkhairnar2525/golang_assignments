package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	Db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:12345678@tcp(0.0.0.0:9090)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("err",err)
	}
	Db = d
}

func GetDB() *gorm.DB {
	return Db
}
