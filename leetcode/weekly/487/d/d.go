package main

import (
	"cmp"
	"slices"
)

// https://space.bilibili.com/206214
func calc(a []int) []int {
	f := make([]int, len(a)) // f[i] 表示以 i 结尾的最长交替子数组的长度
	for i, x := range a {
		if i == 0 || a[i-1] == x {
			f[i] = 1
		} else if i > 1 && a[i-2] != a[i-1] && (a[i-2] < a[i-1]) == (a[i-1] > x) {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 2
		}
	}
	return f
}

func longestAlternating1(nums []int) int {
	n := len(nums)
	pre := calc(nums) // pre[i] 表示以 i 结尾的最长交替子数组的长度

	slices.Reverse(nums)
	suf := calc(nums) // suf[i] 表示以 i 开头的最长交替子数组的长度
	slices.Reverse(suf)
	slices.Reverse(nums)

	// 不删除元素时的最长交替子数组的长度
	ans := slices.Max(pre)

	// 枚举删除 nums[i]
	for i := 1; i < n-1; i++ {
		if nums[i-1] == nums[i+1] { // 无法拼接
			continue
		}

		// 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
		x := 0
		if i > 1 {
			x = cmp.Compare(nums[i-2], nums[i-1])
		}

		y := cmp.Compare(nums[i-1], nums[i+1])

		z := 0
		if i < n-2 {
			z = cmp.Compare(nums[i+1], nums[i+2])
		}

		if x == -y && x == z { // 左右两边可以拼接
			ans = max(ans, pre[i-1]+suf[i+1])
		} else {
			if x == -y {
				ans = max(ans, pre[i-1]+1) // 只拼接 nums[i+1] 
			}
			if z == -y {
				ans = max(ans, suf[i+1]+1) // 只拼接 nums[i-1] 
			}
		}
	}

	return ans
}

func longestAlternating(a []int) int {
	n := len(a)
	f := make([][2][2]int, n)
	for i := range f {
		f[i] = [2][2]int{{1, 1}, {1, 1}}
	}

	ans := 1
	for i := 1; i < n; i++ {
		if a[i-1] != a[i] {
			inc := 0
			if a[i-1] < a[i] {
				inc = 1
			}
			f[i][0][inc] = f[i-1][0][inc^1] + 1
			f[i][1][inc] = f[i-1][1][inc^1] + 1
		}
		if i > 1 && a[i-2] != a[i] {
			inc := 0
			if a[i-2] < a[i] {
				inc = 1
			}
			f[i][1][inc] = max(f[i][1][inc], f[i-2][0][inc^1]+1)
		}
		ans = max(ans, f[i][1][0], f[i][1][1])
	}
	return ans
}
