package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type seg []struct{ l, r, mi, mx int }

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 4*n)
	t.init(a)
	return t
}

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].mi = min(lo.mi, ro.mi)
	t[o].mx = max(lo.mx, ro.mx)
}

func (t seg) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mi = a[l-1]
		t[o].mx = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) (int, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mi, t[o].mx
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	mi1, mx1 := t._query(o<<1, l, r)
	mi2, mx2 := t._query(o<<1|1, l, r)
	return min(mi1, mi2), max(mx1, mx2)
}

func (t seg) init(a []int)              { t._build(a, 1, 1, len(a)) }
func (t seg) query(l, r int) (int, int) { return t._query(1, l, r) }

func longestSubarray(a []int, limit int) (ans int) {
	t := newSegmentTree(a)
	n := len(a)
	for l, r := 1, 1; r <= n; {
		for ; r <= n; r++ {
			mi, mx := t.query(l, r)
			if mx-mi > limit {
				break
			}
			ans = max(ans, r-l+1)
		}
		for ; l <= r; l++ {
			mi, mx := t.query(l, r)
			if mx-mi <= limit {
				break
			}
		}
	}
	return
}
