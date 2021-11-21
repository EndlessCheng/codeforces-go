package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF300A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fprintln(out, 1, a[0])
	if a[n-1] > 0 {
		Fprintln(out, 1, a[n-1])
		Fprint(out, n-2)
		for _, v := range a[1 : n-1] {
			Fprint(out, " ", v)
		}
	} else {
		Fprintln(out, 2, a[1], a[2])
		Fprint(out, n-3)
		for _, v := range a[3:] {
			Fprint(out, " ", v)
		}
	}
}

//func main() { CF300A(os.Stdin, os.Stdout) }
