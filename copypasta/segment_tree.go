package copypasta

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异。
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中。
type stNode struct {
	l, r   int
	val    int64
	maxPos int
}
type segmentTree []stNode // t := make(segmentTree, 4*n)

func (t segmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	// *custom*
	//t[o].val = max(lo.val, ro.val)
	if ro.val >= lo.val { // maxPos 取最右侧；若为 > 符号则取最左侧
		t[o].val, t[o].maxPos = ro.val, ro.maxPos
	} else {
		t[o].val, t[o].maxPos = lo.val, lo.maxPos
	}
}

func (t segmentTree) _build(arr []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		//t[o].val = arr[l] // if arr start at 1
		t[o].val = arr[l-1]
		t[o].maxPos = l - 1
		// *custom*
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)

	// *custom* after built children

	t._pushUp(o)
}

func (t segmentTree) _update(o, idx int, val int64) {
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

func (t segmentTree) _query(o, l, r int) (res int64) {
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

func (t segmentTree) _query2(o, l, r int) (res int64, maxPos int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val, t[o].maxPos
	}
	mid := (t[o].l + t[o].r) >> 1
	// *custom*
	res = -1e9
	if l <= mid {
		res, maxPos = t._query2(o<<1, l, r)
	}
	if mid < r {
		if newRes, newMaxPos := t._query2(o<<1|1, l, r); newRes >= res { // maxPos 取最右侧；若为 > 符号则取最左侧
			res, maxPos = newRes, newMaxPos
		}
	}
	return
}

// if arr start at 1, end at n
//func (t segmentTree) init(arr []int64)          { t._build(arr, 1, 1, len(arr)-1) }
func (t segmentTree) init(arr []int64)             { t._build(arr, 1, 1, len(arr)) }
func (t segmentTree) update(idx int, val int64)    { t._update(1, idx, val) }    // 1<=idx<=n
func (t segmentTree) query(l, r int) int64         { return t._query(1, l, r) }  // [l,r] 1<=l<=r<=n
func (t segmentTree) query2(l, r int) (int64, int) { return t._query2(1, l, r) } // [l,r] 1<=l<=r<=n

//

type lazySTNode struct {
	l, r        int
	sum         int64
	addChildren int64 // 子节点待更新
}
type lazySegmentTree []lazySTNode // t := make(lazySegmentTree, 4*n)

func (t lazySegmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	// *custom*
	t[o].sum = lo.sum + ro.sum
}

func (t lazySegmentTree) _build(arr []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		//t[o].sum = arr[l] // if arr start at 1
		t[o].sum = arr[l-1]
		// *custom*
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)

	// *custom* after built children

	t._pushUp(o)
}

func (t lazySegmentTree) _spread(o int) {
	if add := t[o].addChildren; add != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum += add * int64(lo.r-lo.l+1)
		ro.sum += add * int64(ro.r-ro.l+1)
		lo.addChildren += add
		ro.addChildren += add
		t[o].addChildren = 0
	}
}

func (t lazySegmentTree) _update(o, l, r int, add int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum += add * int64(or-ol+1)
		t[o].addChildren += add
		return
	}
	t._spread(o)
	mid := (ol + or) >> 1
	if l <= mid {
		t._update(o<<1, l, r, add)
	}
	if mid < r {
		t._update(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t lazySegmentTree) _query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t._spread(o)
	mid := (t[o].l + t[o].r) >> 1
	// *custom*
	if l <= mid {
		res += t._query(o<<1, l, r)
	}
	if mid < r {
		res += t._query(o<<1|1, l, r)
	}
	return
}

// if arr start at 1, end at n
//func (t lazySegmentTree) init(arr []int64)           { t._build(arr, 1, 1, len(arr)-1) }
func (t lazySegmentTree) init(arr []int64)           { t._build(arr, 1, 1, len(arr)) }
func (t lazySegmentTree) update(l, r int, val int64) { t._update(1, l, r, val) }  // [l,r] 1<=l<=r<=n
func (t lazySegmentTree) query(l, r int) int64       { return t._query(1, l, r) } // [l,r] 1<=l<=r<=n
