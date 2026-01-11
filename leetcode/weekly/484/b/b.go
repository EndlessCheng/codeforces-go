package main

// https://space.bilibili.com/206214
func centeredSubarrays(nums []int) (ans int) {
	for i := range nums {
		has := map[int]int{}
		s := 0
		for _, x := range nums[i:] {
			has[x] = 1
			s += x
			ans += has[s]
		}
	}
	return
}
