package main

import "fmt"

func sum(nums...int) int  {

	sum := 0
	for _, val := range nums{
		sum += val

	}
	return sum
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}
func main() {

	n := []int {3,4,5,6}
	fmt.Println("sum = ",sum(n...))
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
