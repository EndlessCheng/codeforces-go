下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

根据题意，我们可以遍历每个格子，以这个格子为起点，往上下左右四个方向前进，如果下一个格子的值比当前格子的值大，则可以前进。

定义 $f[i][j]$ 表示以第 $i$ 行第 $j$ 列的格子为起点的路径数。

我们把四个方向可以走的格子所对应的状态 $f[i'][j']$ 累加起来，再加上 $1$，即当前格子组成的长度为 $1$ 的路径，即为 $f[i][j]$。

代码实现时可以用记忆化搜索。

```py [sol1-Python3]
class Solution:
    def countPaths(self, grid: List[List[int]]) -> int:
        MOD = 10 ** 9 + 7
        m, n = len(grid), len(grid[0])
        @cache
        def dfs(i: int, j: int) -> int:
            res = 1
            for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                if 0 <= x < m and 0 <= y < n and grid[x][y] > grid[i][j]:
                    res += dfs(x, y)
            return res % MOD
        return sum(dfs(i, j) for i in range(m) for j in range(n)) % MOD
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7;
    static final int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    int m, n;
    int[][] grid, f;

    public int countPaths(int[][] grid) {
        m = grid.length;
        n = grid[0].length;
        this.grid = grid;
        f = new int[m][n];
        for (int i = 0; i < m; i++) Arrays.fill(f[i], -1);
        var ans = 0;
        for (var i = 0; i < m; ++i)
            for (var j = 0; j < n; ++j)
                ans = (ans + dfs(i, j)) % MOD;
        return ans;
    }

    int dfs(int i, int j) {
        if (f[i][j] != -1) return f[i][j];
        var res = 1;
        for (var d : dirs) {
            int x = i + d[0], y = j + d[1];
            if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j])
                res = (res + (dfs(x, y))) % MOD;
        }
        return f[i][j] = res;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
    const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int countPaths(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        int f[m][n]; memset(f, -1, sizeof(f));
        function<int(int, int)> dfs = [&](int i, int j) -> int {
            if (f[i][j] != -1) return f[i][j];
            int res = 1;
            for (auto &d : dirs) {
                int x = i + d[0], y = j + d[1];
                if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j])
                    res = (res + (dfs(x, y))) % MOD;
            }
            return f[i][j] = res;
        };
        int ans = 0;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                ans = (ans + dfs(i, j)) % MOD;
        return ans;
    }
};
```

```go [sol1-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func countPaths(grid [][]int) (ans int) {
	const mod int = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if f[i][j] != -1 {
			return f[i][j]
		}
		res := 1
		for _, d := range dirs {
			if x, y := i+d.x, j+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j] {
				res = (res + dfs(x, y)) % mod
			}
		}
		f[i][j] = res
		return res
	}
	for i, row := range grid {
		for j := range row {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return
}
```
