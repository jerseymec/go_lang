package main

import "fmt"

func main() {
	ids := []int{1, 23, 12, 33, 42, 12}
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i, id)

	}
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)

	}
}
