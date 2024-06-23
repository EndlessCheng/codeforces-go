package main

// https://space.bilibili.com/206214
func minimumOperations(nums []int) (ans int) {
	for _, x := range nums {
		if x%3 != 0 {
			ans++
		}
	}
	return
}
