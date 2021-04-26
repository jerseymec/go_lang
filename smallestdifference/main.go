package main

import (
	"fmt"
	"math"
	"sort"
)

//func SmallestDifference(array1, array2 []int) []int {
//
//	if len(array1) <=0{
//		return nil
//	}
//	if len(array2) <= 0{
//		return nil
//	}
//	var min1, min2 int
//	sort.Ints(array1)
//	sort.Ints(array2)
//
//	minDiff := math.MaxInt64
//	for i := 0; i < len(array1); i++ {
//		int2 := 0
//		for int2 < len(array2) {
//			switch minDiff > int(math.Abs(float64(array1[i]-array2[int2]))) {
//			case true:
//				minDiff = int(math.Abs(float64(array1[i] - array2[int2])))
//				min1 = array1[i]
//				min2 = array2[int2]
//
//			}
//			int2++
//
//		}
//
//	}
//	return []int{min1,min2}
//}

func SmallestDifference(array1, array2 []int) []int {

	if len(array1) <= 0 {
		return nil
	}
	if len(array2) <= 0 {
		return nil
	}
	var min1, min2 int
	sort.Ints(array1)
	sort.Ints(array2)

	minDiff := math.MaxInt64
	i, j := 0, 0

	for i < len(array1) && j < len(array2) {
		firstnum := array1[i]
		secondnum := array2[j]
		switch {
		case firstnum < secondnum:
			i++
		case secondnum < firstnum:
			j++
		case firstnum == secondnum:
			return []int{firstnum, secondnum}

		}
		if minDiff > int(math.Abs(float64(firstnum-secondnum))) {
			minDiff = int(math.Abs(float64(firstnum - secondnum)))
			min1 = array1[i]
			min2 = array2[j]
		}

	}

	return []int{min1, min2}
}

func main() {
	//array1 := []int {12, 3, 1, 2, -6, 15, -8, 16}
	//array2 := []int {5, 10, 22, 25, 6, -1, 8, 10}
	//array1 := []int {10, 0, 20, 25}
	//array2 := []int {1005, 1006, 1014, 1032, 1031}
	array1 := []int{5, 7, 3, 9, 4, 8, 3, 1}
	array2 := []int{52, 17, 30, 19, 2, 12, -20, -13, -10}

	fmt.Println("Smallest Diff pairs = ", SmallestDifference(array1, array2))

}
