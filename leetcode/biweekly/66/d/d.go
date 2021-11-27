package main

/* 动态规划

先求正金字塔。

定义 $\textit{dp}[i][j]$ 表示金字塔顶端位于 $(i,j)$ 时的最大层数（$1$ 层也算）。如果顶端在 $(i,j)$ 的金字塔最大能有 $x$ 层，那么顶端在 $(i,j)$ 的金字塔也可以有 $x-1,x-2,\cdot,\1$ 层。由于要求区域内格子数目大于 $1$，最后统计答案的时候把 $1$ 层去掉，因此 $(i,j)$ 处可以有 $x-1$ 个金字塔。

我们从 $\textit{grid}$ 的最后一行开始往上递推。转移的策略是在金字塔上套一层**倒 V 型的「外壳」**。具体来说，设从 $(i,j)$ 出发向左下方向前进，求出能达到的连续 $1$ 的个数，根据 $\textit{dp}$ 的定义，这就是 $\textit{dp}[i+1][j-1]+1$；同理，向右下方向前进，能达到的连续 $1$ 的个数为 $\textit{dp}[i+1][j+1]+1$。

那么有

$$
\textit{dp}[i][j] = \min(\textit{dp}[i+1][j]+1, \textit{dp}[i+1][j-1]+1, \textit{dp}[i+1][j+1]+1)
$$

倒金字塔可以将 $\textit{grid}$ 上下颠倒后再求一遍正金字塔即可。

*/

// github.com/EndlessCheng/codeforces-go
func countPyramids(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	f := func() {
		dp[m-1] = grid[m-1]
		for i := m - 2; i >= 0; i-- {
			dp[i][0] = grid[i][0]
			dp[i][n-1] = grid[i][n-1]
			for j := 1; j < n-1; j++ {
				if grid[i][j] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = min(min(dp[i+1][j-1], dp[i+1][j]), dp[i+1][j+1]) + 1
					ans += dp[i][j] - 1
				}
			}
		}
	}
	f()
	for i := 0; i < m/2; i++ {
		grid[i], grid[m-1-i] = grid[m-1-i], grid[i]
	}
	f()
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
