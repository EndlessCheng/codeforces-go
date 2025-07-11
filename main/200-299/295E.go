package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	. "slices"
	"sort"
)

// https://github.com/EndlessCheng
type data95 struct{ c, s, res int }
type seg95 []struct {
	l, r int
	d    data95
}

func (t seg95) merge(l, r data95) data95 {
	return data95{l.c + r.c, l.s + r.s, l.res + r.res + r.s*l.c - l.s*r.c}
}

func (t seg95) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg95) update(o, i, c, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.d.c += c
		cur.d.s += v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, c, v)
	} else {
		t.update(o<<1|1, i, c, v)
	}
	cur.d = t.merge(t[o<<1].d, t[o<<1|1].d)
}

func (t seg95) query(o, l, r int) data95 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].d
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf295E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := Clone(a)
	xs := Clone(a)
	Fscan(in, &k)
	qs := make([]struct{ op, l, r int }, k)
	for i := range qs {
		Fscan(in, &qs[i].op, &qs[i].l, &qs[i].r)
		if qs[i].op == 1 {
			qs[i].l--
			j := qs[i].l
			b[j] += qs[i].r
			xs = append(xs, b[j])
		}
	}
	Sort(xs)
	xs = Compact(xs)
	m := len(xs)

	t := make(seg95, 2<<bits.Len(uint(m-1)))
	t.build(1, 0, m-1)
	lb := sort.SearchInts
	for _, v := range a {
		t.update(1, lb(xs, v), 1, v)
	}
	for _, q := range qs {
		if q.op == 1 {
			i := q.l
			t.update(1, lb(xs, a[i]), -1, -a[i])
			a[i] += q.r
			t.update(1, lb(xs, a[i]), 1, a[i])
		} else {
			l, r := lb(xs, q.l), lb(xs, q.r+1)-1
			if l > r {
				Fprintln(out, 0)
			} else {
				Fprintln(out, t.query(1, l, r).res)
			}
		}
	}
}

//func main() { cf295E(bufio.NewReader(os.Stdin), os.Stdout) }
