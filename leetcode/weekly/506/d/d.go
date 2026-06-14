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
	// 特判：能否把正数都聚在一起
	allPosSum := 0
	allPosCnt := 0
	for _, x := range nums {
		if x > 0 {
			allPosSum += x
			allPosCnt++
		}
	}
	if allPosCnt == 0 { // 没有正数
		return int64(slices.Max(nums))
	}
	// 定长滑动窗口模板，窗口长度为 allPosCnt
	cnt := 0
	for i, x := range nums {
		if x > 0 {
			cnt++
		}
		left := i - allPosCnt + 1
		if left < 0 {
			continue
		}
		if cnt+k >= allPosCnt { // 可以把正数都聚在一起
			return int64(allPosSum)
		}
		if nums[left] > 0 {
			cnt--
		}
	}

	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	widthM = bits.Len(uint(m))
	rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allPosTree := make(fenwick, m+1) // 包含所有正数的树状数组
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		if x > 0 {
			allPosTree.update(rank[i], 1, x)
		}
	}

	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		negTree := make(fenwick, m+1)
		posTree := slices.Clone(allPosTree)
		posSum := allPosSum
		posCnt := allPosCnt
		negCnt := 0
		subSum := 0

		// 枚举子数组右端点
		for right := left; right < n; right++ {
			// x 从子数组外移到子数组内
			x := nums[right]
			rk := rank[right]
			subSum += x
			if x > 0 {
				posTree.update(rk, -1, -x)
				posSum -= x
				posCnt--
			} else if x < 0 {
				negTree.update(rk, 1, x)
				negCnt++
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			needSwap := min(negCnt, posCnt, k)
			if needSwap > 0 {
				inSum := negTree.preSum(needSwap, sorted)
				outSum := posSum - posTree.preSum(posCnt-needSwap, sorted)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}
