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
	dfs = func(x, fa int) (s [14]int) {
		for _, y := range g[x] {
			if y != fa {
				fy := dfs(y, x)
				for j, v := range fy {
					s[j] += v
				}
			}
		}
		for j := range 13 {
			s[j] = max(coins[x]>>j-k+s[j], coins[x]>>(j+1)+s[j+1])
		}
		s[13] = max(s[13]+coins[x]>>13-k, 0)
		return
	}
	return dfs(0, -1)[0]
}

func maximumPoints1(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][14]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, fa int) (res int) {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res1 := coins[i]>>j - k
		res2 := coins[i] >> (j + 1)
		for _, ch := range g[i] {
			if ch != fa {
				res1 += dfs(ch, j, i) // 不右移
				if j < 13 {           // j+1 >= 14 相当于 res2 += 0 无需递归
					res2 += dfs(ch, j+1, i) // 右移
				}
			}
		}
		return max(res1, res2)
	}
	return dfs(0, 0, -1)
}
