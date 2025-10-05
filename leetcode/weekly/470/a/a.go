package main

// https://space.bilibili.com/206214
func alternatingSum(nums []int) (ans int) {
	for i, x := range nums {
		ans += x * (1 - i%2*2)
	}
	return
}
