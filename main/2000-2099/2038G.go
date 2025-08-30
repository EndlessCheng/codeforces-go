package main

import . "fmt"

// https://github.com/EndlessCheng
func cf2038G() {
	q := func(t string) (c int) {
		Println(1, t)
		Scan(&c)
		return
	}
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		Println(0, 1, q("0")-q("00")-q("10")^1)
		Scan(&n)
	}
}

//func main() { cf2038G() }
