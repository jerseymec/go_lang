package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method value reciever
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(lname string) {
	if p.gender == "f" {
		p.lastName = lname
	} else {
		return
	}

}
func main() {
	//initialize a person
	//person1 := Person {firstName: "Charles", lastName: "Smith", city: "New York", gender: "m", age: 24 }
	person1 := Person{"Charles", "Smith", "New York", "f", 24}

	fmt.Println(person1)
	fmt.Println("My name is ", person1.firstName, person1.lastName)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("john")
	fmt.Println(person1.greet())

}
