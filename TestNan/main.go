package main

import "math"

var(
	a = math.Sqrt(-1.0) //NaN
	x = 0.0
	y = 1.0/x //+Inf
	z = 2.0*y
)

func main() {
	m:= map[float64] int{}
	m[a] = 1;m[a]=1;m[a]=1
	m[y] = 2;m[z]=2
	print(len(m))

}
