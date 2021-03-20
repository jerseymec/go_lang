package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	nu := []int{}
	for i, k := range nums {
		if val, ok := m[k]; ok {
			nu = append(nu, val)
			nu = append(nu, i)
			return nu
		} else {
			m[target-k] = i

		}
	}
	return nil
}

func main() {
	n := []int{1, 3, 4, 5, 6, 7}
	t := 8
	fmt.Println("Indexes of two nums whose sum = ", t, " are ", twoSum(n, t))
}
