package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	maps := map[rune]int{}
	if str == "" {
		return -1
	}
	for _, r := range str {
		_, err := maps[r]
		if !err {
			maps[r] = 1
		} else {
			maps[r]++
		}
	}
	for i := 0; i < len(str); i++ {
		if maps[rune(str[i])] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(
		FirstNonRepeatingCharacter("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
}
