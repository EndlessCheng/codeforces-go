package main

// https://space.bilibili.com/206214
func hasTrailingZeros(nums []int) bool {
	even := len(nums)
	for _, x := range nums {
		even -= x % 2
	}
	return even >= 2
}
