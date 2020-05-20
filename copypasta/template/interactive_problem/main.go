package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func run(io func(int) bool) (ans int) {
	var n int
	Scan(&n)

	return
}

func main() {
	io := func(q int) bool {
		Println("?", q)
		// ... or read int and return it
		var s []byte
		Scan(&s)
		return s[0] == 'Y'
	}
	var t int
	for Scan(&t); t > 0; t-- {
		Println("!", run(io))
	}
}
