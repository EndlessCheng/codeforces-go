package main

// https://space.bilibili.com/206214/dynamic
func numberOfPairs(nums []int) []int {
	pairs := 0
	vis := map[int]bool{}
	for _, v := range nums {
		if vis[v] {
			pairs++
		}
		vis[v] = !vis[v]
	}
	return []int{pairs, len(nums) - pairs*2}
}
