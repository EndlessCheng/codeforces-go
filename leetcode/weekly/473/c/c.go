package main

// https://space.bilibili.com/206214
func countStableSubarrays(capacity []int) (ans int64) {
	type pair struct{ x, s int }
	cnt := map[pair]int{}
	sum := capacity[0] // 前缀和
	for i := 1; i < len(capacity); i++ {
		ans += int64(cnt[pair{capacity[i], sum}])
		cnt[pair{capacity[i-1], capacity[i-1] + sum}]++
		sum += capacity[i]
	}
	return
}
