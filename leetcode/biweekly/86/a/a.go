package main

// https://space.bilibili.com/206214
func findSubarrays(nums []int) bool {
	vis := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		s := nums[i-1] + nums[i]
		if vis[s] {
			return true
		}
		vis[s] = true
	}
	return false
}
