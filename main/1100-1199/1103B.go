package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1103B() {
	var s string
	q := func(x, y int) bool {
		Println("?", x, y)
		Scan(&s)
		return s == "y"
	}
	for Scan(&s); s != "end"; Scan(&s) {
		i := 1
		for q(i/2, i) {
			i *= 2
		}
		l, r := i/2, i
		for l+1 < r {
			m := (l + r) / 2
			if q(m, l) {
				r = m
			} else {
				l = m
			}
		}
		Println("!", r)
	}
}

//func main() { cf1103B() }
