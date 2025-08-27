package main

// https://space.bilibili.com/206214
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4]int, m)
	for i := range memo {
		memo[i] = make([][4]int, n)
	}

	var dfs func(int, int, int, bool, int) int
	dfs = func(i, j, k int, canTurn bool, target int) int {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return 0
		}
		// 只在 canTurn=false 时读取和写入 memo
		if !canTurn && memo[i][j][k] > 0 {
			return memo[i][j][k]
		}
		res := dfs(i, j, k, canTurn, 2-target) + 1
		if !canTurn {
			memo[i][j][k] = res
			return res
		}
		maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
		k = (k + 1) % 4
		// 优化二：如果理论最大值没有超过 res，那么不递归
		if min(maxs[k], maxs[(k+3)%4]) > res {
			res = max(res, dfs(i, j, k, false, 2-target)+1)
		}
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x != 1 {
				continue
			}
			maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
			for k, mx := range maxs {                  // 枚举起始方向
				// 优化一：如果理论最大值没有超过 ans，那么不递归
				if mx > ans {
					ans = max(ans, dfs(i, j, k, true, 2)+1)
				}
			}
		}
	}
	return
}
