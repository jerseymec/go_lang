package main

import "fmt"

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func MoveElementToEnd(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1
	n := len(array)

	for left < right {
		for left < n-1 && array[left] != toMove {
			left++
		}
		for right >= 0 && array[right] == toMove {
			right--
		}
		if left < right {
			swap(left, right, array)
		}

	}
	return array
}

func main() {
	array2 := []int{52, 17, 30, 19, 2, 2, 2, -13, -10}
	fmt.Println("Move 2 to the end ", MoveElementToEnd(array2, 44))

}
