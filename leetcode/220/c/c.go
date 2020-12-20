package main

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r int
	val  int
}

func (t seg) set(o int, val int) {
	t[o].val = val
}

func (t seg) do(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.do(lo.val, ro.val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = -1e18
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i int, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t seg) query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.do(vl, vr)
}

func (t seg) queryAll() int { return t[1].val }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

func maxResult(a []int, k int) (ans int) {
	n := len(a)
	if n == 1 {
		return a[0]
	}
	t := newSegmentTree(make([]int, n))
	t.update(1, 1, a[0])
	for i := 1; ; i++ {
		cur := t.query(1, max(1, i-k+1), i+1) + a[i]
		t.update(1, i+1, cur)
		if i == n-1 {
			return cur
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
