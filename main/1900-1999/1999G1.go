package main

import (
	. "fmt"
	"sort"
)

func cf1999G1() {
	var T, s int
	for Scan(&T); T > 0; T-- {
		ans := 2 + sort.Search(997, func(x int) bool {
			x += 2
			Println("?", x, x)
			Scan(&s)
			return s > x*x
		})
		Println("!", ans)
	}
}

//func main() { cf1999G1() }
