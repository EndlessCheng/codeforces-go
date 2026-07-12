package main

// https://space.bilibili.com/206214
func minimumCost(nums []int, k int) int {
	const mod = 1_000_000_007
	total := 0
	for _, x := range nums {
		total += x
	}
	sum := (total - 1) / k % mod
	return sum * (sum + 1) / 2 % mod
}
