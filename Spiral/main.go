package main

import "fmt"

func spiraltraverse(array [][]int) []int {

	if len(array) <= 0 {
		return []int{}
	}

	var results = []int{}
	startRow := 0
	endRow := len(array) - 1
	startCol := 0
	endCol := len(array[0]) - 1

	for startRow <= endRow && startCol <= endCol {
		for c := startCol; c <= endCol; c++ {
			results = append(results, array[startRow][c])
		}
		for r := startRow + 1; r <= endRow; r++ {
			results = append(results, array[r][endCol])
		}
		for c := endCol - 1; c >= startCol; c-- {
			if endRow == startRow {
				break
			}
			results = append(results, array[endRow][c])
		}

		for r := endRow - 1; r > startRow; r-- {
			if startCol == endCol {
				break
			}
			results = append(results, array[r][startCol])
		}
		startRow++
		endRow--
		startCol++
		endCol--
	}
	return results

}
func main() {
	ab := [][]int{
		{1, 2, 3, 4, 5},
		{14, 15, 16, 17, 6},
		{13, 20, 19, 18, 7},
		{12, 11, 10, 9, 8},
	}
	fmt.Println("Spiral read of matrix = ", spiraltraverse(ab))
}
