下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

按题意模拟即可。

```py [sol1-Python3]
class Solution:
    def maxSum(self, grid: List[List[int]]) -> int:
        return max(grid[i - 1][j - 1] + grid[i - 1][j] + grid[i - 1][j + 1] + grid[i][j] +
                   grid[i + 1][j - 1] + grid[i + 1][j] + grid[i + 1][j + 1]
                   for i in range(1, len(grid) - 1) for j in range(1, len(grid[i]) - 1))
```

```go [sol1-Go]
func maxSum(grid [][]int) (ans int) {
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			ans = max(ans, grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1]+grid[i][j]+grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1])
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(1)$，仅用到若干变量。
