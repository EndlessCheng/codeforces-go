package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1076E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, q, d int
	var x int64
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type info struct{ tin, tout, d int }
	is := make([]info, n)
	depT := make([][]int, n)
	t := 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		t++
		is[v].tin = t
		is[v].d = d
		depT[d] = append(depT[d], t)
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
		is[v].tout = t
	}
	f(0, -1, 0)
	add := make([]int64, n)
	diff := make([][]int64, n)
	for i, row := range depT {
		diff[i] = make([]int64, len(row)+1)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &d, &x)
		v--
		add[v] += x
		info := is[v]
		if d += info.d + 1; d < n {
			l := sort.SearchInts(depT[d], info.tin)
			r := sort.SearchInts(depT[d], info.tout+1)
			diff[d][l] += x
			diff[d][r] -= x
		}
	}

	ans := make([]interface{}, n)
	var f2 func(v, fa, d int, s int64)
	f2 = func(v, fa, d int, s int64) {
		s += add[v] - diff[d][0]
		diff[d][1] += diff[d][0]
		diff[d] = diff[d][1:]
		ans[v] = s
		for _, w := range g[v] {
			if w != fa {
				f2(w, v, d+1, s)
			}
		}
	}
	f2(0, -1, 0, 0)
	Fprintln(out, ans...)
}

//func main() { CF1076E(os.Stdin, os.Stdout) }
