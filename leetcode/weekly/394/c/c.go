package main

// https://space.bilibili.com/206214
func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f0, f1, pre := 0, 0, -1
	for i := 0; i < n; i++ {
		cnt := [10]int{}
		for _, row := range grid {
			cnt[row[i]]++
		}
		mx, mx2, x := -1, 0, -1
		for v := range cnt {
			res := 0
			if v != pre {
				res = f0
			} else {
				res = f1
			}
			res += cnt[v]
			if res > mx {
				mx2 = mx
				mx = res
				x = v
			} else if res > mx2 {
				mx2 = res
			}
		}
		f0, f1, pre = mx, mx2, x
	}
	return m*n - f0
}

func minimumOperations2(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	cnt := make([][10]int, n)
	for _, row := range grid {
		for j, x := range row {
			cnt[j][x]++
		}
	}

	memo := make([][11]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for k, c := range cnt[i] {
			if k != j {
				res = max(res, dfs(i-1, k)+c)
			}
		}
		*p = res
		return
	}
	return m*n - dfs(n-1, 10)
}
