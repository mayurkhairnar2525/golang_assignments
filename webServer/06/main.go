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

var Data []string

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Data : %#v", Data)
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
	Data = append(Data, requestPayload.Request)

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get-data", handler).Methods("GET")
	r.HandleFunc("/add-data", greetHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8099", r))
}
