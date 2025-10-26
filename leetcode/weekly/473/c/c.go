package main

// https://space.bilibili.com/206214
func countStableSubarrays(capacity []int) (ans int64) {
	type pair struct{ x, s int }
	cnt := map[pair]int{}
	sum := capacity[0] // 前缀和
	for r := 1; r < len(capacity); r++ {
		ans += int64(cnt[pair{capacity[r], sum}])
		cnt[pair{capacity[r-1], capacity[r-1] + sum}]++
		sum += capacity[r]
	}
	return
}
