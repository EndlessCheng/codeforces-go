package main

import . "fmt"

// https://github.com/EndlessCheng
func cf730B() {
	less := func(i, j int) bool {
		Println("?", i, j)
		c := ""
		Scan(&c)
		return c == "<"
	}
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		mn, mx := 1, 2-n%2
		if n%2 == 0 && less(2, 1) {
			mn, mx = 2, 1
		}
		for i := 3 - n%2; i <= n; i += 2 {
			x, y := i, i+1
			if less(i+1, i) {
				x, y = i+1, i
			}
			if less(x, mn) {
				mn = x
			}
			if less(mx, y) {
				mx = y
			}
		}
		Println("!", mn, mx)
	}
}

//func main() { cf730B() }
