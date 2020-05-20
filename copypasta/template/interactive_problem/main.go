package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func run(io func(int64) bool) (ans int64) {
	var n int
	Scan(&n)

	return
}

func main() {
	io := func(q int64) bool {
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
