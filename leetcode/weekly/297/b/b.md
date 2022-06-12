定义 $f[i][j]$ 表示从第一行出发到达第 $i$ 行第 $j$ 列时的最小路径代价。

枚举从第 $i-1$ 行的第 $k$ 列转移过来，取最小值，则有

$$
f[i][j] = \textit{grid}[i][j] + \min_{k=0}^{n-1} f[i-1][k] + \textit{moveCost}[\textit{grid}[i-1][k]][j]
$$

答案为 $\min(f[m-1])$。

代码实现时可以用滚动数组优化。

```Python [sol1-Python3]
class Solution:
    def minPathCost(self, grid: List[List[int]], moveCost: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        f = grid[0]
        for i in range(1, m):
            f = [g + min(f[k] + moveCost[v][j] for k, v in enumerate(grid[i - 1])) for j, g in enumerate(grid[i])]
        return min(f)
```

```go [sol1-Go]
func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	pre := grid[0]
	for i := 1; i < m; i++ {
		f := make([]int, n)
		for j, g := range grid[i] {
			f[j] = math.MaxInt32
			for k, v := range grid[i-1] {
				f[j] = min(f[j], pre[k]+moveCost[v][j])
			}
			f[j] += g
		}
		pre = f
	}
	ans := math.MaxInt32
	for _, v := range pre {
		ans = min(ans, v)
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
```
