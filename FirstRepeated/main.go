package main

import "fmt"

func FirstDuplicateValue(array []int) int {
	var repeats = map[int]int{}

	for _, a := range array {
		if _, ok := repeats[a]; ok {
			return a
		} else {
			repeats[a]++
		}

	}
	return -1
}

func main() {
	num := []int{2, 1, 5, 12, 3, 13, 4}
	fmt.Println("First repeating num = ", FirstDuplicateValue(num))
}
