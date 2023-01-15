package main

// https://space.bilibili.com/206214
func differenceOfSum(nums []int) (ans int) {
	for _, x := range nums {
		for ans += x; x > 0; x /= 10 {
			ans -= x % 10
		}
	}
	return
}
