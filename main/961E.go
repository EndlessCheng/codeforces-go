package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

type stNode961 struct {
	l, r int
	vals []int
}
type segmentTree961 []stNode961

func (t segmentTree961) _pushUp(o int) []int {
	a, b := t[o<<1].vals, t[o<<1|1].vals
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, n+m)
	for {
		if i == n {
			return append(res, b[j:]...)
		}
		if j == m {
			return append(res, a[i:]...)
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
}

func (t segmentTree961) _build(arr []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].vals = []int{arr[l-1]}
		return
	}
	m := (l + r) >> 1
	t._build(arr, o<<1, l, m)
	t._build(arr, o<<1|1, m+1, r)
	t[o].vals = t._pushUp(o)
}

func (t segmentTree961) _query(o, l, r, x int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		a := t[o].vals
		return len(a) - sort.SearchInts(a, x)
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		res += t._query(o<<1, l, r, x)
	}
	if r > m {
		res += t._query(o<<1|1, l, r, x)
	}
	return
}

func (t segmentTree961) init(arr []int)        { t._build(arr, 1, 1, len(arr)) }
func (t segmentTree961) query(l, r, x int) int { return t._query(1, l, r, x) }

// github.com/EndlessCheng/codeforces-go
func CF961E(_r io.Reader, _w io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	t := make(segmentTree961, 4*n)
	t.init(a)
	ans := int64(0)
	for i, v := range a[:n-1] {
		if v > i+1 {
			ans += int64(t.query(i+2, min(v, n), i+1))
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF961E(os.Stdin, os.Stdout)
//}
