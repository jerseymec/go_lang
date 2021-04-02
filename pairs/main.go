package main

import (
	"fmt"

	"regexp"
)

func Solution(str string) []string {
	n := len(str)
	if n < 2 {
		return []string{str + "_"}
	}
	var splitted []string
	slots := ""
	l := 0
	for _, r := range str {
		if l < 2 {
			slots += string(r)
			l++
		} else {
			splitted = append(splitted, slots)
			slots = string(r)
			l = 1
		}

	}
	if l > 0 {
		if l%2 != 0 {
			slots += string('_')
		}
		splitted = append(splitted, slots)
	}

	return splitted
}

//alternate solution
func splitN(str string, chunk int) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

func splitNreg(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}

func main() {
	fmt.Println(Solution("a"))

}
