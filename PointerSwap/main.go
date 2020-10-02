package main

import "fmt"

func pointerswap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	x := 1
	y := 2
	pointerswap(&x, &y)
	fmt.Println("Swap(*x,*y) = ")
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}
