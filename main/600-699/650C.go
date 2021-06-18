package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF650C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	fa := make([]int, n*m+1)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	a := make([]int, n*m+1)
	id := make([]int, n*m+1)
	for i := 1; i <= n*m; i++ {
		Fscan(in, &a[i])
		fa[i] = i
		id[i] = i
	}
	sort.Slice(id[1:], func(i, j int) bool { return a[id[i+1]] < a[id[j+1]] })

	rank := make([]int, n*m+1)
	mxR := make([]int, n+1)
	mxC := make([]int, m+1)
	for _, i := range id[1:] {
		r, c := (i-1)/m+1, (i-1)%m+1
		fr, fc, fi := find(mxR[r]), find(mxC[c]), find(i)
		x := rank[fr]
		if a[fi] > a[fr] {
			x++
		} else {
			fa[fr] = fi
		}
		y := rank[fc]
		if a[fi] > a[fc] {
			y++
		} else {
			fa[fc] = fi
		}
		rank[fi] = max(x, y)
		mxR[r] = fi
		mxC[c] = fi
	}
	for i := 1; i <= n*m; i++ {
		Fprint(out, rank[find(i)], " ")
		if i%m == 0 {
			Fprintln(out)
		}
	}
}

//func main() { CF650C(os.Stdin, os.Stdout) }
