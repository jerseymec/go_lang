package main

import "fmt"

func main() {
	ids := []int{1, 23, 12, 33, 42, 12}

	//using index
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i, id)

	}
	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)

	}
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum : ", sum)

	emails := map[string]string{"Bod": "Bod@gmsil.com", "Sharon": "Sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name : ", k)
	}

}
