package main

import (
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/routes"
	"log"
	"net/http"
)

func main (){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8090", r))
}
