package main

import "fmt"

func add(hit map[string]map[string]int, path, country string) {
	mm, ok := hit[path]
	if !ok {
		mm = make(map[string]int)
		hit[path] = mm
	}
	mm[country]++

}

func main() {
	c_hits := make(map[string]map[string]int)
	add(c_hits, "/main", "US")
	add(c_hits, "/main", "US")
	add(c_hits, "/test", "US")
	add(c_hits, "/main", "UK")
	add(c_hits, "/main", "TN")
	add(c_hits, "/test", "CA")
	add(c_hits, "/test", "CA")
	for k, i := range c_hits {
		for key, v := range i {
			fmt.Println("Path = ", k, " country = ", key, " count = ", v)
		}
	}

}
