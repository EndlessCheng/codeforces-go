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

func longestArithmetic1(nums []int) (ans int) {
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

func longestArithmetic(nums []int) (ans int) {
	n := len(nums)
	for i := 1; ; {
		// 枚举 i-1 和 i 作为等差子数组的前两项，且我们不改 nums[i-1] 和 nums[i]
		start := i - 1
		d := nums[i] - nums[i-1]

		// 往右移动，直到 nums[i] 不满足等差
		for i++; i < n && nums[i]-nums[i-1] == d; i++ {
		}

		// 现在 [start, i-1] 是等差子数组
		// 要想让子数组更长，要么改 nums[start-1]，要么改 nums[i]

		// 改 nums[start-1]
		if start >= 2 && nums[start]-nums[start-2] == d*2 { // 可以和 nums[start-2] 连起来
			ans = max(ans, i-start+2) // 等差子数组 [start-2, i-1]
			// 继续往左延长的情况等同于上一段继续往右延长，无需重复计算
		} else { // 子数组左端点最远只能到 max(start-1,0)
			ans = max(ans, i-max(start-1, 0)) // 等差子数组 [max(start-1,0), i-1]
		}

		if i == n {
			return
		}

		// 改 nums[i]
		if i < n-1 && nums[i+1]-nums[i-1] == d*2 { // 可以和 nums[i+1] 连起来
			// 继续往右延长
			j := i + 2
			for ; j < n && nums[j]-nums[j-1] == d; j++ {
			}
			ans = max(ans, j-start) // 等差子数组 [start, j-1]
		} else { // 子数组右端点最远只能到 i
			ans = max(ans, i-start+1) // 等差子数组 [start, i]
		}
	}
}
