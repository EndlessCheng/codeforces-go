package main

import (
	"cmp"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func countHillValley1(nums []int) (ans int) {
	nums = slices.Compact(nums)
	// 去重后，nums 中的相邻元素都不同
	for i := 1; i < len(nums)-1; i++ {
		if (nums[i-1] < nums[i]) == (nums[i] > nums[i+1]) {
			ans++
		}
	}
	return
}

func countHillValley2(nums []int) (ans int) {
	pre := nums[0] // 上个连续相同段的元素
	for i := 1; i < len(nums)-1; i++ {
		cur := nums[i]   // 当前连续相同段的元素
		nxt := nums[i+1] // 下个连续相同段的元素
		if cur == nxt {  // 同一段
			continue
		}
		// 注意 pre 可能等于 cur，比如 [1,1,2] 中 pre = cur = 1，nxt = 2
		if pre != cur && (pre < cur) == (cur > nxt) {
			ans++
		}
		pre = cur
	}
	return
}

func countHillValley3(nums []int) (ans int) {
	preState := 0
	for i, x := range nums[:len(nums)-1] {
		y := nums[i+1]
		if x > y {
			if preState == -1 { // x 是峰
				ans++
			}
			preState = 1
		} else if x < y {
			if preState == 1 { // x 是谷
				ans++
			}
			preState = -1
		}
	}
	return ans
}

func countHillValley(nums []int) (ans int) {
	preState := 0
	for i, x := range nums[:len(nums)-1] {
		curState := cmp.Compare(x, nums[i+1])
		if curState == 0 {
			continue
		}
		if preState == -curState {
			ans++
		}
		preState = curState
	}
	return ans
}
