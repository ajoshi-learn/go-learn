package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	people := GetPeopleData()
	json.NewEncoder(writer).Encode(people)
}
