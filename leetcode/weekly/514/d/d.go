package main

import "math/bits"

// https://space.bilibili.com/206214
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
type data struct {
	cnt, pre, suf, len int
	hasPeak            bool
}

type seg []data

func newSegmentTree(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func (seg) mergeData(a, b data) data {
	pre := a.pre
	if !a.hasPeak {
		pre += b.pre
	}
	suf := b.suf
	if !b.hasPeak {
		suf += a.suf
	}
	return data{a.cnt + b.cnt + a.len*b.len - a.suf*b.pre, pre, suf, a.len + b.len, a.hasPeak || b.hasPeak}
}

func (t seg) maintain(node int) {
	t[node] = t.mergeData(t[node*2], t[node*2+1])
}

func (t seg) build(a []int, node, l, r int) {
	if l == r { // 叶子
		hasPeak := 0 < l && l < len(a)-1 && a[l-1] < a[l] && a[l] > a[l+1]
		t[node] = data{0, 1, 1, 1, hasPeak} // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

func (t seg) update(node, l, r, i int, hasPeak bool) {
	if l == r { // 叶子（到达目标）
		t[node].hasPeak = hasPeak
		return
	}
	m := (l + r) >> 1
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i, hasPeak)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i, hasPeak)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) >> 1
	if qr <= m { // [ql, qr] 与右子树无交集，仅需递归左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 与左子树无交集，仅需递归右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	// [ql, qr] 与左右子树均有交集，分别递归，然后合并结果
	return t.mergeData(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func countOfPeaks(nums []int, queries [][]int) (ans []int64) {
	n := len(nums)
	t := newSegmentTree(nums)
	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, int64(t.query(1, 0, n-1, q[1], q[2]).cnt))
			continue
		}
		i := q[1]
		nums[i] = q[2]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			// 注：这里可以优化一下，如果更新前后 hasPeak 不变，则不调用 t.update
			hasPeak := nums[j-1] < nums[j] && nums[j] > nums[j+1]
			t.update(1, 0, n-1, j, hasPeak)
		}
	}
	return
}
