package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums[i] - nums[j] })

	// rank[i] 表示 nums[i] 是 nums 中的第几小，或者说节点 i 在 idx 中的下标
	rank := make([]int, n)
	for i, j := range idx {
		rank[j] = i
	}

	// 双指针，从第 i 小的数开始，向左一步，最远能跳到第 left 小的数
	pa := make([][]int, n)
	mx := bits.Len(uint(n))
	left := 0
	for i, j := range idx {
		for nums[j]-nums[idx[left]] > maxDiff {
			left++
		}
		pa[i] = make([]int, mx)
		pa[i][0] = left
	}

	// 倍增
	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			pa[x][i+1] = pa[p][i]
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if l == r { // 不用跳
			continue
		}
		l, r = rank[l], rank[r]
		if l > r { // 保证 l 在 r 左边
			l, r = r, l
		}
		// 从 r 开始，向左跳到 l
		res := 0
		for k := mx - 1; k >= 0; k-- {
			if pa[r][k] > l {
				res |= 1 << k
				r = pa[r][k]
			}
		}
		if pa[r][0] > l { // 无法跳到 l
			ans[qi] = -1
		} else {
			ans[qi] = res + 1 // 再跳一步
		}
	}
	return ans
}
