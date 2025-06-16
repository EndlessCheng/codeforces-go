package main

import "slices"

const mod = 1_000_000_007

var pow2 = [100_000]int{1}

func init() {
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func numSubseq(nums []int, target int) (ans int) {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] <= target {
			// nums[left] 可以作为子序列的最小值
			// 其余下标在 [left+1,right] 中的数选或不选都可以
			ans += pow2[right-left]
			left++
		} else {
			// nums[right] 太大了，即使与剩余元素的最小值 nums[left] 相加也不满足要求
			right--
		}
	}
	return ans % mod
}
