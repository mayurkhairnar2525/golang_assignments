package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,EveryOne!"))
	vars := mux.Vars(r)
	title := vars["Title"]
	page := vars["Page"]
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/title/{title}/page/{page}", handler)
	log.Fatal(http.ListenAndServe(":8090", r))
}
