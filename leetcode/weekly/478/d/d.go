package main

import (
	"runtime/debug"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
// 有大量指针的题目，关闭 GC 更快
func init() { debug.SetGCPercent(-1) } 

type node struct {
	lo, ro   *node
	l, r     int
	cnt, sum int
}

func (o *node) maintain() {
	o.cnt = o.lo.cnt + o.ro.cnt
	o.sum = o.lo.sum + o.ro.sum
}

func build(l, r int) *node {
	o := &node{l: l, r: r}
	if l == r {
		return o
	}
	mid := (l + r) / 2
	o.lo = build(l, mid)
	o.ro = build(mid+1, r)
	return o
}

// 在线段树的位置 i 添加 val
// 注意这里传的不是指针，会把 node 复制一份，而这正好是我们需要的
func (o node) add(i, val int) *node {
	if o.l == o.r {
		o.cnt++
		o.sum += val
		return &o
	}
	mid := (o.l + o.r) / 2
	if i <= mid {
		o.lo = o.lo.add(i, val)
	} else {
		o.ro = o.ro.add(i, val)
	}
	o.maintain()
	return &o
}

// 查询 old 和 o 对应区间的第 k 小，k 从 1 开始
func (o *node) kth(old *node, k int) int {
	if o.l == o.r {
		return o.l
	}
	cntL := o.lo.cnt - old.lo.cnt
	if k <= cntL { // 答案在左子树中
		return o.lo.kth(old.lo, k)
	}
	return o.ro.kth(old.ro, k-cntL) // 答案在右子树中
}

// 查询 old 和 o 对应区间，有多少个数 <= i，这些数的元素和是多少
func (o *node) query(old *node, i int) (int, int) {
	if o.r <= i {
		return o.cnt - old.cnt, o.sum - old.sum
	}
	cnt, tot := o.lo.query(old.lo, i)
	mid := (o.l + o.r) / 2
	if i > mid {
		c, t := o.ro.query(old.ro, i)
		cnt += c
		tot += t
	}
	return cnt, tot
}

func minOperations(nums []int, k int, queries [][]int) []int64 {
	n := len(nums)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]%k != nums[i-1]%k {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}

	// 准备离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	t := make([]*node, n+1)
	t[0] = build(0, len(sorted)-1)
	for i, x := range nums {
		j := sort.SearchInts(sorted, x) // 离散化
		t[i+1] = t[i].add(j, x)         // 构建可持久化线段树
	}

	ans := make([]int64, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if left[r] > l { // 无解
			ans[qi] = -1
			continue
		}

		r++ // 改成左闭右开，方便计算

		// 计算区间中位数
		sz := r - l
		i := t[r].kth(t[l], sz/2+1)
		median := sorted[i] // 离散化后的值 -> 原始值

		// 计算区间所有元素到中位数的距离和
		total := t[r].sum - t[l].sum // 区间元素和
		cntLeft, sumLeft := t[r].query(t[l], i)
		sum := median*cntLeft - sumLeft              // 蓝色面积
		sum += total - sumLeft - median*(sz-cntLeft) // 绿色面积
		ans[qi] = int64(sum / k)                     // 操作次数 = 距离和 / k
	}
	return ans
}
