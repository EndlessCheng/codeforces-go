下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

按要求模拟。

优化 1：由于行和列都可以看成是 $1$ 的个数减去 $0$ 的个数，所以统计的时候，可以把 $0$ 当成 $-1$。

优化 2：答案可以直接填到 $\textit{grid}$ 中。

```py [sol1-Python3]
class Solution:
    def onesMinusZeros(self, grid: List[List[int]]) -> List[List[int]]:
        r = [0] * len(grid)
        c = [0] * len(grid[0])
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                r[i] += x * 2 - 1
                c[j] += x * 2 - 1  # 1 -> 1, 0 -> -1
        for i, x in enumerate(r):
            for j, y in enumerate(c):
                grid[i][j] = x + y
        return grid
```

```go [sol1-Go]
func onesMinusZeros(grid [][]int) [][]int {
	r := make([]int, len(grid))
	c := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, x := range row {
			r[i] += x*2 - 1
			c[j] += x*2 - 1 // 1 -> 1, 0 -> -1
		}
	}
	for i, x := range r {
		for j, y := range c {
			grid[i][j] = x + y
		}
	}
	return grid
}
```

#### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(m+n)$。
