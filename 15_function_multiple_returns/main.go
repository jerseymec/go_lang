package main

import (
	"fmt"
	"math/rand"
)

func vals() (int, int) {

	return rand.Intn(100), rand.Intn(100)
}
func main() {
	a, b := vals()
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
	_, c := vals()
	fmt.Println("c= ", c)
}
