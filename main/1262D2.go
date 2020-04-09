package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

type node1262 struct {
	l, r   int
	lo, ro *node1262
	sum    int
}
type pst1262 []*node1262

func (t pst1262) _build(l, r int) *node1262 {
	o := &node1262{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = t._build(l, m)
	o.ro = t._build(m+1, r)
	return o
}

func (t pst1262) _update(o *node1262, idx int) *node1262 {
	tmp := *o
	o = &tmp
	if o.l == o.r {
		o.sum++
		return o
	}
	if idx <= o.lo.r {
		o.lo = t._update(o.lo, idx)
	} else {
		o.ro = t._update(o.ro, idx)
	}
	o.sum = o.lo.sum + o.ro.sum
	return o
}

func (t pst1262) _queryKth(o1, o2 *node1262, k int) (idx int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := o2.lo.sum - o1.lo.sum; d >= k {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-d)
	}
}

func (t pst1262) init(n int)                     { t[0] = t._build(1, n) }
func (t pst1262) update(ver, idx int)            { t[ver+1] = t._update(t[ver], idx) }
func (t pst1262) queryKth(l, r, k int) (idx int) { return t._queryKth(t[l-1], t[r], k) }

// github.com/EndlessCheng/codeforces-go
func CF1262D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ v, i int }

	var n, m, k, p int
	Fscan(in, &n)
	a := make([]int, n)
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &a[i])
		ps[i] = pair{a[i], i}
	}

	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.v > b.v || a.v == b.v && a.i < b.i })
	t := make(pst1262, n+1)
	t.init(n)
	for i, p := range ps {
		t.update(i, p.i+1)
	}
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &k, &p)
		Fprintln(out, a[t.queryKth(1, k, p)-1])
	}
}

//func main() { CF1262D2(os.Stdin, os.Stdout) }
