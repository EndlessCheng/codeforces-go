package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF25A(in io.Reader, out io.Writer) {
	var n, c, tar int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] &= 1
	}
	for _, v := range a[:3] {
		c += v
	}
	if c < 2 {
		tar = 1
	}
	for i, v := range a {
		if v == tar {
			Fprint(out, i+1)
		}
	}
}

//func main() { CF25A(os.Stdin, os.Stdout) }
