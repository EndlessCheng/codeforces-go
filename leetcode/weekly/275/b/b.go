package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(nums []int) int {
	tot1 := 0
	for _, num := range nums {
		tot1 += num
	}
	cnt1, maxCnt1 := 0, 0
	nums = append(nums, nums...) // 断环成链
	for i, num := range nums {
		cnt1 += num
		if i >= tot1 { // 滑窗
			cnt1 -= nums[i-tot1]
			if cnt1 > maxCnt1 {
				maxCnt1 = cnt1
			}
		}
	}
	return tot1 - maxCnt1
}
