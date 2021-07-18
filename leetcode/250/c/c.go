package main

/* DP 优化技巧：拆项+前后缀最大值

本文用 $n$ 表示行数，$m$ 表示列数。

定义 $f[i][j]$ 表示前 $i$ 行中，第 $i$ 行选择 $\textit{points}[i][j]$ 时的最大得分，则有

$$
f[i][j] = \textit{points}[i][j] + \max f[i-1][k] - |k-j|
$$

由于转移是 $O(m)$ 的，所以总体复杂度是 $O(nm^2)$ 的，我们需要优化。

拆掉绝对值符号，将上式变形为

$$
f[i][j] =
\begin{cases}
\textit{points}[i][j] + \max f[i-1][k] - (j - k),&k\le j\\
\textit{points}[i][j] + \max f[i-1][k] - (k - j),&k > j
\end{cases}
$$

将 $j$ 提出来，化简为

$$
f[i][j] =
\begin{cases}
\textit{points}[i][j] - j + \max f[i-1][k] + k,&k\le j\\
\textit{points}[i][j] + j + \max f[i-1][k] - k,&k > j
\end{cases}
$$

由上式可知，在计算 $f[i][j]$ 时，我们需要知道位置 $j$ 左侧的 $f[i-1][k] + k$ 的最大值，以及位置 $j$ 右侧的 $f[i-1][k] - k$ 的最大值。这可以在计算完一整行 $f[i-1][]$ 之后，在计算下一行 $f[i][]$ 之前，预处理出来。

代码实现时，$f$ 的第一维可以压缩掉，且预处理过程可以只处理 $f[i-1][k] - k$ 的最大值，$f[i-1][k] + k$ 的最大值可以一边遍历 $\textit{points}[i][]$ 一边计算。

*/

// github.com/EndlessCheng/codeforces-go
func maxPoints(points [][]int) int64 {
	ans := 0
	m := len(points[0])
	f := make([][2]int, m)
	sufMax := make([]int, m) // 后缀最大值
	for i, row := range points {
		if i == 0 {
			for j, v := range row {
				ans = max(ans, v)
				f[j][0] = v + j
				f[j][1] = v - j
			}
		} else {
			preMax := int(-1e9)
			for j, v := range row {
				preMax = max(preMax, f[j][0])
				res := max(v-j+preMax, v+j+sufMax[j]) // 左侧和右侧的最大值即为选择 points[i][j] 时的计算结果
				ans = max(ans, res) // 直接更新答案，这样下面就不直接存储 res 了，改为存储 res + j 和 res - j
				f[j][0] = res + j
				f[j][1] = res - j
			}
		}
		sufMax[m-1] = f[m-1][1]
		for j := m - 2; j >= 0; j-- {
			sufMax[j] = max(sufMax[j+1], f[j][1])
		}
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
