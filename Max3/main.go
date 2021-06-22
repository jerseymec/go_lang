package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {

	diction := map[int]bool{}

	for _, i := range nums {

		diction[i] = true

	}

	list_nums := []int{}

	for key := range diction {

		list_nums = append(list_nums, key)

	}

	if len(list_nums) <= 2 {

		return max(list_nums)

	}
	sort.Sort(sort.Reverse(sort.IntSlice(list_nums)))
	return list_nums[2]

}

func max(a []int) int {

	max := -10000000

	for _, i := range a {

		if i > max {

			max = i

		}

	}

	return max

}
func main() {
	n := []int{1, 2, 1, 2, 3, 1}
	fmt.Println(thirdMax(n))
}
