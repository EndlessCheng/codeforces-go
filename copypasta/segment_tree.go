package copypasta

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//

type node struct {
	l, r int
	val  int
}
type segmentTree []node // t := make(segmentTree, 4*n)

func (t segmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	// *custom*
	t[o].val = max(lo.val, ro.val)
}

func (t segmentTree) _build(arr []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = arr[l]
		// *custom*
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)

	// *custom* after built children

	t._pushUp(o)
}

func (t segmentTree) _update(o, idx int, val int) {
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
	mid := (t[o].l + t[o].r) >> 1
	// *custom*
	res = -1e9
	if l <= mid {
		res = max(res, t._query(o<<1, l, r))
	}
	if mid < r {
		res = max(res, t._query(o<<1|1, l, r))
	}
	return
}

// NOTE: arr must start at 1
func (t segmentTree) init(arr []int)          { t._build(arr, 1, 1, len(arr)-1) }
func (t segmentTree) update(idx int, val int) { t._update(1, idx, val) }
func (t segmentTree) query(l, r int)          { t._query(1, l, r) } // [l,r]

//

type node2 struct {
	l, r int
	val  int
	// TODO
}
type lazySegmentTree []node2 // t := make(lazySegmentTree, 4*n)
