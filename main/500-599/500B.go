package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF500B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([]string, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	vis := make([]bool, n)
	ids := []int{}
	var f func(int)
	f = func(v int) {
		ids = append(ids, v)
		vis[v] = true
		for w, b := range g[v] {
			if b == '1' && !vis[w] {
				f(w)
			}
		}
	}
	for i, b := range vis {
		if !b {
			ids = nil
			f(i)
			sort.Ints(ids)
			b := []int{}
			for _, id := range ids {
				b = append(b, a[id])
			}
			sort.Ints(b)
			for j, id := range ids {
				a[id] = b[j]
			}
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF500B(os.Stdin, os.Stdout) }
