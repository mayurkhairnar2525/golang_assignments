package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/config"
)

var db *gorm.DB

type BookManagement struct {
	gorm.Model
	  Name         string `json:"name"`
	  Author       string `json:"author"`
	  Prices       string `json:"prices"`
	  Available    string `json:"available"`
	  PageQuality  string `json:"pagequality"`
	  LaunchedYear string `json:"launchedyear"`
	  Isbn         string `json:"isbn"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BookManagement{})
}

func (b *BookManagement) CreateBook() *BookManagement {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllBooks() []BookManagement {
	var Books []BookManagement
	db.Find(&Books)
	return Books
}

func GetBookByAuthor(Author string) (*BookManagement , *gorm.DB){
	var getBook BookManagement
	db:=db.Where("author=?", Author).Find(&getBook)
	return &getBook, db

}
