package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"math/bits"
	"os"
	"reflect"
	"sort"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var s []byte
	var m int
	Fscan(in, &s, &m)
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, ri := range rank {
		if h > 0 {
			h--
		}
		if ri > 0 {
			for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[ri] = h
	}
	st := make([][19]int, n)
	for i, v := range height {
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
	lcp := func(i, j int) int {
		if i == j {
			return n - i
		}
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		return _q(ri+1, rj+1)
	}

	type pair struct{ l, r int }
	ps := make([]pair, m)
	for i := range ps {
		Fscan(in, &ps[i].l, &ps[i].r)
		ps[i].l--
	}
	sort.Slice(ps, func(i, j int) bool {
		a, b := ps[i], ps[j]
		la, lb := a.r-a.l, b.r-b.l
		l := lcp(a.l, b.l)
		if l >= la || l >= lb {
			return la < lb || la == lb && (a.l < b.l || a.l == b.l && a.r < b.r)
		}
		return rank[a.l] < rank[b.l]
	})
	for _, p := range ps {
		Fprintln(out, p.l+1, p.r)
	}
}

func main() { run(os.Stdin, os.Stdout) }
