package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1342A(_r io.Reader, _w io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	var t, x, y, a, b int64
	for Fscan(_r, &t); t > 0; t-- {
		Fscan(_r, &x, &y, &a, &b)
		if x > y {
			x, y = y, x
		}
		Fprintln(_w, min(a*(x+y), a*(y-x)+b*x))
	}
}

//func main() { CF1342A(os.Stdin, os.Stdout) }
