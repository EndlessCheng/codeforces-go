package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func minInversionCount(nums []int, k int) int64 {
	// 离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x) + 1
	}

	t := make(fenwick, len(sorted)+1)
	inv := 0
	ans := math.MaxInt

	for i, in := range nums {
		// 1. 入
		t.update(in, 1)
		inv += min(i+1, k) - t.pre(in) // 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

		left := i + 1 - k
		if left < 0 { // 尚未形成第一个窗口
			continue
		}

		// 2. 更新答案
		ans = min(ans, inv)
		if ans == 0 { // 已经最小了，无需再计算
			break
		}

		// 3. 出
		out := nums[left]
		inv -= t.pre(out - 1)
		t.update(out, -1)
	}
	return int64(ans)
}
