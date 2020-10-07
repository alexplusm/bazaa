package main

import (
	"fmt"
	"github.com/Alexplusm/bazaa/golang/mypackage"
)

func slicesPart1() {
	array := make([]int, 0, 15)

	fmt.Println(array, len(array), cap(array))

	/* FILL SLICE */
	// for i := 0; i < 15; i++ {
	// 	array = append(array, i)
	// }

	array = append(array, []int{1, 2, 3}...)

	fmt.Println("filled", array, len(array), cap(array))

	fmt.Println("----------------------------")

	slice1 := []int{1, 3, 4}

	fmt.Println("slice1_0", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 42)

	fmt.Println("slice1_1", slice1, len(slice1), cap(slice1))

	slice2 := slice1
	slice1[2] = 666
	fmt.Println("slice1_2", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2_0", slice2, len(slice2), cap(slice2))

	slice3 := slice1
	slice1 = append(slice1, 47)

	slice1[2] = 777

	fmt.Println("slice1_3", slice1, len(slice1), cap(slice1))
	fmt.Println("slice3_0", slice3, len(slice3), cap(slice3))

	slice4 := slice1
	slice1 = append(slice1, []int{0}...)

	slice1[2] = 999

	fmt.Println("slice1_4", slice1, len(slice1), cap(slice1))
	fmt.Println("slice4_0", slice4, len(slice4), cap(slice4))

	slice1 = append(slice1, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}...)
	slice1[2] = 123

	fmt.Println("slice1_5", slice1, len(slice1), cap(slice1))
	fmt.Println("slice4_1", slice4, len(slice4), cap(slice4))

	// incorrenct slice copy
	var slice5 []int
	copy(slice5, slice1)

	fmt.Println("slice1_6", slice1, len(slice1), cap(slice1))
	fmt.Println("slice5_0", slice5, len(slice5), cap(slice5))

	// correct slice cope

	slice6 := make([]int, len(slice1), len(slice1))
	slice6 = slice1
	fmt.Println("slice1_7", slice1, len(slice1), cap(slice1))
	fmt.Println("slice6_0", slice6, len(slice6), cap(slice6))
}

func slicesPart2() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("slice1_0", slice1, len(slice1), cap(slice1))
	fmt.Println("part slice1[:4]", slice1[:4], len(slice1[:4]), cap(slice1[:4]))
	fmt.Println("part slice1[2:]", slice1[2:], len(slice1[2:]), cap(slice1[2:]))

	slice2 := append(slice1[:2], slice1[4:]...)
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	arr := [...]int{1, 2, 3}

	slice3 := arr[:]
	fmt.Println("arr", arr, len(arr), cap(arr))
	fmt.Println("slice3", slice3, len(slice3), cap(slice3))

	slice3[1] = 666
	fmt.Println("update slice3[1]")
	fmt.Println("arr", arr, len(arr), cap(arr))
	fmt.Println("slice3", slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, []int{1, 2, 3}...)
	slice3[1] = 123

	fmt.Println("update slice3[1] after append")
	fmt.Println("arr", arr, len(arr), cap(arr))
	fmt.Println("slice3", slice3, len(slice3), cap(slice3))
}

func maps() {
	var m1 map[string]string

	m1 = map[string]string{"kek": "1"}

	fmt.Println("m1", m1)
	val, exist := m1["kek"]

	fmt.Println(val, exist)

	m2 := m1

	key := "kek"

	fmt.Println("m2", m2, len(m2))
	delete(m1, key)
	fmt.Println("m2", m2, len(m2))
}

//  TODO: func which copy map

func iterations() {

	slice := [...]int{1, 2, 3, 4}

	for index := range slice {
		fmt.Println(index, slice[index])
	}

	for index, value := range slice {
		fmt.Println("index with value", index, value)
	}

	for _, value := range slice {
		fmt.Println("only value", value)
	}

	mm := map[string]string{"first": "1", "second": "2"}

	for key := range mm {
		fmt.Println("map key:", key, mm[key])
	}
	for key, value := range mm {
		fmt.Println("map key and val", key, value)
	}
}

func sswitch() {
	mm := map[string]string{}

	// mm["1"] = "first"
	mm["1"] = "one"
	mm["2"] = "second"
	// mm["flag"] = "kek"
	mm["flag"] = "lol"

	switch mm["1"] {
	case "first", "2":
		fmt.Println("digits", mm["first"])
	case "one":
		if mm["flag"] == "lol" {
			fmt.Println("tikaem", mm["flag"])
			break
		}
		fmt.Println("ne tikaem", mm["flag"])
		fallthrough // to next case

	default:
		fmt.Println("default")
	}
}

func runes() {

	// str1 := "Привет"
	str1 := "你好，世界"

	fmt.Println("one byte of rune", str1[1])
	// iteration on symbols
	for index, runeValue := range str1 {
		fmt.Println(index, runeValue)
	}

	bytes := []byte(str1)

	for index, value := range bytes {
		fmt.Println("byte value", index, value)
	}
}

func packageTest() {
	fmt.Println(mypackage.Kek)
}

func main() {
	// slicesPart1()
	// slicesPart2()
	// maps()
	// iterations()
	// sswitch()
	// runes()

	packageTest()
}
