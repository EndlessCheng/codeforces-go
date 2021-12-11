package main

// 建图 + 暴力枚举所有起点

// github.com/EndlessCheng/codeforces-go
func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	g := make([][]int, n)
	for i, p := range bombs {
		for j, q := range bombs {
			if j != i && (q[0]-p[0])*(q[0]-p[0])+(q[1]-p[1])*(q[1]-p[1]) <= p[2]*p[2] {
				g[i] = append(g[i], j) // 有向图
			}
		}
	}
	for i := range g {
		vis := make([]bool, n)
		cnt := 0
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			cnt++
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		dfs(i)
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
