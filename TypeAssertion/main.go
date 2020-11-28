package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

type Values struct {
	Now time.Time
}

var (
	myString = "Hello World"
	myInt    = 10
	myBool   = false
	myStruct = Values{Now: time.Now()}
)

func main() {
	// f := run
	// f := assertString
	// f := assertStringNoPanic
	f := assertReader
	if err := f(myStruct); err != nil {
		// if err := f(myString); err != nil {
		log.Fatal("main : error : ", err)
	}
}
func assertReader(data interface{}) error {
	_ = data.(io.Reader)
	fmt.Println("Data implements the io.Reader interface")
	return nil

}

func run(data interface{}) error {
	// str := data.(string)
	fmt.Println("My Data: ", data)
	return nil
}

func assertString(data interface{}) error {
	str := data.(string)
	fmt.Println("My Data : ", str)
	return nil
}

func assertStringNoPanic(data interface{}) error {
	str, ok := data.(string)
	if !ok {
		return fmt.Errorf("Invalid type: data has %T", data)
	}
	fmt.Println("My Data : ", str)
	return nil
}
