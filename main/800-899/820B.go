package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF820B(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var s, a int
	Fscan(in, &s, &a)
	Fprint(out, 2, 1, min(2+max(1, (a*s+90)/180), s))
}

//func main() { CF820B(os.Stdin, os.Stdout) }
