package main

import (
	"errors"
	"fmt"
	"log"
)

func doError() (string, error) {
	return "", errors.New("This is my error")
}

func doNoError() (string, error) {
	return "My response", nil
}

func doFmtError() error {
	errCode := 401
	return fmt.Errorf("This is ny error code: %d", errCode)
}
func main() {
	resp, err := doError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)

	}
	fmt.Println("My Message: ", resp)
	resp, err = doNoError()
	if err != nil {
		log.Printf("This should not print")
	}
	fmt.Println("My Response: ", resp)
	err = doFmtError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)
	}
}
