package main

import "fmt"

func Bubblesort(array []int) []int {
	n := len(array)
	for i, _ := range array {
		for j := 1; j < n-i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array
}

func main() {
	nums := []int{3, 7, 1, 9, 2, 20, 5}
	fmt.Println("BubbleSort of ", nums, " is = ", Bubblesort(nums))
	fmt.Println("Is = ", Bubblesort(nums))
}
