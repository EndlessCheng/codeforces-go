package main

import "sort"

type node struct{ l, r, val int }
type segmentTree []node

func (segmentTree) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t segmentTree) _build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
}

func (t segmentTree) _pushUp(o int) { t[o].val = t.max(t[o<<1].val, t[o<<1|1].val) }

func (t segmentTree) _update(o, idx, val int) {
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

func (t segmentTree) _query(o, l, r int) (res int) {
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
	return t.max(vl, vr)
}

func (t segmentTree) init(n int)          { t._build(1, 1, n) }
func (t segmentTree) update(idx, val int) { t._update(1, idx+1, val) }
func (t segmentTree) query(l, r int) int  { return t._query(1, l+1, r+1) }

func maxJumps(a []int, d int) (ans int) {
	mins := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val < ans {
				ans = val
			}
		}
		return ans
	}
	maxs := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val > ans {
				ans = val
			}
		}
		return ans
	}

	n := len(a)
	type pair struct{ v, i int }
	posL := make([]int, n)
	stack := []pair{{1e9, -1}}
	for i, v := range a {
		for {
			if top := stack[len(stack)-1]; top.v >= v {
				posL[i] = top.i + 1
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}
	posR := make([]int, n)
	stack = []pair{{1e9, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stack[len(stack)-1]; top.v >= v {
				posR[i] = top.i - 1
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	ps := make([]pair, n)
	for i := range ps {
		ps[i] = pair{a[i], i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })

	t := make(segmentTree, 4*n)
	t.init(n)
	for _, p := range ps {
		i := p.i
		l := maxs(0, i-d, posL[i])
		r := mins(n-1, i+d, posR[i])
		rangeMax := t.query(l, r)
		t.update(i, rangeMax+1)
	}
	return t.query(0, n-1)
}
