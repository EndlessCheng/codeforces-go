package main

import (
	. "fmt"
	"math/bits"
)

func cf1934D2() {
	var T, n, p, q uint
	r := func() {
		Scan(&p, &q)
		if bits.OnesCount(p)%2 == 0 {
			n = p
		} else {
			n = q
		}
	}
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		if bits.OnesCount(n)%2 > 0 {
			Println("second")
			r()
		} else {
			Println("first")
		}
		for n > 0 {
			hb := uint(1) << (bits.Len(n) - 1)
			Println(n^hb, hb)
			r()
		}
	}
}

//func main() { cf1934D2() }
