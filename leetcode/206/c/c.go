package main

// github.com/EndlessCheng/codeforces-go
func minCostConnectPoints(points [][]int) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dis := func(p, q []int) int { return abs(p[0]-q[0]) + abs(p[1]-q[1]) }
	n := len(points)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = dis(points[i], points[j])
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const inf int = 1e7
	minWeights := make([]int, n)
	for i := range minWeights {
		minWeights[i] = inf
	}
	minWeights[0] = 0
	used := make([]bool, n)
	for {
		v := -1
		for i, u := range used {
			if !u && (v == -1 || minWeights[i] < minWeights[v]) {
				v = i
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		ans += minWeights[v]
		for w := range minWeights {
			minWeights[w] = min(minWeights[w], dist[v][w])
		}
	}
	return
}
