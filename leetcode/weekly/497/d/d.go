package main

import "math/bits"

// https://space.bilibili.com/206214
var targetGcd int

// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
// 线段树有两个下标，一个是线段树节点的下标，另一个是线段树维护的区间的下标
// 节点的下标：从 1 开始，如果你想改成从 0 开始，需要把左右儿子下标分别改成 node*2+1 和 node*2+2
// 区间的下标：从 0 开始
type seg []int

// 线段树维护数组 a
func newSegmentTreeWithArray(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

// 合并左右儿子的 val 到当前节点的 val
func (t seg) maintain(node int) {
	t[node] = gcd(t[node*2], t[node*2+1])
}

// 用 a 初始化线段树
// 时间复杂度 O(n)
func (t seg) build(a []int, node, l, r int) {
	if l == r { // 叶子
		if a[l]%targetGcd == 0 {
			t[node] = a[l] // 初始化叶节点的值
		}
		return
	}
	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

// 更新 a[i]
// 调用 t.update(1, 0, n-1, i, val)
// 0 <= i <= n-1
// 时间复杂度 O(log n)
func (t seg) update(node, l, r, i, val int) {
	if l == r { // 叶子（到达目标）
		if val%targetGcd == 0 {
			t[node] = val
		} else {
			t[node] = 0 // 0 和任何数 g 的 GCD 都是 g，所以设置为 0 不影响所有数的 GCD
		}
		return
	}
	m := (l + r) / 2
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i, val)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i, val)
	}
	t.maintain(node)
}

// 返回用 GCD 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
// 调用 t.query(1, 0, n-1, ql, qr)
// 0 <= ql <= qr <= n-1
// 时间复杂度 O(log n)
func (t seg) query(node, l, r, ql, qr int) int {
	if ql > qr {
		return 0
	}
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) / 2
	if qr <= m { // [ql, qr] 在左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	lRes := t.query(node*2, l, m, ql, qr)
	rRes := t.query(node*2+1, m+1, r, ql, qr)
	return gcd(lRes, rRes)
}

func (t seg) check(n int) bool {
	for i := range n {
		if gcd(t.query(1, 0, n-1, 0, i-1), t.query(1, 0, n-1, i+1, n-1)) == targetGcd {
			return true
		}
	}
	return false
}

func countGoodSubseq(nums []int, p int, queries [][]int) (ans int) {
	n := len(nums)
	cntP := 0
	for _, x := range nums {
		if x%p == 0 {
			cntP++
		}
	}

	targetGcd = p
	t := newSegmentTreeWithArray(nums)

	for _, q := range queries {
		i, x := q[0], q[1]

		if nums[i]%p == 0 {
			cntP--
		}
		if x%p == 0 {
			cntP++
		}
		nums[i] = x
		t.update(1, 0, n-1, q[0], x)

		if t[1] == p && (cntP < n || n > 6 || t.check(n)) {
			ans++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
