package main

// https://space.bilibili.com/206214
func sumOfSquares(nums []int) (ans int) {
	for i, x := range nums {
		if len(nums)%(i+1) == 0 {
			ans += x * x
		}
	}
	return
}
