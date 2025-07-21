package main

import "math/bits"

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
	return f.pre(r) - f.pre(l-1)
}

var popDepthList [51]int

func init() {
	for i := 2; i < len(popDepthList); i++ {
		popDepthList[i] = popDepthList[bits.OnesCount(uint(i))] + 1
	}
}

func popDepth(x uint64) int {
	if x == 1 {
		return 0
	}
	return popDepthList[bits.OnesCount64(x)] + 1
}

func popcountDepth(nums []int64, queries [][]int64) (ans []int) {
	n := len(nums)
	f := [6]fenwick{}
	for i := range f {
		f[i] = newFenwickTree(n)
	}
	update := func(i, delta int) {
		d := popDepth(uint64(nums[i]))
		f[d].update(i+1, delta)
	}

	for i := range n {
		update(i, 1) // 添加
	}

	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, f[q[3]].query(int(q[1])+1, int(q[2])+1))
		} else {
			i := int(q[1])
			update(i, -1) // 撤销旧的
			nums[i] = q[2]
			update(i, 1) // 添加新的
		}
	}
	return
}
