package main

import (
	"fmt"
)

func caesars(str string, k int) string {
	k = k % 26
	var res []rune
	var num int
	for _, c := range str {
		switch {
		case 97 <= c && c <= 122:
			num = int(c) + k

			if num > 122 {
				num = 96 + num%122
			}
			res = append(res, rune(num))

		case 65 <= c && c <= 90:
			num = int(c) + k

			if num > 90 {
				num = 64 + num%90
			}
			res = append(res, rune(num))
		default:
			fmt.Printf("%s is not an alphabet", c)
		}
	}
	return string(res)
}
func main() {
	str := "xyz%"
	fmt.Println("encrypted: ", caesars(str, 2))
}
