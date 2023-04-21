package main

import (
	"fmt"
)

func main() {
	// explicit
	var a int = 10

	// initialize default
	var b int

	/*
		0	numbers
		false	booleans
		""	string
		nil	reference types (slice,pointer,map,channel,function)
	*/

	// short (recommended)
	c := 12

	// := declare and initialize
	// = is an assigment

	fmt.Println(a, b, c)

	// multi-variable declaration
	x, y, z := 10, 20, 30

	fmt.Println(x, y, z)

	// swap content of d and e
	d := 1
	e := 2
	d, e = e, d

	fmt.Println(d, e)

	// https://pkg.go.dev/fmt#hdr-Printing

	i := 32
	fmt.Printf("%d, %T\n", i, i)

	var Number int
	fmt.Println("Number=", Number)

	// print value and type of variable j
	j := 2
	fmt.Printf("%d - %T\n", j, j)

	var k1 uint8 = 20
	// k1 = 256 compiler error overflow
	// k1 = -2 compiler error overflow
	fmt.Printf("%d - %T\n", k1, k1)

	var k2 uint16 = 30
	fmt.Printf("%d - %T\n", k2, k2)
	// k1 = k2 //compiler error wrong variable type
	// k2 = k1 compiler error wrong variable type

	k2 = uint16(k1)

	fmt.Printf("%d - %T\n", k2, k2)

}
