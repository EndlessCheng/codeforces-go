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
	Q := func(q int64) (resp bool) {
		Println("?", q)
		// ... or read int and return it
		var s []byte
		Fscan(in, &s)
		resp = s[0] == 'Y'
		return
	}
	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n int
		Fscan(in, &n)
		ans := run(n, Q)
		Println("!", ans)
		// some problems need to read an extra string like "Correct" or "Incorrect" after guessed the answer
		//var s []byte
		//Fscan(in, &s)
	}
}
