package main

// https://space.bilibili.com/206214/dynamic
func zeroFilledSubarray(nums []int) (ans int64) {
	c := 0
	for _, num := range nums {
		if num == 0 {
			c++
			ans += int64(c)
		} else {
			c = 0
		}
	}
	return
}
