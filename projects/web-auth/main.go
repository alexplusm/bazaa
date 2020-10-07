package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{First: "Alex"}
	p2 := person{First: "Lee"}

	persons1 := []person{p1, p2}

	bs, err := json.Marshal(persons1)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("JSON", string(bs))

	persons2 := []person{}
	err = json.Unmarshal(bs, &persons2)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Go struct", persons2)
}
