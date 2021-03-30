package main

import "fmt"

func FindOutlier(integers []int) int {
	odd := 0
	even := 0
	l := len(integers)
	if l < 3 {
		return 0
	}
	for i, n := range integers {
		if n%2 == 0 {
			even++
		} else {
			odd++
		}
		if i >= 3 {
			break
		}
	}
	if even > odd {
		for _, n := range integers {
			if n%2 != 0 {
				return n
			}
		}
	} else {
		for _, n := range integers {
			if n%2 == 0 {
				return n
			}
		}

	}
	return 0
}

func FindOutlier2(integers []int) int {
	var even, odd *int
	for i, integer := range integers {
		if even != nil && odd != nil {
			if integer%2 == 0 {
				return *odd
			} else {
				return *even
			}
		}
		if integer%2 == 0 {
			even = &integers[i]
		} else {
			odd = &integers[i]
		}

	}
	return integers[len(integers)-1]
}

func main() {
	ints := []int{2, 4, 8, 6, 5, 2, 6, 8}
	fmt.Println("Outlier is ", FindOutlier2(ints))
	fmt.Println("Outlier is ", FindOutlier(ints))
}
