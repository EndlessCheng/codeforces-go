package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick67 []int

func newFenwickTree67(n int) fenwick67 {
	t := make(fenwick67, n+1)
	for i := range t {
		t[i] = -1e18
	}
	return t
}

func (f fenwick67) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], v)
	}
}

func (f fenwick67) pre(i int) int {
	res := int(-1e18)
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

func cf1667B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}

		sorted := slices.Clone(s)
		slices.Sort(sorted)
		sorted = slices.Compact(sorted)
		m := len(sorted)

		lessT := newFenwickTree67(m)
		grT := newFenwickTree67(m)
		maxF := newFenwickTree67(m)

		f := 0
		for r, v := range s {
			v = sort.SearchInts(sorted, v)
			if r > 0 {
				f = max(lessT.pre(v)+r, grT.pre(m-1-v)-r, maxF[v])
			}
			lessT.update(v+1, f-r)
			grT.update(m-v, f+r)
			maxF[v] = max(maxF[v], f)
		}
		Fprintln(out, f)
	}
}

//func main() { cf1667B(bufio.NewReader(os.Stdin), os.Stdout) }
