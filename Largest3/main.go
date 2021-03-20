package main

import (
	"fmt"
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		updateLargest(three, num)
	}
	return three
}

func updateLargest(three []int, num int) {
	if num > three[2] {
		shiftAndUpdate(three, num, 2)
	} else if num > three[1] {
		shiftAndUpdate(three, num, 1)
	} else if num > three[0] {
		shiftAndUpdate(three, num, 0)
	}
}

func shiftAndUpdate(three []int, num, ind int) {
	for i := 0; i < ind+1; i++ {
		if i == ind {
			three[i] = num
		} else {
			three[i] = three[i+1]
		}
	}
}

func main() {
	test_array := []int{1, 6, 2, 5, 9, 10, 3}
	fmt.Println("The 3 largest = ", FindThreeLargestNumbers(test_array))
}
