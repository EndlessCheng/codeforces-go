package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func solve(a []int, low, high int) (res int) {
	if low == high || len(a) <= 1 {
		return
	}

	var lowSt, highSt, b, c []int
	mid := (low + high) / 2

	for i, x := range a {
		if x <= mid { // x 在下部，作为 nums[i]
			for len(lowSt) > 0 && a[lowSt[len(lowSt)-1]] < x {
				lowSt = lowSt[:len(lowSt)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
			}
			lowSt = append(lowSt, i)
			b = append(b, x)
		} else { // x 在上部，作为 nums[j]
			// 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
			for len(highSt) > 0 && a[highSt[len(highSt)-1]] >= x {
				highSt = highSt[:len(highSt)-1]
			}
			res += len(lowSt)
			if len(highSt) > 0 {
				// lowSt 中 < highSt[len(highSt)-1] 的下标不能作为 nums[i]
				res -= sort.SearchInts(lowSt, highSt[len(highSt)-1])
			}
			highSt = append(highSt, i)
			c = append(c, x)
		}
	}

	return res + solve(b, low, mid) + solve(c, mid+1, high)
}

func shadowPairs(nums []int) int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x)
	}

	return solve(nums, 0, len(sorted)-1)
}
