package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1609D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, d, v, w, ex int
	Fscan(in, &n, &d)
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for ; d > 0; d-- {
		Fscan(in, &v, &w)
		v, w = find(v), find(w)
		if v == w {
			ex++
		} else {
			if sz[v] > sz[w] {
				v, w = w, v
			}
			sz[w] += sz[v]
			sz[v] = 0 // 注意这里
			fa[v] = w
		}
		a := append([]int(nil), sz...)
		sort.Ints(a) // or kth
		s := -1
		for i := n; i >= n-ex; i-- {
			s += a[i]
		}
		Fprintln(out, s)
	}
}

//func main() { CF1609D(os.Stdin, os.Stdout) }
