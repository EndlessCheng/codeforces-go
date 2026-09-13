package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1557E() {
	var T, p int
	var s string
	q := func(x, y int) string {
		Println(x, y)
		p = y
		Scan(&s)
		return s
	}
	var solve func(int) bool
	solve = func(i int) bool {
		j := 1
		if p == 1 {
			j = 2
		}
		for ; j <= 8; j++ {
			s = q(i, j)
			if s == "Done" {
				return true
			} else if s[0] == 'D' {
				return false
			} else if s[0] == 'U' {
				return solve(i)
			}
		}
		return false
	}

	for Scan(&T); T > 0; T-- {
		p = 1
		for i := 1; i <= 8; i++ {
			if q(i, p) == "Done" || solve(i) {
				break
			}
		}
	}
}

//func main() { cf1557E() }
