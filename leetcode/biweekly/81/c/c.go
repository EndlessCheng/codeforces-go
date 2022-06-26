package main

// https://space.bilibili.com/206214/dynamic
func maximumXOR(nums []int) (ans int) {
	for _, num := range nums {
		ans |= num
	}
	return
}
