package main

import (
	"fmt"
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	var sortedInterval = make([][]int, len(intervals))
	copy(sortedInterval, intervals)
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(sortedInterval, func(i, j int) bool {
		return sortedInterval[i][0] < sortedInterval[j][0]
	})
	var merged [][]int
	merged = append(merged, sortedInterval[0])
	m := 1
	for i := 1; i < len(intervals); i++ {

		if merged[m-1][1] >= sortedInterval[i][0] {
			merged[m-1][1] = max(merged[m-1][1], sortedInterval[i][1])
		} else {
			merged = append(merged, sortedInterval[i])
			m++
		}

	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func main() {
	//intervals := [][]int{
	//{1, 2},
	//{-3, 5},
	//{4, 7},
	//{6, 8},
	//{9, 10},
	//}

	interval2 := [][]int{
		{89, 90},
		{-10, 20},
		{-50, 0},
		{70, 90},
		{90, 91},
		{90, 95},
	}
	fmt.Println("Final merged intervals are : ", MergeOverlappingIntervals(interval2))

}
