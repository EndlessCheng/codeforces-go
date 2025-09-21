package main

// https://space.bilibili.com/206214
func evenNumberBitwiseORs(nums []int) (ans int) {
	for _, x := range nums {
		if x%2 == 0 {
			ans |= x
		}
	}
	return
}
