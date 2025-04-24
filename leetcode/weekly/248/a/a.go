package main

// github.com/EndlessCheng/codeforces-go
func buildArray(nums []int) []int {
	for i, x := range nums {
		if x < 0 { // 已标记，说明之前计算过
			continue
		}
		cur := i
		for nums[cur] != i {
			nxt := nums[cur]
			nums[cur] = ^nums[nxt] // 把下一个数搬过来，同时做标记（取反）
			cur = nxt
		}
		nums[cur] = ^x // 这一圈的最后一个数，它的下一个数是起点 nums[i]
	}
	for i, x := range nums {
		nums[i] = ^x // 复原
	}
	return nums
}

func buildArray1(nums []int) []int {
	ans := make([]int, len(nums))
	for i, x := range nums {
		ans[i] = nums[x]
	}
	return ans
}
