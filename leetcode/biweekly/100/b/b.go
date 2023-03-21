package main

// https://space.bilibili.com/206214
func maximizeGreatness(nums []int) int {
	maxCnt := 0
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
		maxCnt = max(maxCnt, cnt[v])
	}
	return len(nums) - maxCnt
}

func max(a, b int) int { if a < b { return b }; return a }
