package main

import "fmt"

func isMonotonic(A []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			decreasing = false
		}
		if A[i] < A[i-1] {
			increasing = false
		}
	}
	return increasing || decreasing
}

func main() {
	arr := []int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 11}
	fmt.Println(isMonotonic(arr))
}
