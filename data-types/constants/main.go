package main

import (
	"fmt"
)

func main() {

	var k1 uint8 = 20
	k1 /= 2
	const pi = 3.14159

	fmt.Println(25, pi, k1)

	const (
		a = 2
		b
		c
		d = a * 10
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)

	// iota = const generator

	const (
		a2 = iota + 1
		b2
		c2
		d2 = a2 * 10
		e2
		f2 = iota
	)
	fmt.Println(a2, b2, c2, d2, e2, f2)
}
