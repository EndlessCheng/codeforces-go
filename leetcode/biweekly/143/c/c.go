package main

import (
	"maps"
	"slices"
)

// https://space.bilibili.com/206214
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	slices.Sort(nums)

	n := len(nums)
	var cnt, left, right, left2 int
	for i, x := range nums {
		for nums[left2] < x-k*2 {
			left2++
		}
		ans = max(ans, min(i-left2+1, numOperations))

		cnt++
		// 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
		if i < n-1 && x == nums[i+1] {
			continue
		}
		for nums[left] < x-k {
			left++
		}
		for right < n && nums[right] <= x+k {
			right++
		}
		ans = max(ans, min(right-left, cnt+numOperations))
		cnt = 0
	}

	return
}

func maxFrequency2(nums []int, k, numOperations int) (ans int) {
	cnt := map[int]int{}
	diff := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		diff[x] += 0 // 把 x 插入 diff，以保证下面能遍历到 x
		diff[x-k]++  // 把 [x-k, x+k] 中的每个整数的出现次数都加一
		diff[x+k+1]--
	}

	sumD := 0
	for _, x := range slices.Sorted(maps.Keys(diff)) {
		sumD += diff[x]
		ans = max(ans, min(sumD, cnt[x]+numOperations))
	}
	return
}
