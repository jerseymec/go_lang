package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C"},
			[]string{"D", "E", "F"},
			[]string{"G", "K", "L"},
		},
	)
	fmt.Println(df)
}
