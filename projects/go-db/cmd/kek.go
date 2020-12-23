package main

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

func main() {
	serv := services.ValidateFacesService{}
	img := "0-DVN_b_SVAO_541_1-04_08_2020_13_00_30.jpg"
	a, b := serv.Validate("media_root/" + img)
	fmt.Println(a, b)
}
