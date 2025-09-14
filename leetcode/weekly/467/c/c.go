package main

import "slices"

// https://space.bilibili.com/206214
func subsequenceSumAfterCapping(nums []int, k int) []bool {
	slices.Sort(nums)

	n := len(nums)
	ans := make([]bool, n)
	f := make([]bool, k+1)
	f[0] = true // 不选元素，和为 0

	i := 0
	for x := 1; x <= n; x++ {
		// 增量地考虑所有等于 x 的数
		for i < n && nums[i] == x {
			for j := k; j >= nums[i]; j-- {
				f[j] = f[j] || f[j-nums[i]] // 0-1 背包：不选 or 选
			}
			i++
		}

		// 枚举（从大于 x 的数中）选了 j 个 x
		for j := range min(n-i, k/x) + 1 {
			if f[k-j*x] {
				ans[x-1] = true
				break
			}
		}
	}
	return ans
}
