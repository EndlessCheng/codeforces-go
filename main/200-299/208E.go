package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF208E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, t, q, k int
	Fscan(in, &n)
	g := make([][]int, n)
	roots := []int{}
	for w := 0; w < n; w++ {
		Fscan(in, &v)
		if v > 0 {
			g[v-1] = append(g[v-1], w)
		} else {
			roots = append(roots, w)
		}
	}

	const mx = 17
	pa := make([][mx]int, n)
	type info struct{ in, out, d int }
	is := make([]info, n)
	depT := make([][]int, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		pa[v][0] = p
		t++
		is[v].in = t
		is[v].d = d
		depT[d] = append(depT[d], t)
		for _, w := range g[v] {
			if w != p {
				f(w, v, d+1)
			}
		}
		is[v].out = t
	}
	for _, root := range roots {
		f(root, -1, 0)
	}
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}

	uptoKthPa := func(v, k int) int {
		for i := 0; i < mx && v != -1; i++ {
			if k>>i&1 == 1 {
				v = pa[v][i]
			}
		}
		return v
	}
	query := func(v, d int) int {
		i := is[v]
		d += i.d
		l := sort.SearchInts(depT[d], i.in)
		r := sort.SearchInts(depT[d], i.out+1)
		return r - l
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &k)
		if v = uptoKthPa(v-1, k); v == -1 {
			Fprint(out, "0 ")
		} else {
			Fprint(out, query(v, k)-1, " ")
		}
	}
}

//func main() { CF208E(os.Stdin, os.Stdout) }
