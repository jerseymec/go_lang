package main

import (
	"fmt"
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	// Write your code here.

	sort.Ints(queries)
	if len(queries) <= 1 {
		return 0
	}

	if len(queries) < 2 {
		return queries[0]
	}

	sum := queries[0]

	for i := 2; i < len(queries); i++ {
		fmt.Println(queries[i])
		for j := 0; j < i; j++ {
			sum += queries[j]
		}

	}
	return sum
}

func main() {
	array := []int{1, 2, 3, 4, 6}
	fmt.Println("Minimum Wait Time = ", MinimumWaitingTime(array))
}
