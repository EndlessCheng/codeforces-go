package main

// https://space.bilibili.com/206214
func maximumOr(nums []int, k int) int64 {
	allOr, fixed := 0, 0
	for _, x := range nums {
		// 如果在计算 allOr |= x 之前，allOr 和 x 有公共的 1
		// 那就意味着有多个 nums[i] 在这些比特位上都是 1
		fixed |= allOr & x // 把公共的 1 记录到 fixed 中
		allOr |= x         // 所有数的 OR
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, (allOr^x)|fixed|x<<k)
	}
	return int64(ans)
}

func maximumOr2(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i+1]
	}

	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i])
		pre |= x
	}
	return int64(ans)
}
