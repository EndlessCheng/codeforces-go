package main

// github.com/EndlessCheng/codeforces-go
func countPaths(n int, roads [][]int) (ans int) {
	type nb struct{ to, wt int }
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e18
		}
	}
	for _, r := range roads {
		v, w, wt := r[0], r[1], r[2]
		g[v][w] = wt
		g[w][v] = wt
	}

	// 求 0 到其余点的最短路
	d := make([]int, n)
	for i := range d {
		d[i] = 1e18
	}
	d[0] = 0
	used := make([]bool, n)
	for {
		v := -1
		for w, u := range used {
			if !u && (v < 0 || d[w] < d[v]) {
				v = w
			}
		}
		if v < 0 {
			break
		}
		used[v] = true
		for w, wt := range g[v] {
			if newD := d[v] + wt; newD < d[w] {
				d[w] = newD
			}
		}
	}

	// 最短路构成了一个 DAG，这里不需要建一个新图，直接根据距离来判断每条边是否在 DAG 上
	// 计算新图的入度数组
	deg := make([]int, n)
	for v, r := range g {
		for w, wt := range r {
			if d[v]+wt == d[w] {
				deg[w]++
			}
		}
	}

	// 在新图上跑拓扑排序
	dp := make([]int, n) // dp[i] 表示 0 到 i 的最短路个数
	dp[0] = 1
	q := []int{0}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for w, wt := range g[v] {
			if d[v]+wt == d[w] {
				dp[w] = (dp[w] + dp[v]) % (1e9 + 7)
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
	}
	return dp[n-1]
}
