package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	signed := 0

	if x*-1 > 0 {
		signed = -1
		x *= -1
	} else {
		signed = 1
	}
	str_int := strconv.Itoa(x)
	r := ""
	for _, c := range str_int {
		r = string(c) + r
	}
	reversed_int := string(r)
	if reversed, err := strconv.Atoi(reversed_int); err != nil {
		return 0
	} else {
		if reversed > math.MaxInt32 {
			return 0
		} else {
			return reversed * signed
		}
	}

}

func reversed1(x int) int {
	var ans int64 = 0
	for x != 0 {
		ans = ans*10 + int64(x%10)
		x = x / 10
	}
	x = int(ans)
	if math.MinInt32 > x || x > math.MaxInt32 {
		return 0
	}
	return x
}
func main() {
	n := -989898986469
	//n := 1534236469
	fmt.Println(reversed1(n))
}
