package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

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
	people = append(people, Person{ID: "1", FirstName: "Carlos", Lastname: "Carv", Address: &Address{City: "Asuncion", State: "Central"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Method("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Method("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Method("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Method("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
