package main

import "fmt"

func main() {
	s := []int{-3, 1, 2, 3, 4}
	fmt.Println("Sorted suqred list =  ", SortedSquaredArray(s))
}

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	r := len(array) - 1
	indx := r
	l := 0
	sqrd := make([]int, len(array))
	for indx >= 0 {
		if abs(array[l]) >= abs(array[r]) {
			sqrd[indx] = array[l] * array[l]
			indx--
			l++
		} else {
			sqrd[indx] = array[r] * array[r]
			indx--
			r--
		}
	}
	return sqrd
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
