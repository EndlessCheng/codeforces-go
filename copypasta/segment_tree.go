package copypasta

import "math/bits"

// https://oi-wiki.org/ds/seg/
// 题目推荐 https://cp-algorithms.com/data_structures/segment_tree.html#toc-tgt-12
// TODO: zkw 线段树 https://codeforces.com/blog/entry/18051

// TIPS: 对于复杂的区间操作（如区间开方），可以判断区间元素是否相同，然后 lazy
// TIPS: 一般来说会有一个核心函数，如 min/max/+/gcd/...
// CF961E: 用归并树查询区间内大于等于某个数的元素个数（rank）

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异。
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中。
type stNode struct {
	l, r   int
	val    int64
	maxPos int
}
type segmentTree []stNode // t := make(segmentTree, 4*n)

func (segmentTree) min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func (segmentTree) max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func (segmentTree) gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (t segmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	// *custom* 核心函数
	t[o].val = t.max(lo.val, ro.val)
	//if ro.val >= lo.val { // maxPos 取最右侧；若为 > 符号则取最左侧
	//	t[o].val, t[o].maxPos = ro.val, ro.maxPos
	//} else {
	//	t[o].val, t[o].maxPos = lo.val, lo.maxPos
	//}
}

func (t segmentTree) _build(arr []int64, o, l, r int) {
	t[o].l, t[o].r = l, r // 注意：一定要初始化 l 和 r
	if l == r {
		//t[o].val = arr[l] // if arr start at 1
		t[o].val = arr[l-1]
		//t[o].maxPos = l - 1
		return
	}
	m := (l + r) >> 1
	t._build(arr, o<<1, l, m)
	t._build(arr, o<<1|1, m+1, r)
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
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	vl := t._query(o<<1, l, r)
	vr := t._query(o<<1|1, l, r)
	// *custom* 核心函数
	return t.max(vl, vr)
}

// if arr start at 1, end at n
//func (t segmentTree) init(arr []int64)          { t._build(arr, 1, 1, len(arr)-1) }
func (t segmentTree) init(arr []int64)          { t._build(arr, 1, 1, len(arr)) }
func (t segmentTree) update(idx int, val int64) { t._update(1, idx, val) }   // 1<=idx<=n
func (t segmentTree) query(l, r int) int64      { return t._query(1, l, r) } // [l,r] 1<=l<=r<=n

// others
func (t segmentTree) _query2(o, l, r int) (res int64, maxPos int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val, t[o].maxPos
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query2(o<<1, l, r)
	}
	if l > m {
		return t._query2(o<<1|1, l, r)
	}
	vl, pl := t._query2(o<<1, l, r)
	vr, pr := t._query2(o<<1|1, l, r)
	if vl <= vr { // 取等号时，返回的是 pl，即最左侧的位置；若写成 <，则会在取等号时返回 pr，即最右侧的位置
		return vl, pl
	}
	return vr, pr
}
func (t segmentTree) query2(l, r int) (int64, int) { return t._query2(1, l, r) } // [l,r] 1<=l<=r<=n
//func (t segmentTree) queryAll() int64              { return t._query(1, 1, n) }

//

// 模板题 https://www.luogu.com.cn/problem/P3373
type lazySTNode struct {
	l, r        int
	sum         int64
	addChildren int64 // 子节点待更新
}
type lazySegmentTree []lazySTNode // t := make(lazySegmentTree, 4*n)

func (t lazySegmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = lo.sum + ro.sum // mod
}

func (t lazySegmentTree) _build(arr []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		//t[o].sum = arr[l] // if arr start at 1
		t[o].sum = arr[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(arr, o<<1, l, m)
	t._build(arr, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t lazySegmentTree) _spread(o int) {
	if add := t[o].addChildren; add != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum += add * int64(lo.r-lo.l+1)
		ro.sum += add * int64(ro.r-ro.l+1)
		lo.addChildren += add
		ro.addChildren += add
		//lo.sum = (lo.sum + add*int64(lo.r-lo.l+1)) % mod
		//ro.sum = (ro.sum + add*int64(ro.r-ro.l+1)) % mod
		//lo.addChildren = (lo.addChildren + add) % mod
		//ro.addChildren = (ro.addChildren + add) % mod
		t[o].addChildren = 0
	}
}

func (t lazySegmentTree) _update(o, l, r int, add int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum += add * int64(or-ol+1)
		t[o].addChildren += add
		//t[o].sum = (t[o].sum + add*int64(or-ol+1)) % mod
		//t[o].addChildren = (t[o].addChildren + add) % mod
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, add)
	}
	if m < r {
		t._update(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t lazySegmentTree) _query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t._spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		res += t._query(o<<1, l, r)
	}
	if m < r {
		res += t._query(o<<1|1, l, r)
	}
	//res %= mod
	return
}

