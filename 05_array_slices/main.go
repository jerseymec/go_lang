package main

import "fmt"

func main() {
	//Arrays
	var friutArr [2]string

	//Assign values
	friutArr[0] = "Apples"
	friutArr[1] = "Orange"

	fruits := [2]string{"Apples", "Oranges"}

	fmt.Println(friutArr)
	fmt.Println(friutArr[1])
	fmt.Println(fruits)
}
