package main

// https://space.bilibili.com/206214
func countCompleteDayPairs(hours []int) (ans int) {
	cnt := [24]int{}
	for _, v := range hours {
		v%=24
		ans += cnt[(24-v)%24]
		cnt[v]++
	}
	return
}
