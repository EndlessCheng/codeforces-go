package main

// https://space.bilibili.com/206214
func numberOfGoodPartitions(nums []int) int {
	r := map[int]int{}
	for i, x := range nums {
		r[x] = i
	}
	ans := 1
	maxR := 0
	for i, x := range nums[:len(nums)-1] {
		maxR = max(maxR, r[x])
		if maxR == i {
			ans = ans * 2 % 1_000_000_007
		}
	}
	return ans
}
