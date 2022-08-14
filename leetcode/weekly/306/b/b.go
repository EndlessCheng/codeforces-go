package main

// https://space.bilibili.com/206214
func edgeScore(edges []int) (ans int) {
	score := make([]int, len(edges))
	for i, to := range edges {
		score[to] += i
		if score[to] > score[ans] || score[to] == score[ans] && to < ans {
			ans = to
		}
	}
	return
}
