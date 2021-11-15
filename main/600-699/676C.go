package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF676C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, l, mxC int
	var s string
	Fscan(in, &n, &k, &s)
	c := [2]int{}
	for r, b := range s {
		b &= 1
		c[b]++
		if c[b] > mxC {
			mxC = c[b]
		}
		if r-l+1-mxC > k {
			c[s[l]&1]--
			l++
		}
	}
	Fprint(out, n-l)
}

//func main() { CF676C(os.Stdin, os.Stdout) }
