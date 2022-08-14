下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

把最大值直接保存在 $3\times 3$ 矩阵的左上角，这样可以无需创建返回矩阵。

```py [sol1-Python3]
class Solution:
    def largestLocal(self, grid: List[List[int]]) -> List[List[int]]:
        n = len(grid)
        for i in range(n - 2):
            for j in range(n - 2):
                grid[i][j] = max(max(row[j:j + 3]) for row in grid[i:i + 3])
            grid[i].pop()
            grid[i].pop()  # 不用切片，避免拷贝
        grid.pop()
        grid.pop()
        return grid
```

```go [sol1-Go]
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i, row := range grid[:n-2] {
		for j := 0; j < n-2; j++ {
			mx := 0
			for _, r := range grid[i : i+3] {
				for _, v := range r[j : j+3] {
					mx = max(mx, v)
				}
			}
			row[j] = mx
		}
		grid[i] = row[:n-2]
	}
	return grid[:n-2]
}

func max(a, b int) int { if b > a { return b }; return a}
```
