package copypasta

// 本质上来说，每个点对应着 O(logn) 个线段树上的节点，每个区间可以拆分成至多 O(logn) 个线段树上的区间

// https://oi-wiki.org/ds/seg/
// https://cp-algorithms.com/data_structures/segment_tree.html
// https://codeforces.com/blog/entry/18051
// https://codeforces.com/blog/entry/15890
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/SegmentTree.java.html
// https://leetcode.com/articles/a-recursive-approach-to-segment-trees-range-sum-queries-lazy-propagation/
// 总结得比较详细 https://www.acwing.com/blog/content/1684/

// EXTRA: 离散化线段树
// https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/F https://www.luogu.com.cn/problem/P5848

// EXTRA: zkw 线段树
// 虽说写起来很快，但是 debug 就比较难受了
// 张昆玮《统计的力量》
// https://zhuanlan.zhihu.com/p/29876526
// https://zhuanlan.zhihu.com/p/29937723
// https://codeforces.com/blog/entry/18051 Efficient and easy segment trees

// todo 李超线段树
//  https://zhuanlan.zhihu.com/p/64946571
//  https://www.luogu.com.cn/blog/fzber0103/Li-Chao-Tree

// todo 经典题 ai>aj>ak 三元组的数量 https://codeforces.com/problemset/problem/61/E
// 根号线段树思想 https://codeforces.com/problemset/problem/920/F
// GCD 好题 https://codeforces.com/problemset/problem/914/D
// todo 整理 CF961E: 用归并树查询区间内大于等于某个数的元素个数（rank）     其他方法？
// 题目推荐 https://cp-algorithms.com/data_structures/segment_tree.html#toc-tgt-12
// LC https://leetcode-cn.com/tag/segment-tree/
// 另见 dp.go 的动态 DP 部分

// 某个 do
func (seg) maxPos(a, b int64, pa, pb int) (int64, int) {
	if b >= a { // >= 为相同元素最右侧位置；若为 > 符号则是相同元素最左侧位置
		return b, pb
	}
	return a, pa
}

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中
type seg []struct {
	l, r int
	val  int64
}

func (t seg) set(o int, val int64) {
	t[o].val = val
}

func (t seg) do(a, b int64) int64 {
	// + * | & ^ min max gcd maxPos mulMatrix ...
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.do(lo.val, ro.val)
}

func (t seg) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i int, val int64) {
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
func (t seg) query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	//defer t.maintain(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.do(vl, vr)
}

func (t seg) queryAll() int64 { return t[1].val }

// a starts at 0
func newSegmentTree(a []int64) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

//

// 单个更新操作：
// + min/max https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/A https://codeforces.com/problemset/problem/1321/E
// + Σ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/D https://www.luogu.com.cn/problem/P3372
// | & https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/C
// = min https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/E
// = Σ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/F
// https://codeforces.com/problemset/problem/1114/F
//
// 多个更新操作复合：
// * + Σ https://www.luogu.com.cn/problem/P3373
// = + Σ https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/A
//
// EXTRA: 多项式更新 Competitive Programmer’s Handbook Ch.28
// 比如区间加等差数列（差分法）https://www.luogu.com.cn/problem/P1438 https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/B
type lazyST []struct {
	l, r int
	todo int64
	sum  int64
}

func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = lo.sum + ro.sum // % mod
}

func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazyST) do(o int, v int64) {
	to := &t[o]
	to.todo += v                     // % mod
	to.sum += int64(to.r-to.l+1) * v // % mod
}

func (t lazyST) spread(o int) {
	if add := t[o].todo; add != 0 {
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
	}
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	//defer t.maintain(o)
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return vl + vr // % mod
}

func (t lazyST) queryAll() int64 { return t[1].sum }

// a starts at 0
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

//

// todo 权值线段树
// 浅谈权值线段树到主席树 https://www.luogu.com.cn/blog/your-alpha1022/WeightSegmentTree-ChairmanTree

//

// EXTRA: 线段树合并
// todo https://www.luogu.com.cn/blog/styx-ferryman/xian-duan-shu-ge-bing-zong-ru-men-dao-fang-qi
// todo https://www.luogu.com.cn/problem/P4556

//

