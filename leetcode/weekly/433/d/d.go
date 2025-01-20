package main

import "math"

// https://space.bilibili.com/206214
func minMaxSubarraySum(nums []int, k int) int64 {
	count := func(m int) int {
		if m <= k {
			return (m + 1) * m / 2
		}
		return (m*2 - k + 1) * k / 2
	}

	// 计算最小值的贡献
	sumSubarrayMins := func() (res int) {
		st := []int{-1} // 哨兵
		for r, x := range nums {
			for len(st) > 1 && nums[st[len(st)-1]] >= x {
				i := st[len(st)-1]
				st = st[:len(st)-1]
				l := st[len(st)-1]
				cnt := count(r-l-1) - count(i-l-1) - count(r-i-1)
				res += nums[i] * cnt // 累加贡献
			}
			st = append(st, r)
		}
		return
	}

	nums = append(nums, math.MinInt)
	ans := sumSubarrayMins()
	// 所有元素取反（求最大值），就可以复用同一份代码了
	for i := range nums {
		nums[i] = -nums[i]
	}
	ans -= sumSubarrayMins()
	return int64(ans)
}

func minMaxSubarraySum1(nums []int, k int) int64 {
	// 计算最小值的贡献
	sumSubarrayMins := func() (res int) {
		n := len(nums)
		// 左边界 left[i] 为左侧严格小于 nums[i] 的最近元素位置（不存在时为 -1）
		left := make([]int, n)
		// 右边界 right[i] 为右侧小于等于 nums[i] 的最近元素位置（不存在时为 n）
		right := make([]int, n)
		st := []int{-1} // 哨兵，方便赋值 left
		for i, x := range nums {
			for len(st) > 1 && x <= nums[st[len(st)-1]] {
				right[st[len(st)-1]] = i // i 是栈顶的右边界
				st = st[:len(st)-1]
			}
			left[i] = st[len(st)-1]
			st = append(st, i)
		}
		for _, i := range st[1:] {
			right[i] = n
		}

		for i, x := range nums {
			l, r := left[i], right[i]
			if r-l-1 <= k {
				cnt := (i - left[i]) * (right[i] - i)
				res += x * cnt // 累加贡献
			} else {
				l = max(l, i-k)
				r = min(r, i+k)
				// 左端点 > r-k 的子数组个数
				cnt := (r - i) * (i - (r - k))
				// 左端点 <= r-k 的子数组个数
				cnt2 := (l + r + k - i*2 + 1) * (r - l - k) / 2
				res += x * (cnt + cnt2) // 累加贡献
			}
		}
		return
	}
	ans := sumSubarrayMins()
	// 所有元素取反（求最大值），就可以复用同一份代码了
	for i := range nums {
		nums[i] = -nums[i]
	}
	ans -= sumSubarrayMins()
	return int64(ans)
}
