package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1003A(in io.Reader, out io.Writer) {
	c := [101]int{}
	var n, v, mx int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if c[v]++; c[v] > mx {
			mx = c[v]
		}
	}
	Fprint(out, mx)
}

//func main() { CF1003A(os.Stdin, os.Stdout) }
