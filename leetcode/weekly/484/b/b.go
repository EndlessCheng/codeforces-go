package main

// https://space.bilibili.com/206214
func centeredSubarrays(nums []int) (ans int) {
	has := map[int]int{}
	for i := range nums {
		clear(has)
		s := 0
		for _, x := range nums[i:] {
			has[x] = 1
			s += x
			ans += has[s]
		}
	}
	return
}
