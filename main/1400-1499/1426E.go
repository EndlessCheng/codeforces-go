package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1426E(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var n, a, b, c, x, y, z int
	Fscan(in, &n, &a, &b, &c, &x, &y, &z)
	Fprint(out, max(max(max(0, a-x-z), b-x-y), c-y-z), min(a, y)+min(b, z)+min(c, x))
}

//func main() { CF1426E(os.Stdin, os.Stdout) }
