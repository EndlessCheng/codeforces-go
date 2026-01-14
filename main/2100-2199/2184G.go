package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg84 []struct{ l, r, mn int }

func (t seg84) maintain(o int) {
	t[o].mn = min(t[o<<1].mn, t[o<<1|1].mn)
}

func (t seg84) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mn = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg84) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.mn = v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}

var gMin84 int

func (t seg84) query(o, ql, qr int) int {
	tl, tr := t[o].l, t[o].r
	if ql <= tl && tr <= qr {
		mn := min(gMin84, t[o].mn)
		if mn == tr-ql {
			return 1
		}
		if mn > tr-ql {
			gMin84 = mn
			return -1
		}
		if tl == tr {
			return 0
		}
		res := t.query(o<<1, ql, qr)
		if res < 0 {
			res = t.query(o<<1|1, ql, qr)
		}
		return max(res, 0)
	}

	m := (tl + tr) >> 1
	if qr <= m {
		return t.query(o<<1, ql, qr)
	}
	if m < ql {
		return t.query(o<<1|1, ql, qr)
	}

	res := t.query(o<<1, ql, qr)
	if res < 0 {
		res = t.query(o<<1|1, ql, qr)
	}
	return res
}

func cf2184G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, op, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		t := make(seg84, 2<<bits.Len(uint(n-1)))
		t.build(a, 1, 0, n-1)
		for range q {
			Fscan(in, &op, &l, &r)
			if op == 1 {
				t.update(1, l-1, r)
			} else {
				gMin84 = 1e9
				Fprintln(out, max(t.query(1, l-1, r-1), 0))
			}
		}
	}
}

//func main() { cf2184G(bufio.NewReader(os.Stdin), os.Stdout) }
