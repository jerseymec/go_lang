package main

import "fmt"

func largestUniqueNumber(A []int) int {
	cnter := map[int]int{}
	for _, n := range A {
		if _, ok := cnter[n]; !ok {
			cnter[n] = 1
		} else {
			cnter[n] += 1
		}
	}
	mx := -1
	for k, v := range cnter {
		if k > mx && v == 1 {
			mx = k
		}
	}
	return mx

}
func main() {
	//a := []int{5,7,3,9,4,9,8,3,1}
	a := []int{9, 9, 8, 8}
	fmt.Println("Largest unique: ", largestUniqueNumber(a))
}
