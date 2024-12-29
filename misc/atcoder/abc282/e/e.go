package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type tuple struct{ v, w, wt int }
	b := make([]tuple, 0, n*n)
	for i, v := range a {
		for j := i + 1; j < n; j++ {
			w := a[j]
			b = append(b, tuple{i, j, (pow(v, w, m) + pow(w, v, m)) % m})
		}
	}
	sort.Slice(b, func(i, j int) bool { return b[i].wt > b[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	sum := 0
	for _, e := range b {
		v, w, wt := e.v, e.w, e.wt
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
			sum += wt
		}
	}
	Fprint(out, sum)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
