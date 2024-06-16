package main

// https://space.bilibili.com/206214
func countCompleteDayPairs(hours []int) (ans int64) {
	cnt := [24]int{}
	for _, t := range hours {
		// 先查询 cnt，再更新 cnt，因为题目要求 i<j
		// 如果先更新，再查询，就把 i=j 的情况也考虑进去了
		ans += int64(cnt[(24-t%24)%24])
		cnt[t%24]++
	}
	return
}
