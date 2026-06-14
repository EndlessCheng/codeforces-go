package main

import (
	"math"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
var widthM int

type pair struct{ cnt, sum int }
type fenwick []pair

// 添加 num 个 val，其中 val 离散化后的值为 i
// 如果 num < 0，表示减少 -num 个 val
func (t fenwick) update(i, num, val int) {
	for ; i < len(t); i += i & -i {
		t[i].cnt += num
		t[i].sum += val
	}
}

// 返回第 k 小的数（k 从 1 开始）
func (t fenwick) kth(k int, sorted []int) int {
	i := 0
	for b := 1 << (widthM - 1); b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			i = nxt
		}
	}
	return sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (t fenwick) preSum(k int, sorted []int) (res int) {
	i := 0
	for b := 1 << (widthM - 1); b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			res += t[nxt].sum
			i = nxt
		}
	}
	// 加上等于第 k 小的数
	res += sorted[i] * k
	return
}

func maxSum(nums []int, k int) int64 {
	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	widthM = bits.Len(uint(m))

	rank := make([]int, n)        // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allTree := make(fenwick, m+1) // 包含所有元素的树状数组
	total := 0
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		allTree.update(rank[i], 1, x)
		total += x
	}

	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		inTree := make(fenwick, m+1)
		outTree := slices.Clone(allTree)
		needSwap := 0
		subSum := 0

		// 枚举子数组右端点
		for right := left; right < n; right++ {
			// x 从子数组外移到子数组内
			x := nums[right]
			rk := rank[right]
			subSum += x
			inTree.update(rk, 1, x)
			outTree.update(rk, -1, -x)

			inc := false
			sz := right - left + 1
			if needSwap < k && needSwap < sz && needSwap < n-sz {
				// 能否多交换一次
				if inTree.kth(needSwap+1, sorted) < outTree.kth(n-sz-needSwap, sorted) {
					inc = true
					needSwap++
				}
			}

			if !inc && needSwap > 0 {
				// 是否要减少交换次数
				if inTree.kth(needSwap, sorted) >= outTree.kth(n-sz-needSwap+1, sorted) {
					needSwap--
				}
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			if needSwap > 0 {
				inSum := inTree.preSum(needSwap, sorted)
				outSum := total - subSum - outTree.preSum(n-sz-needSwap, sorted)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}

// 优化前
func maxSum1(nums []int, k int) int64 {
	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	rank := make([]int, n)        // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allTree := make(fenwick, m+1) // 包含所有元素的树状数组
	widthM = bits.Len(uint(m))
	total := 0
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		allTree.update(rank[i], 1, x)
		total += x
	}

	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		inTree := make(fenwick, m+1)
		outTree := slices.Clone(allTree)
		needSwap := 0
		subSum := 0

		// 枚举子数组右端点
		for right := left; right < n; right++ {
			// x 从子数组外移到子数组内
			x := nums[right]
			rk := rank[right]
			subSum += x
			inTree.update(rk, 1, x)
			outTree.update(rk, -1, -x)

			inc := false
			sz := right - left + 1
			if needSwap < k && needSwap < sz && needSwap < n-sz {
				// 能否多交换一次
				if inTree.kth(needSwap+1, sorted) < outTree.kth(n-sz-needSwap, sorted) {
					inc = true
					needSwap++
				}
			}

			if !inc && needSwap > 0 {
				// 是否要减少交换次数
				if inTree.kth(needSwap, sorted) >= outTree.kth(n-sz-needSwap+1, sorted) {
					needSwap--
				}
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			if needSwap > 0 {
				inSum := inTree.preSum(needSwap, sorted)
				outSum := total - subSum - outTree.preSum(n-sz-needSwap, sorted)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}
