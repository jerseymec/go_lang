package main

import (
	"fmt"
	"sort"
)

func main() {
	rh := []int{15, 18, 11, 13, 14}
	bh := []int{6, 9, 2, 4, 5}
	fmt.Println("Class Photo is Possible: ", ClassPhotos(rh, bh))
}

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	// Write your code here.
	sort.Sort(sort.Reverse(sort.IntSlice(redShirtHeights)))
	sort.Sort(sort.Reverse(sort.IntSlice(blueShirtHeights)))
	if blueShirtHeights[0] == redShirtHeights[0] {
		return false
	}
	backrowColor := "Red"
	if blueShirtHeights[0] > redShirtHeights[0] {
		backrowColor = "Blue"

	}

	for indx, _ := range redShirtHeights {

		redshirts := redShirtHeights[indx]
		blueshirts := blueShirtHeights[indx]

		if backrowColor == "Red" {
			if blueshirts >= redshirts {
				return false
			}
		} else {
			if redshirts >= blueshirts {
				return false
			}
		}

	}

	return true
}
