package main

import "fmt"

func main() {
	//Define Map
	emails := make(map[string]string)

	//
	emails["Charles"] = "hellojersey@gmail.com"
	emails["Bob"] = "Bobdhillon@gmail.com"
	emails["Mike"] = "Meike@gmail.com"
	fmt.Println(emails)
	fmt.Println(emails["Charles"])
	fmt.Println(len(emails))

	//delete maps elements
	delete(emails, "Mike")
	fmt.Println(emails)

	emails1 := map[string]string{"Bod": "Bod@gmsil.com", "Sharon": "Sharon@gmail.com"}
	fmt.Println(emails1["Bod"])
	fmt.Println(len(emails1))

}
