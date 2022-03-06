package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1395C(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

O:
	for m := 0; ; m++ {
	o:
		for _, v := range a {
			for _, w := range b {
				if v&w|m == m {
					continue o
				}
			}
			continue O
		}
		Fprint(out, m)
		return
	}
}

//func main() { CF1395C(os.Stdin, os.Stdout) }
