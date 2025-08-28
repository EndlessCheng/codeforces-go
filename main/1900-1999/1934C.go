package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1934C() {
	q := func(x, y int) (r int) {
		Println("?", x, y)
		Scan(&r)
		return
	}
	var T, n, m int
	for Scan(&T); T > 0; T-- {
		Scan(&n, &m)
		d := q(1, 1)
		x, y := d+1, 1
		if d >= n {
			x, y = n, d-n+2
		}
		e := q(x, y)
		if q(x-e/2, y+e/2) == 0 {
			Println("!", x-e/2, y+e/2)
			continue
		}
		if d < m {
			x, y = 1, d+1
		} else {
			x, y = d-m+2, m
		}
		e = q(x, y)
		Println("!", x+e/2, y-e/2)
	}
}

//func main() { cf1934C() }
