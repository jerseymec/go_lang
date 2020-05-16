package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("a= %d, b= %d\n", a, b)
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use * to read val from address
	fmt.Println("b= ", *b)
	*b = 10

	fmt.Println(a)
	fmt.Println(*b)
}
