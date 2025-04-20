package main

// https://space.bilibili.com/206214
func maximumPossibleSize(nums []int) (ans int) {
	mx := 0
	for _, x := range nums {
		if x >= mx {
			mx = x
			ans++
		}
	}
	return
}
