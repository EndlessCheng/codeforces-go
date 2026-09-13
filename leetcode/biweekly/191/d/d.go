package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func distantSubarrays1(nums []int, goal, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	sorted := slices.Clone(sum)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	ans := n * (n + 1) / 2
	t := newFenwickTree(len(sorted))
	for _, s := range sum {
		// 离散化后的值加一，方便使用树状数组
		l := sort.SearchInts(sorted, s-goal-k+1) + 1
		r := sort.SearchInts(sorted, s-goal+k)
		ans -= t.query(l, r)
		t.update(sort.SearchInts(sorted, s)+1, 1)
	}
	return int64(ans)
}

func distantSubarrays(nums []int, goal, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	
	
}
