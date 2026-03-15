package main

import "slices"

// https://space.bilibili.com/206214
func calc(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	pre[1] = 2
	for i := 2; i < n; i++ {
		if nums[i-2]+nums[i] == nums[i-1]*2 { // 三个数等差
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 2
		}
	}
	return pre
}

func longestArithmetic(nums []int) (ans int) {
	n := len(nums)
	pre := calc(nums)
	ans = slices.Max(pre) + 1
	if ans >= n { // 整个数组是等差的，或者修改端点元素后是等差的
		return n
	}

	slices.Reverse(nums)
	suf := calc(nums)
	slices.Reverse(suf)
	slices.Reverse(nums)
	// 注意 max(pre) == max(suf)，无需重复计算

	for i := 1; i < n-1; i++ {
		// 把 nums[i] 改成 d2/2
		d2 := nums[i+1] - nums[i-1]
		if d2%2 != 0 { // d2/2 必须是整数
			continue
		}

		okLeft := i > 1 && nums[i-1]-nums[i-2] == d2/2
		okRight := i+2 < n && nums[i+2]-nums[i+1] == d2/2

		if okLeft && okRight {
			ans = max(ans, pre[i-1]+1+suf[i+1])
		} else if okLeft {
			ans = max(ans, pre[i-1]+2)
		} else if okRight {
			ans = max(ans, suf[i+1]+2)
		}
	}

	return ans
}
