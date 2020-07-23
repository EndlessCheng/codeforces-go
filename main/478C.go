package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF478C(in io.Reader, out io.Writer) {
	min := func(a ...int64) int64 {
		res := a[0]
		for _, v := range a[1:] {
			if v < res {
				res = v
			}
		}
		return res
	}

	var r, g, b int64
	Fscan(in, &r, &g, &b)
	Fprint(out, min((r+g+b)/3, r+g, r+b, g+b))
}

//func main() { CF478C(os.Stdin, os.Stdout) }
