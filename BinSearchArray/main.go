package main

import "fmt"

func BinarySearch(array []int, target int) int {
	// Write your code here.
	r := len(array) - 1
	return BinSearch(array, 0, r, target)
}

func BinSearch(list []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if list[mid] > target {
		return BinSearch(list, left, mid-1, target)
	} else if list[mid] < target {
		return BinSearch(list, mid+1, right, target)
	}

	return mid
}

func NonRecursiveBin(array []int, target int) int {
	l := 0
	r := len(array) - 1

	for l <= r {
		mid := (l + r) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 5, 7, 9, 10, 12, 14, 15}
	//fmt.Println("Index of 1 is:  ", BinarySearch(nums,1))
	fmt.Println("Index of 1 is:  ", NonRecursiveBin(nums, 8))
}
