package main

import (
	"log"
)

func main() {
	farm(42)
}

//interface has sound method so multiple struct types can meet this methods, allows for future derivations
type Animal interface {
	Sound() string
}

type Dog struct {
	name string
}

func (d Dog) Sound() string {
	return d.name + " says woof"
}

func NewDog() *Dog {
	return &Dog{name: "Doug"}
}

type Cat struct {
	name string
}

func NewCat() Cat {
	return Cat{"Jimmy"}
}

func (c Cat) Sound() string {
	return c.name + " says meow"
}

// using a newDog instantiation method abstracts the object creation knowledge from farm function
// allows the use of complex logic to instantiate object
func farm(x int) {
	var a Animal
	if x > 42 {
		a = NewDog()
	} else {
		a = NewCat()
	}

	log.Print(a.Sound())
}
