package main

import "fmt"

//O(N) time | O(N) Space
func ArrayOfProducts(array []int) []int {
	n := len(array)
	products := make([]int, n)
	for i := 0; i < n; i++ {
		products[i] = 1
	}

	leftrunning := 1
	for i := 0; i < n; i++ {
		products[i] = leftrunning
		leftrunning *= array[i]
	}
	rightrunning := 1
	for i := n - 1; i >= 0; i-- {
		products[i] *= rightrunning
		rightrunning *= array[i]
	}

	return products
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println("Array of Products = ", ArrayOfProducts(a))

}
