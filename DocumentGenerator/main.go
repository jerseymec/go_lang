package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	c := map[rune]int{}
	//docu := map[rune]int {}

	for _, chr := range characters {
		if _, ok := c[chr]; !ok {
			c[chr] = 1
		} else {
			c[chr]++
		}
	}
	//for _,chr := range document{
	//	if _,ok := docu[chr];!ok{
	//		docu[chr] = 1
	//	}else{
	//		docu[chr] ++
	//	}
	//}
	//for k,v := range docu{
	//	if cnt,ok := c[k]; !ok{
	//		return false
	//	}else{
	//		if v > cnt{
	//			return false
	//		}
	//	}
	//}
	for _, s := range document {
		if cnt, ok := c[s]; ok && cnt > 0 {
			c[s]--
		} else {
			return false
		}
	}

	return true
}

func main() {
	strs := "Bste!hetsi ogEAxpelrt x @"
	docu := "AlgoExpert is the Best!@"
	fmt.Println("Document can be generated: ", GenerateDocument(strs, docu))

}
