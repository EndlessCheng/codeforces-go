package main

// https://space.bilibili.com/206214
func differenceOfSum(nums []int) (ans int) {
	for _, x := range nums {
		ans += x // 累加元素和
		for x > 0 {
			ans -= x % 10 // 减去数位和
			x /= 10
		}
	}
	return
}
