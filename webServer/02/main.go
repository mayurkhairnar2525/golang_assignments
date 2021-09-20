package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type incomingRequest struct {
	Request string `json:"request"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,EveryOne!"))
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}
func greetHandler(w http.ResponseWriter, r *http.Request) {

	read, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error %s", err)
		return
	}
	defer r.Body.Close()
	requestPayload := incomingRequest{}
	err = json.Unmarshal(read, &requestPayload)
	if err != nil {
		fmt.Fprintf(w, "error %s", err)
		return
	}
	fmt.Fprintf(w, "Hello %s", requestPayload.Request)

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/title/{title}/page/{page}", handler).Methods("GET")
	r.HandleFunc("/greet", greetHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", r))
}
