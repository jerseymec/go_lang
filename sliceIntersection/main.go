package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	elem := 0
	for i < len(nums1) && j < len(nums2) {
		switch {
		case nums1[i] == nums2[j]:
			elem = nums1[i]
			res = append(res, elem)
			for i < len(nums1) && nums1[i] == elem {
				i++
			}
			for j < len(nums2) && nums2[j] == elem {
				j++
			}
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		}
	}
	return res
}

func intersection2(nums1, nums2 []int) []int {
	tag := make(map[int]int)
	var comm []int
	for _, num := range nums1 {
		tag[num] = 1
	}
	for _, n := range nums2 {
		if tag[n] == 1 {
			comm = append(comm, n)
			tag[n] = 2
		}
	}
	return comm
}
func main() {
	a := []int{4, 9, 5}
	b := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection2(a, b))
}
