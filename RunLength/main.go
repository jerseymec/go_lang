package main

import (
	"fmt"
	"strconv"
)

func EncodeF(s string) string {
	res := ""

	s_len := len(s)

	chr := string(s[0])
	if s_len <= 9 {
		return strconv.Itoa(s_len) + chr /// if its less than 9 chrs conv length to str and add chr
	} else {
		for s_len > 0 {
			qo := s_len / 9
			rem := s_len % 9
			if qo == 0 && rem > 0 {
				res += strconv.Itoa(rem) + chr // keep adding multiples of 9 plus chr
				break
			}
			for qo > 0 {

				res += "9" + chr
				qo--
			}

			s_len %= 9

		}
	}
	return res
}
func RunLengthEncoding(str string) string {
	l := len(str)
	var enc string
	subs := ""
	i := 0

	for i < l {
		subs = string(str[i])             // Add first char
		i++                               //increment index
		for i < l && str[i-1] == str[i] { // as long as the current chr matches the last one
			subs += string(str[i]) //create substring
			i++
		}
		if len(subs) > 0 {
			enc += EncodeF(subs) // if there is a subs call Encoder
		}

	}

	return enc
}
func main() {
	s := "AAAAAAAAAAAAABBCCCCDD"
	fmt.Println("Run length encoding is ", RunLengthEncoding(s))
}
