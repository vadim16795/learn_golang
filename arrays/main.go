package main

import (
	"fmt"
)

func main() {
	var nums [2]int

	var sum1 int
	var sum2 int

	fmt.Printf("nums=%v type=%T len=%d\n", nums, nums, len(nums))

	for i := range nums {
		sum1 += i
		sum2 += nums[i]
	}

	fmt.Println(sum1, sum2)

	nums2 := [...]int{10, 15, 30}
	fmt.Println(nums2)

	nums2[1] = 20
	fmt.Println(nums2)

	for i := range nums2 {
		sum2 += nums2[i]
	}

	fmt.Println(sum2)

	for j := 0; j < len(nums2); j++ {
		fmt.Print(nums2[j], " ")
	}

}
