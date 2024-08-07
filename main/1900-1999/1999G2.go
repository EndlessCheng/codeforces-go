package main

import . "fmt"

func cf1999G2() {
	var T, s int
	for Scan(&T); T > 0; T-- {
		l, r := 1, 999
		for l+1 < r {
			sz := (r - l + 1) / 3
			p, q := l+sz, r-sz
			Println("?", p, q)
			Scan(&s)
			if p*q == s {
				l = q
			} else if p*(q+1) == s {
				l, r = p, q
			} else {
				r = p
			}
		}
		Println("!", r)
	}
}

//func main() { cf1999G2() }
