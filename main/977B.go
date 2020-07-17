package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF977B(in io.Reader, out io.Writer) {
	var n int
	var s, m string
	Fscan(in, &n, &s)
	c := map[string]int{}
	for i := 2; i <= n; i++ {
		t := s[i-2 : i]
		if c[t]++; c[t] > c[m] {
			m = t
		}
	}
	Fprint(out, m)
}

//func main() { CF977B(os.Stdin, os.Stdout) }
