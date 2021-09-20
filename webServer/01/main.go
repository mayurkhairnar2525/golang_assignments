package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
// Request Handler
func handler (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello Everyone!"))
}

func main (){
	// New router
	// Instant of router
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/",handler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080",r))
}

