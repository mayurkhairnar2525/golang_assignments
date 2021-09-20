package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	   fmt.Println("Endpoint Hit: homePage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main (){
	 r := mux.NewRouter()
	r.HandleFunc("/",homePage).Methods("GET")

	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	r.HandleFunc("/article",returnAllArticles).Methods("POST")
	 log.Fatal(http.ListenAndServe(":8095",r))
}
