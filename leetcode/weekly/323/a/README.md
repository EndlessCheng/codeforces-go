[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/)

问题等价于把每行排序之后，求每列的最大值。

```py [sol-Python3]
class Solution:
    def deleteGreatestValue(self, grid: List[List[int]]) -> int:
        for row in grid:
            row.sort()
        return sum(map(max, zip(*grid)))  # 累加每列的最大值
```

```java [sol-Java]
class Solution {
    public int deleteGreatestValue(int[][] grid) {
        for (var row : grid)
            Arrays.sort(row);
        int ans = 0, n = grid[0].length;
        for (int j = 0; j < n; j++) {
            int mx = 0;
            for (var row : grid)
                mx = Math.max(mx, row[j]);
            ans += mx;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int deleteGreatestValue(vector<vector<int>> &grid) {
        for (auto &row: grid)
            sort(row.begin(), row.end());
        int ans = 0, n = grid[0].size();
        for (int j = 0; j < n; j++) {
            int mx = 0;
            for (auto &row: grid)
                mx = max(mx, row[j]);
            ans += mx;
        }
        return ans;
    }
};
```

```go [sol-Go]
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

```js [sol-JavaScript]
var deleteGreatestValue = function (grid) {
    for (let row of grid)
        row.sort((a, b) => a - b);
    let ans = 0;
    const n = grid[0].length;
    for (let j = 0; j < n; j++) {
        let mx = 0;
        for (const row of grid)
            mx = Math.max(mx, row[j]);
        ans += mx;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log n)$，其中 $m$ 和 $n$ 分别为矩阵 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$，忽略排序的栈开销，仅用到若干额外变量。（Python 忽略 `*` 的开销）
