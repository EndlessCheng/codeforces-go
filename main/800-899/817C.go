package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF817C(in io.Reader, out io.Writer) {
	var n, s int64
	Fscan(in, &n, &s)
	c := n - s
	if c < 0 {
		c = 0
	}
	for i := s + 1; i <= n && i < s+200; i++ {
		ds := int64(0)
		for x := i; x > 0; x /= 10 {
			ds += x % 10
		}
		if i-ds < s {
			c--
		}
	}
	Fprint(out, c)
}

//func main() { CF817C(os.Stdin, os.Stdout) }
