package main

import "math/bits"

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 从 1 开始
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 从 1 开始
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

// 不写记忆化更快
func popDepth(x uint64) int {
	if x == 1 {
		return 0
	}
	return popDepth(uint64(bits.OnesCount64(x))) + 1
}

func popcountDepth(nums []int64, queries [][]int64) (ans []int) {
	n := len(nums)
	f := [6]fenwick{}
	for i := range f {
		f[i] = make(fenwick, n+1)
	}
	update := func(i int, x int64, delta int) {
		d := popDepth(uint64(x))
		if d <= 5 {
			f[d].update(i+1, delta)
		}
	}

	for i, x := range nums {
		update(i, x, 1)
	}

	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, f[q[3]].query(int(q[1])+1, int(q[2])+1))
		} else {
			i := int(q[1])
			update(i, nums[i], -1) // 撤销旧的
			nums[i] = q[2]
			update(i, nums[i], 1) // 添加新的
		}
	}
	return
}
