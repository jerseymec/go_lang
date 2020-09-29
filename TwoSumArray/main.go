package main

import (
	"fmt"
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	right := len(array) - 1
	left := 0
	sort.Ints(array)
	curr_sum := 0
	for left < right {
		curr_sum = array[left] + array[right]
		if curr_sum == target {
			return []int{array[left], array[right]}
		} else if curr_sum < target {
			left += 1
		} else if curr_sum > target {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	fmt.Println(TwoNumberSum(array, 15))
}
