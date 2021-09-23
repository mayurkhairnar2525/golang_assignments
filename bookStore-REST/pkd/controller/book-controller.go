package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/config"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/models"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/utils"
	"log"
	"net/http"
)

var NewBook models.BookManagement

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.BookManagement{}

	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		log.Fatalln("Error", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal("error", err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookAuthor := vars["author"]
	utils.ParseBody(r, GetBook)
	bookDetails, err := models.GetBookByAuthor(bookAuthor)
	if err != nil {
		log.Fatal("err occurred", err)
	}
	res, errr := json.Marshal(bookDetails)
	if errr != nil {
		log.Fatal("err occurred", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.BookManagement
	config.Db.First(&book, params["isbn"])
	json.NewDecoder(r.Body).Decode(&book)
	config.Db.Save(&book)
	json.NewEncoder(w).Encode("Updates")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.BookManagement
	config.Db.Delete(&book, params["id"])
	json.NewEncoder(w).Encode("Deleted")
}
