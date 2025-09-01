package main

import (
	. "fmt"
)

// https://github.com/EndlessCheng
func cf1729E() {
	q := func(a, b int) (d int) {
		Println("?", a, b)
		Scan(&d)
		return
	}
	for b := 2; ; b++ {
		d := q(1, b)
		if d < 0 {
			Println("!", b-1)
			return
		}
		d2 := q(b, 1)
		if d != d2 {
			Println("!", d+d2)
			return
		}
	}
}

//func main() { cf1729E() }
