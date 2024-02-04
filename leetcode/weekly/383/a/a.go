package main

// https://space.bilibili.com/206214
func returnToBoundaryCount(nums []int) (ans int) {
	sum := 0
	for _, x := range nums {
		sum += x
		if sum == 0 {
			ans++
		}
	}
	return
}
