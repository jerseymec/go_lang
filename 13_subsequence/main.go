package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	arr_n := len(array)
	seq_n := len(sequence)
	if (arr_n <= 0) || (seq_n <= 0) {
		return false

	} else {
		j := 0
		for i, seq := range sequence {
			for (j < arr_n) && (array[j] != seq) {
				j++
			}
			if (j >= arr_n) && (i < seq_n) {
				return false
			}
			j++

		}

	}

	return true
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{5, 1, 22, 25, 6, -1, 8, 10, 12}
	fmt.Println(IsValidSubsequence(array, sequence))
}
