package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

func ProductSum(array SpecialArray) int {
	// Write your code here.
	return PS(array, 1)
}
func PS(array SpecialArray, mult int) int {
	sum := 0
	for _, elem := range array {
		if ct, ok := elem.(SpecialArray); ok {
			sum += PS(ct, mult+1)
		} else if ct, ok := elem.(int); ok {
			sum += ct
		}
	}
	return sum * mult
}

func main() {
	input := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	fmt.Println("Sum = ", ProductSum(input))
}
