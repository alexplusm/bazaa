package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", decodeListPerson)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Panic("Server doesn't start", err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Alex",
	}
	p2 := person{
		First: "Lee",
	}

	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("Encoded bad data!", err)
	}
}

func decodeListPerson(w http.ResponseWriter, r *http.Request) {
	var people []person
	err := json.NewDecoder(r.Body).Decode(&people)

	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Request body", people)
}

func decodeOnePerson(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)

	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Request body", p1)
}
