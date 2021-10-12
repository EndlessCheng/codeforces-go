package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1163B2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	c := make([]int, 1e5+1)
	tot := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		c[v]++
		c := c[v]
		tot[c] += c
		if c := tot[c]; c >= i && c != n {
			ans = c
		}
	}
	Fprint(out, ans+1)
}

//func main() { CF1163B2(os.Stdin, os.Stdout) }
