package main

import (
	. "fmt"
)

// https://github.com/EndlessCheng
func cf2013C() {
	q := func(t string) (ok bool) {
		Println("?", t)
		Scan(&ok)
		return
	}
	var T, n int
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		s := ""
		for len(s) < n {
			if q(s + "0") {
				s += "0"
			} else if q(s + "1") {
				s += "1"
			} else {
				break
			}
		}
		for len(s) < n {
			if q("0" + s) {
				s = "0" + s
			} else {
				s = "1" + s
			}
		}
		Println("!", s)
	}
}

//func main() { cf2013C() }
