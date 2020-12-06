package main

import (
	"fmt"
	"strconv"
)

func runlengthenc(sentence string) string {
	encodedstring := []byte{}
	//prev_char := ''
	currentRun := 1
	for i := 1; i < len(sentence); i++ {
		if sentence[i] != sentence[i-1] || currentRun == 9 {
			encodedstring = append(encodedstring, strconv.Itoa(currentRun)[0])
			encodedstring = append(encodedstring, sentence[i-1])
			currentRun = 0
		}
		currentRun++
	}
	encodedstring = append(encodedstring, strconv.Itoa(currentRun)[0])
	encodedstring = append(encodedstring, sentence[len(sentence)-1])
	return string(encodedstring)

}

func main() {
	inp := "AAAAAAAAAAAAAAAAAAAA"
	fmt.Printf("Encoded version of %s is %s", inp, runlengthenc(inp))
}
