package main

// https://space.bilibili.com/206214
func isPossibleToSplit(nums []int) bool {
	cnt := map[int]int{}
	for _, v := range nums {
		if cnt[v] == 2 {
			return false
		}
		cnt[v]++
	}
	return true
}
