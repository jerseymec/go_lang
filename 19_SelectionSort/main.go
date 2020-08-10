package main

import "fmt"

func SelectionSort(array []int) []int {

	len_arr := len(array)
	elem := 0
	for i := 0; i < len_arr; i++ {
		for j := i + 1; j < len_arr; j++ {
			if array[j] < array[i] {
				elem = array[i]
				array[i] = array[j]
				array[j] = elem
			}
		}
	}
	return array
}

func main() {
	seq := []int{2, 4, 52, 5, 8, 9}
	fmt.Println("Sort of [2,4,52,5,8,9] = ", SelectionSort(seq))
}
