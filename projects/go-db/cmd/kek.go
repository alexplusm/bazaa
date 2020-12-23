package main

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

func main() {
	serv := services.ValidateFacesService{}
	img := "0-DVN_SAO_1094_0-07_08_2020_13_10_30.jpg"
	a, b := serv.Validate(img)
	fmt.Println(a, b)
}
