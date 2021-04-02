package main

import "fmt"

//Insertsort the smallest element bubbles to the front

func Insertionsort(array []int) []int {
	l := len(array)
	for i := 1; i < l; i++ {
		j := i
		for j > 0 {
			if array[j] < array[j-1] {
				swap(j, j-1, array)
			}
			j--
		}
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	nums := []int{4, 1, 7, 2, 9, 2, 1}
	fmt.Println("Sorted version of ", nums, "is ")
	fmt.Println(Insertionsort(nums))
}
