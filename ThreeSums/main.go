package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {

	triplets := [][]int{}
	sort.Ints(array)
	for i := 0; i < len(array)-2; i++ {

		left := i + 1
		right := len(array) - 1
		for left < right {
			currentSum := array[i] + array[left] + array[right]

			switch {
			case currentSum < target:
				left++
			case currentSum > target:
				right--
			case currentSum == target:
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left++
				right--
			}

		}
	}

	return triplets
}

func main() {
	nums := []int{12, 3, 1, 2, -6, 5, -8, 6}
	fmt.Println("3 Sums =  ", ThreeNumberSum(nums, 0))
}
