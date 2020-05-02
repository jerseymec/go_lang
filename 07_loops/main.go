package main

import "fmt"

func main() {
	//long method
	i := 1
	for i <= 10 {
		i++
		fmt.Println("i = ", i)
	}

	//short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)

	}

	//fizzbuzz

	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Printf("%d : FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d : Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d : Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
