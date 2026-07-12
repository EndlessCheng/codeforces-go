package main

// https://space.bilibili.com/206214
func minimumCost(nums []int, k int) int {
	const mod = 1_000_000_007
	sum := 0 // 总操作次数
	left := k
	for _, x := range nums {
		if left < x {
			op := (x-left-1)/k + 1 // 把 left 增大到 >= x，至少操作 op 次
			sum += op
			left += op * k
		}
		left -= x
	}

	// 1 + 2 + ... + sum
	sum %= mod
	return sum * (sum + 1) / 2 % mod
}
