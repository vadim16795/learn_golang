package main

import "fmt"

func main() {
	black := "#"
	white := " "
	finalblack := ""
	finalwhite := ""
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			finalblack = finalblack + black
			finalwhite = finalwhite + white
		} else {
			finalblack = finalblack + white
			finalwhite = finalwhite + black
		}
	}
	// fmt.Println(finalblack)
	// fmt.Println(finalwhite)
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			fmt.Println(finalblack)
		} else {
			fmt.Println(finalwhite)
		}
	}

}

// print("stackoverflow")

// def chessboard(n):
//     finalSentence1 = ""
//     finalSentence2 = ""
//     for i in range(n): #we add 0 and 1 as much as we have n
//         if i%2 == 0: #
//             finalSentence1 += "1"
//             finalSentence2 += "0"
//         else:
//             finalSentence1 += "0"
//             finalSentence2 += "1"

//     for i in range(n): #we print as much as we have n
//         if i%2 == 0:
//             print(finalSentence1)
//         else:
//             print(finalSentence2)

// chessboard(3)
