package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF347A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	a[0], a[n-1] = a[n-1], a[0]
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF347A(os.Stdin, os.Stdout) }
