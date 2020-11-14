package main

import "fmt"

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	a := 5
	b := &a

	fmt.Printf("a= %d, b= %d\n", a, b)
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use * to read val from address
	fmt.Println("b= ", *b)
	*b = 10
	i := 5
	fmt.Println("i = ", i)

	zeroptr(&i)
	fmt.Println("i = ", i)

	fmt.Println(a)
	fmt.Println(*b)
}
