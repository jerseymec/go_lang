package main

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	phone := ""
	if len(numbers) == 10 {
		phone = "(" + strconv.FormatUint(uint64(numbers[0]), 10)
		phone += strconv.FormatUint(uint64(numbers[1]), 10)
		phone += strconv.FormatUint(uint64(numbers[2]), 10) + ") "

		phone += strconv.FormatUint(uint64(numbers[3]), 10)
		phone += strconv.FormatUint(uint64(numbers[4]), 10)
		phone += strconv.FormatUint(uint64(numbers[5]), 10) + "-"
		phone += strconv.FormatUint(uint64(numbers[6]), 10)
		phone += strconv.FormatUint(uint64(numbers[7]), 10)
		phone += strconv.FormatUint(uint64(numbers[8]), 10)
		phone += strconv.FormatUint(uint64(numbers[9]), 10)

		//+ ") "+strconv.FormatUint(numbers[3:],10)+"- "+strconv.FormatUint(numbers[6:],10)

	} else {
		return "invalid"
	}
	return phone

}

//alternate solutions

func formatPhone(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}

func main() {
	a := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//fmt.Println(CreatePhoneNumber(a))
	fmt.Println(formatPhone(a))

}
