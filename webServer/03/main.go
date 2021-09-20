package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hiHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello EveryOne!")

}
func byHandler (w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Bye EveryOne!")
}

func main (){

	r := mux.NewRouter()
	r.HandleFunc("/hi",hiHandler)
	r.HandleFunc("/bye",byHandler)
	http.ListenAndServe(":8095",r)

}
