package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("Good Day"))
	fmt.Println("Good Day"[5]) // ascii representation of D

	s := "SuperHelpful"

	fmt.Println(s[0:5]) // Super
	fmt.Println(s[5:])

	for i := 0; i <= len(s); i++ {
		println(s[i:])
	}

	for i := 0; i <= len(s); i++ {
		println(s[:i])
	}
	fmt.Println("####")
	fmt.Println(s[:])

	s += "Friend"
	fmt.Println(s)
}