// if arr start at 1, end at n
//func (t lazySegmentTree) init(arr []int64)           { t._build(arr, 1, 1, len(arr)-1) }
func (t lazySegmentTree) init(arr []int64)           { t._build(arr, 1, 1, len(arr)) }
func (t lazySegmentTree) update(l, r int, val int64) { t._update(1, l, r, val) }  // [l,r] 1<=l<=r<=n
func (t lazySegmentTree) query(l, r int) int64       { return t._query(1, l, r) } // [l,r] 1<=l<=r<=n

//

// 可持久化线段树（又称函数式线段树、主席树） Persistent Segment Tree
// https://oi-wiki.org/ds/persistent-seg/
// 模板题 https://www.luogu.com.cn/problem/P3834
// TODO: 补题 https://codeforces.ml/contest/786/problem/C
// TODO: 补题 https://codeforces.ml/contest/840/problem/D
type pstNode struct {
	l, r   int
	lo, ro *pstNode // 由于使用了指针，pstNode 必须存放于一个分配了足够空间的 slice 中，避免扩容时改变了内存位置
	sum    int64
}
type pst struct {
	nodes        []pstNode
	versionRoots []*pstNode
}

// 区间长度，版本数，最大更新次数
func newPST(n, versions, maxUpdateTimes int) *pst {
	// https://oi-wiki.org/ds/persistent-seg/
	maxNodeSize := 2*n + (bits.Len(uint(n))+1)*maxUpdateTimes
	return &pst{
		make([]pstNode, 0, maxNodeSize),
		make([]*pstNode, versions+1),
	}
}

func (*pst) _pushUp(o *pstNode) {
	o.sum = o.lo.sum + o.ro.sum
}

func (t *pst) _build(l, r int) *pstNode {
	t.nodes = append(t.nodes, pstNode{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = t._build(l, m)
	o.ro = t._build(m+1, r)
	//t._pushUp(o)
	return o
}

func (t *pst) _buildArr(arr []int64, l, r int) *pstNode {
	t.nodes = append(t.nodes, pstNode{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		o.sum = arr[l]
		return o
	}
	m := (l + r) >> 1
	o.lo = t._buildArr(arr, l, m)
	o.ro = t._buildArr(arr, m+1, r)
	t._pushUp(o)
	return o
}

func (t *pst) _update(o *pstNode, idx int, val int64) *pstNode {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.sum += val
		//o.sum = val
		return o
	}
	if m := o.lo.r; idx <= m {
		o.lo = t._update(o.lo, idx, val)
	} else {
		o.ro = t._update(o.ro, idx, val)
	}
	t._pushUp(o)
	return o
}

func (t *pst) _query(o *pstNode, l, r int) (res int64) {
	if l <= o.l && o.r <= r {
		return o.sum
	}
	m := o.lo.r
	if l <= m {
		res += t._query(o.lo, l, r)
	}
	if m < r {
		res += t._query(o.ro, l, r)
	}
	return
}

func (t *pst) _queryKth(o1, o2 *pstNode, k int) (allKth int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := int(o2.lo.sum - o1.lo.sum); d >= k {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-d)
	}
}

// 初始化，创建版本为 0 的线段树
func (t *pst) init(n int) {
	t.versionRoots[0] = t._build(1, n)
}
func (t *pst) initArr(arr []int64) { // arr start at 1
	t.versionRoots[0] = t._buildArr(arr, 1, len(arr)-1)
}

func (t *pst) copy(dstVersion, srcVersion int) {
	t.versionRoots[dstVersion] = t.versionRoots[srcVersion]
}

// 基于版本为 srcVersion 的线段树，更新其 idx 位置上的值 += val（1<=idx<=n）
// 用更新后的结果覆盖 dstVersion 版本
// dstVersion 和 srcVersion 可以相同
// 若求区间第 k 大，遍历 kthArr 作为 idx 传入
func (t *pst) update(dstVersion, srcVersion int, idx int, val int64) {
	t.versionRoots[dstVersion] = t._update(t.versionRoots[srcVersion], idx, val)
}

// 查询第 version 个版本下的区间和
// [l,r] 1<=l<=r<=n
func (t *pst) query(version int, l, r int) (sum int64) {
	return t._query(t.versionRoots[version], l, r)
}

// 查询区间第 k 大/小在整个数组上的名次 1<=allKth<=n，即排序后的数组下标 (+1)
// [l,r] 1<=l<=r<=n
func (t *pst) queryKth(l, r int, k int) (allKth int) {
	return t._queryKth(t.versionRoots[l-1], t.versionRoots[r], k)
	// 	sortedArr := make([]int, n)
	//	copy(sortedArr, a)
	//	sort.Ints(sortedArr)
}
