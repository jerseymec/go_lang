package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	filesize := 4000000000.
	fmt.Println("GB = ", GB)
	fmt.Println("KB = ", KB)

	fmt.Printf("filesize = %.2fGB", filesize/GB)

}
