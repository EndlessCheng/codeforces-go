package main

import (
	. "fmt"
	"sort"
)

// https://github.com/EndlessCheng
func cf1698D() {
	var T, n, v int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		Println("!", 1+sort.Search(n-1, func(m int) bool {
			m++
			Println("?", 1, m)
			cnt := 0
			for i := 0; i < m; i++ {
				Scan(&v)
				if v <= m {
					cnt++
				}
			}
			return cnt%2 > 0
		}))
	}
}

//func main() { cf1698D() }
