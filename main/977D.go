package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF977D(in io.Reader, out io.Writer) {
	g := [64][]int64{}
	var n int
	var v int64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c := 0
		for w := v; w&1 == 0; w >>= 1 {
			c++
		}
		g[c] = append(g[c], v)
	}
	for _, a := range g {
		sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		for _, v := range a {
			Fprint(out, v, " ")
		}
	}
}

//func main() { CF977D(os.Stdin, os.Stdout) }
