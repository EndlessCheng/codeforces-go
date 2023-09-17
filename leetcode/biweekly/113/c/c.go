package main

// https://space.bilibili.com/206214
func countPairs(coordinates [][]int, k int) (ans int) {
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := 0; i <= k; i++ {
			ans += cnt[pair{x ^ i, y ^ (k - i)}]
		}
		cnt[pair{x, y}]++
	}
	return
}
