package main

type seg []struct{ l, r, val int }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = max(lo.val, ro.val)
}

func (t seg) _build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	vl := t._query(o<<1, l, r)
	vr := t._query(o<<1|1, l, r)
	return max(vl, vr)
}

func (t seg) init(n int)          { t._build(1, 1, n) }
func (t seg) update(idx, val int) { t._update(1, idx, val) }
func (t seg) query(l, r int) int  { return t._query(1, l, r) }

func constrainedSubsetSum(a []int, k int) (ans int) {
	ans = a[0]
	n := len(a)
	t := make(seg, 4*n)
	t.init(n)
	t.update(1, a[0])
	l, r := 1, 1
	for i := 1; i < n; i++ {
		v := max(a[i], t.query(l, r)+a[i])
		ans = max(ans, v)
		t.update(i+1, v)
		if r < n {
			r++
		}
		if r-l+1 > k {
			l++
		}
	}
	return
}
