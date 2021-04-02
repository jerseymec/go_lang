package main

import "fmt"

func IsPalindrome(str string) bool {
	// Write your code here.
	n := len(str)
	if n <= 1 {
		return true
	}
	return str[0] == str[n-1] && IsPalindrome(str[1:n-1])
}

func isPalindrome2(str string) bool {
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		r--
		l++
	}
	return true
}

func main() {
	//str := "abcdcba"
	str := "abcddcba"
	fmt.Println("Is Palindrome: ", IsPalindrome(str))
	fmt.Println("Is Palindrome: ", isPalindrome2(str))
}
