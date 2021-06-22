package main

import "fmt"

func LongestPeak(array []int) int {
	var currentpeak int
	n := len(array)
	if n <= 1 {
		return 0
	}
	i := 1
	for i < n-1 {
		if !(array[i-1] < array[i] && array[i] > array[i+1]) {
			i++
			continue
		}
		lindx := i - 2
		rindx := i + 2
		for lindx >= 0 && array[lindx] < array[lindx+1] {
			lindx--
		}
		for rindx < n && array[rindx] < array[rindx-1] {
			rindx++
		}
		if currentpeak < rindx-lindx-1 {
			currentpeak = rindx - lindx - 1
		}
		i = rindx

	}
	return currentpeak

}

func main() {
	ran := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	fmt.Println("longest peak is ", LongestPeak(ran))

}
