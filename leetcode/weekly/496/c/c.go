package main

// https://space.bilibili.com/206214
func minIncrease(nums []int) int64 {
	n := len(nums)
	suf := 0
	for i := n - 2; i > 0; i -= 2 {
		suf += max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
	}

	if n%2 > 0 {
		// 修改所有奇数下标
		return int64(suf)
	}

	ans := suf // 修改 [2,n-2] 中的所有偶数下标
	pre := 0
	// 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
	for i := 1; i < n-1; i += 2 {
		pre += max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
		suf -= max(max(nums[i], nums[i+2])-nums[i+1]+1, 0) // 撤销 i+1，撤销后 suf 对应 [i+3,n-2]
		ans = min(ans, pre+suf)
	}

	return int64(ans)
}
