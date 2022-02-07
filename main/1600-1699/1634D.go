package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func q(i, j, k int) (d int) {
	Println("?", i, j, k)
	Scan(&d)
	return
}

func CF1634D() {
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		a, b, c := 1, 2, 3
		d := q(1, 2, 3)
		for i := 4; i <= n; i++ {
			x, y := q(a, b, i), q(a, c, i)
			if x <= d && y <= d {
				continue
			}
			if y > x {
				b, d = i, y
			} else if x > d {
				c, d = i, x
			}
		}
		i := 1
		for i == a || i == b || i == c {
			i++
		}
		if q(a, b, i) == d {
			Println("!", a, b)
		} else if q(a, c, i) == d {
			Println("!", a, c)
		} else {
			Println("!", b, c)
		}
	}
}

//func main() { CF1634D() }
