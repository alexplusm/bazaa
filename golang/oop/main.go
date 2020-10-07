package main

import "fmt"

// MyInt kek
type MyInt int

func (i MyInt) printKek() {
	fmt.Println(i)
}

func (i *MyInt) add(n int) {
	*i = *i + MyInt(n)
}

//MyStruct lol
type MyStruct struct {
	num  int
	name string
}

type mySlice []MyStruct

func part1() {
	var myInt MyInt

	myInt.printKek()
	myInt.add(5)
	myInt.add(5)
	myInt.add(5)
	myInt.printKek()
}

func part2() {
	mySlice1 := mySlice{{123, "kek"}, {456, "aza"}}
	fmt.Println(mySlice1, len(mySlice1))
}

type bird struct {
	fly bool
}

func (b *bird) toggleFly() {
	(*b).fly = !(*b).fly
}

type animal struct {
	bird
	fly int
}

func (a *animal) toggleFly() {
	(*a).fly = 1000
}

func part3() {
	b := bird{true}
	a := animal{b, 3}

	fmt.Println(a, a.bird.fly)

	a.toggleFly()
	fmt.Println("after toggle", a, a.bird.fly)
}

func main() {
	// part1()
	// part2()
	part3()
}
