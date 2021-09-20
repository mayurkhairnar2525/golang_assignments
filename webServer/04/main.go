package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homeHandler ( w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello from the home site")
}

func contactHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Thanks for contacting to the home site")
}
func main (){

	r := mux.NewRouter()
	r.HandleFunc("/",homeHandler)
	r.HandleFunc("/contact",contactHandler)
	http.ListenAndServe(":9095",r)
}
