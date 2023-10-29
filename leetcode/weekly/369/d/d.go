package main

// https://space.bilibili.com/206214
func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) [14]int
	dfs = func(x, fa int) (res1 [14]int) {
		res2 := [14]int{}
		for _, y := range g[x] {
			if y != fa {
				r := dfs(y, x)
				for j, v := range r {
					res1[j] += v
					if j < 13 {
						res2[j] += r[j+1]
					}
				}
			}
		}
		for j := 0; j < 14; j++ {
			res1[j] = max(res1[j]+coins[x]>>j-k, res2[j]+coins[x]>>(j+1))
		}
		return
	}
	return dfs(0, -1)[0]
}

func max(a, b int) int { if b > a { return b }; return a }
