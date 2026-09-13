package main

import . "fmt"

// https://github.com/EndlessCheng
func cf1557E() {
	var T int
	var s string
	q := func(x, y int) {
		Println(x, y)
		Scan(&s)
	}
o:
	for Scan(&T); T > 0; T-- {
		x, y := 1, 1
		q(x, y)
		for s != "Done" {
			i := 0
			if y == 1 {
				i = 1
			}
			for ; i <= 8; i++ {
				q(x, y)
				if s == "Done" {
					continue o
				}
				if s[0] == 'D' {
					break
				}
				if s[0] == 'U' {
					if y == 1 {
						i = 1
					}
				}
				y = i
			}
			x++
			q(x, y)
		}
	}
}

//func main() { cf1557E() }
