package main

func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const inf int = 1e9
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}
	for _, e := range edges {
		v, w, d := e[0], e[1], e[2]
		dist[v][w] = d
		dist[w][v] = d
	}
	for k := range dist {
		for i := range dist {
			for j := range dist {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	minCnt := inf
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for j, d := range dist[i] {
			if j != i && d <= distanceThreshold {
				cnt++
			}
		}
		if cnt < minCnt {
			minCnt = cnt
			ans = i
		}
	}
	return
}
