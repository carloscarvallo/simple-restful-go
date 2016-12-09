package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPersonEndpoint(w http.ResponseWritter, req *http.Request) {

}

func GetPeopleEndpoint(w http.ResponseWritter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWritter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWritter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":12345", router))
}
