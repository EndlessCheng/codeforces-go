package main

import (
	"bufio"
	. "fmt"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(n int, Q func(int64) bool) (ans int64) {

	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	Q := func(q int64) bool {
		Println("?", q)
		// ... or read int and return it
		var s []byte
		Fscan(in, &s)
		return s[0] == 'Y'
	}
	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n int
		Fscan(in, &n)
		Println("!", run(n, Q))
	}
}
