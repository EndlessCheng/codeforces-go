下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

DFS 统计每个包含正数的连通块的元素和，最大值即为答案。

```py [sol1-Python3]
class Solution:
    def findMaxFish(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def dfs(i: int, j: int) -> int:
            if i < 0 or i >= m or j < 0 or j >= n or grid[i][j] == 0:
                return 0
            s = grid[i][j]
            grid[i][j] = 0  # 标记成访问过
            for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                s += dfs(x, y)  # 四方向移动
            return s
        return max(max(dfs(i, j) for j in range(n)) for i in range(m))
```

```java [sol1-Java]
class Solution {
    private final static int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    private int[][] grid;

    public int findMaxFish(int[][] grid) {
        this.grid = grid;
        int m = grid.length, n = grid[0].length, ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                ans = Math.max(ans, dfs(i, j));
        return ans;
    }

    private int dfs(int x, int y) {
        int m = grid.length, n = grid[0].length;
        if (x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0)
            return 0;
        int sum = grid[x][y];
        grid[x][y] = 0; // 标记成访问过
        for (var d : DIRS) // 四方向移动
            sum += dfs(x + d[0], y + d[1]);
        return sum;
    }
}
```

```cpp [sol1-C++]
class Solution {
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int findMaxFish(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size(), ans = 0;
        function<int(int, int)> dfs = [&](int x, int y) -> int {
            if (x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0)
                return 0;
            int sum = grid[x][y];
            grid[x][y] = 0; // 标记成访问过
            for (auto &d: dirs) // 四方向移动
                sum += dfs(x + d[0], y + d[1]);
            return sum;
        };
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                ans = max(ans, dfs(i, j));
        return ans;
    }
};
```

```go [sol1-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y]
		grid[x][y] = 0 // 标记成访问过
		for _, d := range dirs { // 四方向移动
			sum += dfs(x+d.x, y+d.y)
		}
		return sum
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。递归需要 $\mathcal{O}(mn)$ 的栈空间。
