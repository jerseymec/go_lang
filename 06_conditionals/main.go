//package _6_conditionals
package main

import "fmt"

func main() {
	x := 10
	y := 10
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)

	}

	//else if
	color := "blue"

	if color == "red" {
		fmt.Println("Color is ", color)
	} else if color == "blue" {
		fmt.Println("Color is ", color)

	} else {
		fmt.Println("Color is neither red no blue")
	}

	//switch

	switch color {
	case "red":
		fmt.Println("Color is ", color)
	case "blue":
		fmt.Println("Color is ", color)
	default:
		fmt.Println("Color is neither red nor blue")

	}

}
