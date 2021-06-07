package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
	"sort"
)

// todo 目前最快做法，待研究 https://codeforces.com/contest/484/submission/100555356

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type data84 struct {
	mx, pre, suf int
	full         bool
}
type node84 struct {
	lo, ro *node84
	l, r   int
	data84
}

func max84(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func op84(a, b data84) (c data84) {
	c.pre = a.pre
	if a.full {
		c.pre += b.pre
	}
	c.suf = b.suf
	if b.full {
		c.suf += a.suf
	}
	c.mx = max84(max84(max84(a.mx, b.mx), max84(c.pre, c.suf)), a.suf+b.pre)
	c.full = a.full && b.full
	return
}

func (o *node84) maintain() {
	o.data84 = op84(o.lo.data84, o.ro.data84)
}

func build84(l, r int) *node84 {
	o := &node84{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build84(l, m)
	o.ro = build84(m+1, r)
	return o
}

func (o node84) insert(i int) *node84 {
	if o.l == o.r {
		o.mx, o.pre, o.suf, o.full = 1, 1, 1, true
		return &o
	}
	if m := o.lo.r; i <= m {
		o.lo = o.lo.insert(i)
	} else {
		o.ro = o.ro.insert(i)
	}
	o.maintain()
	return &o
}

func (o *node84) query(l, r int) data84 {
	if l <= o.l && o.r <= r {
		return o.data84
	}
	m := o.lo.r
	if r <= m {
		return o.lo.query(l, r)
	}
	if m < l {
		return o.ro.query(l, r)
	}
	return op84(o.lo.query(l, r), o.ro.query(l, r))
}

func CF484E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, w int
	Fscan(in, &n)
	type pair struct{ h, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].h)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].h > a[j].h })

	t := make([]*node84, n+1)
	t[0] = build84(1, n)
	for i, p := range a {
		t[i+1] = t[i].insert(p.i + 1)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r, &w)
		i := sort.Search(n-1, func(i int) bool { return t[i+1].query(l, r).mx >= w })
		Fprintln(out, a[i].h)
	}
}

//func main() { CF484E(os.Stdin, os.Stdout) }
