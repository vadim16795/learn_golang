package main

import (
	"fmt"
)

func main() {

	Hello()

	var k1 uint8 = 20
	f1 := 3.14

	fmt.Printf("%f - %T\n", f1, f1)          // %f - float64 value
	fmt.Printf("%5.3f %.2f \n", f1, 214.437) // %5.3f 5digits overall , 3 digits after . (decimal) , %.2f only care about digits after decimal point

	fmt.Println(k1+12, k1-12, k1*2, k1/2, 5%3)

	k1++ // k1 = k1 + 1
	fmt.Println(k1)

	k1 += 10 // k1 = k1 + 10
	fmt.Println(k1)

	fmt.Println(5/2, 5.0/2.0)
}

func Hello() {

	println("Hello,World!")
}
