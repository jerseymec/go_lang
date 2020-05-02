package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Printf("%T\n", b)

	var d = true
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	var e int
	fmt.Println(e)

	var age int32 = 23
	size1 := 45.67
	var size float32 = 2.3
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", size1)

	f := "apple"
	fmt.Println(f)
}
