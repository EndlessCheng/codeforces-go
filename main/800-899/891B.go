package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF891B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	ans := make([]int, n)
	for i, p := range a[:n-1] {
		ans[p.i] = a[i+1].v
	}
	ans[a[n-1].i] = a[0].v
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF891B(os.Stdin, os.Stdout) }
