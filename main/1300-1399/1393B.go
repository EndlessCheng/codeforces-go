package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1393B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v int
	var op string
	c := [1e5 + 1]int{}
	sufC := [2e5 + 1]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v]++
		sufC[c[v]]++
	}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op, &v); op == "+" {
			c[v]++
			sufC[c[v]]++
		} else {
			sufC[c[v]]--
			c[v]--
		}
		if sufC[8] > 0 || sufC[6] > 0 && sufC[2] > 1 || sufC[4] > 1 || sufC[4] > 0 && sufC[2] > 2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1393B(os.Stdin, os.Stdout) }
