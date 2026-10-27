package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type seg []struct {
	l, r      int
	size, sum int
	flip      bool
}

func (t seg) apply(o int) {
	cur := &t[o]
	cur.sum = cur.size - cur.sum
	cur.flip = !cur.flip
}

func (t seg) spread(o int) {
	f := t[o].flip
	if !f {
		return
	}
	t.apply(o << 1)
	t.apply(o<<1 | 1)
	t[o].flip = false
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].size = a[l+1] - a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o].size = t[o<<1].size + t[o<<1|1].size
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t[o].sum = t[o<<1].sum + t[o<<1|1].sum
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var Q int
	Fscan(in, &Q)
	qs := make([]struct{ op, l, r int }, Q)
	a := make([]int, 0, Q*2)
	for i := range qs {
		Fscan(in, &qs[i].op, &qs[i].l, &qs[i].r)
		a = append(a, qs[i].l, qs[i].r+1)
	}

	slices.Sort(a)
	a = slices.Compact(a)
	m := len(a)

	t := make(seg, 2<<bits.Len(uint(m-2)))
	t.build(a, 1, 0, m-2)

	for _, q := range qs {
		l := sort.SearchInts(a, q.l)
		r := sort.SearchInts(a, q.r+1) - 1
		if q.op == 1 {
			t.update(1, l, r)
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