// 可持久化线段树（又称函数式线段树、主席树） Persistent Segment Tree
// NOTE: 对于 CF 上的 2e5 数据，直接 new(pstNode) 比在数组上创建指针要慢 700ms
//       因此建议对时限比较紧的题目改成在数组上创建指针，这里为了简化代码使用的是 new(pstNode)
// todo 区间最值操作与区间历史最值详解 https://www.luogu.com.cn/blog/Hakurei-Reimu/seg-beats
// https://oi-wiki.org/ds/persistent-seg/
// 模板题 https://www.luogu.com.cn/problem/P3834
// todo 区间元素种类数 https://codeforces.com/problemset/problem/620/E
// todo 整理模板+重写 https://codeforces.com/problemset/problem/1262/D2
//                  https://codeforces.com/problemset/problem/813/E
// TODO 查询出现次数大于区间长度一半的元素(强制在线) LC1157 https://leetcode-cn.com/problems/online-majority-element-in-subarray/
// TODO: 补题 https://codeforces.com/contest/786/problem/C
//  https://codeforces.com/problemset/problem/840/D
//  https://codeforces.com/problemset/problem/893/F
//  https://codeforces.com/problemset/problem/837/G
// todo 标记永久化 一本通 p239
type pstNode struct {
	l, r   int
	lo, ro *pstNode
	sum    int64
}
type pst []*pstNode // t := make(pst, versions+1) // 一般 versions 为 n

func (pst) _pushUp(o *pstNode) {
	o.sum = o.lo.sum + o.ro.sum
}

func (t pst) _build(l, r int) *pstNode {
	o := &pstNode{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = t._build(l, m)
	o.ro = t._build(m+1, r)
	//t._pushUp(o)
	return o
}

func (t pst) _buildArr(arr []int64, l, r int) *pstNode {
	o := &pstNode{l: l, r: r}
	if l == r {
		// arr starts at 1
		o.sum = arr[l]
		return o
	}
	m := (l + r) >> 1
	o.lo = t._buildArr(arr, l, m)
	o.ro = t._buildArr(arr, m+1, r)
	t._pushUp(o)
	return o
}

func (t pst) _update(o *pstNode, idx int, val int64) *pstNode {
	tmp := *o
	o = &tmp
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

func (t pst) _query(o *pstNode, l, r int) (res int64) {
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

func (t pst) _queryKth(o1, o2 *pstNode, k int) (allKth int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := o2.lo.sum - o1.lo.sum; d >= int64(k) {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-int(d))
	}
}

func (t pst) init(n int)              { t[0] = t._build(1, n) }                  // 创建版本为 0 的线段树
func (t pst) initArr(arr []int64)     { t[0] = t._buildArr(arr, 1, len(arr)-1) } // arr starts at 1
func (t pst) copy(dstVer, srcVer int) { t[dstVer] = t[srcVer] }

// 单点更新：基于版本为 srcVer 的线段树，用更新后的结果覆盖 dstVer 版本
// 1<=idx<=n   dstVer 可以和 srcVer 相同
// EXTRA: 求区间第 k 大时，遍历 kthArr 作为 idx 传入，见 segment_tree_test.go
func (t pst) update(dstVer, srcVer, idx int, val int64) { t[dstVer] = t._update(t[srcVer], idx, val) }
func (t pst) updateKth(ver, kth int)                    { t[ver+1] = t._update(t[ver], kth, 1) }

// 查询第 ver 个版本下的区间值
// [l,r] 1<=l<=r<=n
func (t pst) query(ver, l, r int) int64 { return t._query(t[ver], l, r) }

// EXTRA: 查询区间第 k 大/小在整个数组上的名次 1<=allKth<=n，即排序后的数组下标 (+1)
// [l,r] 1<=l<=r<=n
func (t pst) queryKth(l, r, kth int) (allKth int) { return t._queryKth(t[l-1], t[r], kth) }

//sortedArr := make([]int, n)
//copy(sortedArr, a)
//sort.Ints(sortedArr)

//

// Wavelet Trees
// todo https://www.geeksforgeeks.org/wavelet-trees-introduction/
// https://codeforces.com/blog/entry/52854
// https://www.youtube.com/watch?v=4aSv9PcecDw
