package main

// github.com/EndlessCheng/codeforces-go
func minTrioDegree(n int, edges [][]int) (ans int) {
	conn := make([][]bool, n)
	for i := range conn {
		conn[i] = make([]bool, n)
	}
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0]-1, e[1]-1
		conn[v][w] = true
		conn[w][v] = true
		deg[v]++
		deg[w]++
	}
	ans = 1e9
	for i, cs := range conn {
		for j := i + 1; j < n; j++ {
			if !cs[j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if cs[k] && conn[j][k] {
					ans = min(ans, deg[i]+deg[j]+deg[k]-6)
				}
			}
		}
	}
	if ans == 1e9 {
		ans = -1
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
