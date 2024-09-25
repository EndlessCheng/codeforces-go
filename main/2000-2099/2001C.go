package main

import . "fmt"

// https://github.com/EndlessCheng
func cf2001C() {
	var T, n, m int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		ans := []any{"!"}
		ok := make([]bool, n+1)
		for i := 2; n > 1; n-- {
			for ok[i] {
				i++
			}
			v, w := 1, i
			for {
				Println("?", v, w)
				Scan(&m)
				if m == v {
					ok[w] = true
					ans = append(ans, v, w)
					break
				}
				if ok[m] {
					v = m
				} else {
					w = m
				}
			}
		}
		Println(ans...)
	}
}

//func main() { cf2001C() }
