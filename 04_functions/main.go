package main

import "fmt"

func greetings(name string) string {
	return "Hello " + name
}
func sums(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greetings("Charles"))
	fmt.Println(sums(2, 3))
}
