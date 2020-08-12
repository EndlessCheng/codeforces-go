package main

import (
	. "fmt"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1063C() {
	var n int
	var s0, s string
	Scan(&n)
	Println(1, 0)
	Scan(&s0)
	if n == 1 {
		Println(0, 1, 1, 1)
		return
	}
	x := 1
	sort.Search(1e9, func(y int) bool {
		x++
		Println(x, y)
		Scan(&s)
		if x == n {
			yy := y + 1
			if s != s0 {
				yy -= 2
			}
			Println(0, y, x, yy)
			os.Exit(0)
		}
		return s != s0
	})
}

//func main() { CF1063C() }
