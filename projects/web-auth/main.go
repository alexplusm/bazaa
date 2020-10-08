package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

func base64Test() {
	str := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	fmt.Println(str)
}

func passwordHash(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating hash from password: %w", err)
	}
	return bs, nil
}

func comparePasswords(password string, hasedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hasedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("Error password: %w", err)
	}
	return nil
}

func testPasswordHashing() {
	pass := "12345"
	hash, err := passwordHash(pass)
	if err != nil {
		panic(err)
	}
	err = comparePasswords("1234", hash)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Login in!")
}

func runServer() {
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

func keyGen() []byte {
	var key []byte = []byte{}
	var i byte

	for i = 0; i < 64; i++ {
		key = append(key, i)
	}

	return key
}

func main() {
	// base64Test()
	// runServer()
	// testPasswordHashing()
}
