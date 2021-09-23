package routes

import (
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookStore-REST/pkd/controller"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{name}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books/{isbn}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{Id}", controller.DeleteBook).Methods("DELETE")
	http.ListenAndServe(":8090", router)
}

