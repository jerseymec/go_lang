package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println("Minimum coin value that cannot be constructed = ", NonConstructibleChange(coins))
}

func NonConstructibleChange(coins []int) int {
	// if coin > currentsum +1 then you will not be able to generate currentsum+1 value using previous coins
	sort.Ints(coins)
	currentChangeCreated := 0
	for _, coin := range coins {
		if coin > currentChangeCreated+1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}
	return currentChangeCreated + 1
}
