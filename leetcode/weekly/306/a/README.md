本题 [视频讲解](https://www.bilibili.com/video/BV1rS4y1s721) 已出炉，欢迎点赞~

---

把最大值直接保存在 $3\times 3$ 矩阵的左上角，这样无需创建返回矩阵。

```py [sol1-Python3]
class Solution:
    def largestLocal(self, grid: List[List[int]]) -> List[List[int]]:
        n = len(grid)
        for i in range(n - 2):
            for j in range(n - 2):
                grid[i][j] = max(max(row[j:j + 3]) for row in grid[i:i + 3])
            del grid[i][-2:]
        del grid[-2:]  # 不要 return grid[:-2] 那样会有额外的拷贝
        return grid
```

```cpp [sol1-C++]
class Solution {
public:
    vector<vector<int>> largestLocal(vector<vector<int>> &grid) {
        int n = grid.size();
        for (int i = 0; i < n - 2; ++i) {
            for (int j = 0; j < n - 2; ++j)
                for (int k = i; k < i + 3; ++k)
                    grid[i][j] = max(grid[i][j], *max_element(grid[k].begin() + j, grid[k].begin() + j + 3));
            grid[i].resize(n - 2);
        }
        grid.resize(n - 2);
        return grid;
    }
};
```

```go [sol1-Go]
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i, row := range grid[:n-2] {
		for j := 0; j < n-2; j++ {
			mx := 0
			for _, r := range grid[i : i+3] {
				for _, x := range r[j : j+3] {
					mx = max(mx, x)
				}
			}
			row[j] = mx
		}
		grid[i] = row[:n-2]
	}
	return grid[:n-2]
}

func max(a, b int) int { if b > a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

#### 思考题

如果把 $3\times 3$ 改成一个比较大的 $h\times w$，你能想出一个非暴力做法吗？

见我的模板库中的[「二维单调队列」](https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/monotone_queue.go#L214)。
