[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

问题等价于把每行排序之后，求每列的最大值。

```py [sol1-Python3]
class Solution:
    def deleteGreatestValue(self, grid: List[List[int]]) -> int:
        for row in grid:
            row.sort()
        return sum(max(col) for col in zip(*grid))
```

```go [sol1-Go]
func deleteGreatestValue(grid [][]int) (ans int) {
	for _, row := range grid {
		sort.Ints(row)
	}
	for j := range grid[0] {
		mx := 0
		for _, row := range grid {
			mx = max(mx, row[j])
		}
		ans += mx
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(mn\log n)$，其中 $m$ 和 $n$ 分别为矩阵 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(1)$，忽略排序的栈开销，仅用到若干额外变量。（Python 忽略 `*` 的开销）
