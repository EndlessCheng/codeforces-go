[视频讲解](https://www.bilibili.com/video/BV1ka4y137ua/)（周赛 345 第三题）

# 方法一：动态规划

### 前置知识：动态规划、记忆化搜索

见 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)。

### 记忆化搜索

写一个递归函数 $\textit{dfs}(i,j)$，返回并记录从 $(i,j)$ 出发时的答案。

枚举向右边的三个方向走，如果对应的格子没有出界，且格子值大于 $\textit{grid}[i][j]$，则有

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i-1,j+1)+1,\textit{dfs}(i,j+1)+1,\textit{dfs}(i+1,j+1)+1)
$$

递归边界：$\textit{dfs}(i,n-1)=0$。

递归入口：$\textit{dfs}(i,0)$。取最大值即为答案。

```py [sol1-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        @cache
        def dfs(i: int, j: int) -> int:
            if j == n - 1:
                return 0
            res = 0
            for k in i - 1, i, i + 1:
                if 0 <= k < m and grid[k][j + 1] > grid[i][j]:
                    res = max(res, dfs(k, j + 1) + 1)
            return res
        return max(dfs(i, 0) for i in range(m))
```

```go [sol1-Go]
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j == n-1 {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				res = max(res, dfs(k, j+1)+1)
			}
		}
		*p = res // 记忆化
		return
	}
	for i := 0; i < m; i++ {
		ans = max(ans, dfs(i, 0))
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。

### 1:1 翻译成递推

根据视频中讲的，我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

做法：

- $\textit{dfs}$ 改成 $f$ 数组。
- 递归改成循环（每个参数都对应一层循环）。由于递归是从左向右移动，所以递推是从右到左移动。
- 递归边界改成 $f$ 数组的初始值。本题可以直接把 $f[i][j]$ 初始化为 $0$。

```Python [sol2-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        f = [[0] * n for _ in range(m)]
        for j in range(n - 2, -1, -1):
            for i, row in enumerate(grid):
                for k in i - 1, i, i + 1:
                    if 0 <= k < m and grid[k][j + 1] > row[j]:
                        f[i][j] = max(f[i][j], f[k][j + 1] + 1)
        return max(row[0] for row in f)
```

```java [sol2-Java]
class Solution {
    public int maxMoves(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        var f = new int[m][n];
        for (int j = n - 2; j >= 0; j--)
            for (int i = 0; i < m; i++)
                for (int k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++)
                    if (grid[k][j + 1] > grid[i][j])
                        f[i][j] = Math.max(f[i][j], f[k][j + 1] + 1);
        int ans = 0;
        for (int i = 0; i < m; i++)
            ans = Math.max(ans, f[i][0]);
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int maxMoves(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size(), f[m][n];
        memset(f, 0, sizeof(f));
        for (int j = n - 2; j >= 0; j--)
            for (int i = 0; i < m; i++)
                for (int k = max(i - 1, 0); k < min(i + 2, m); k++)
                    if (grid[k][j + 1] > grid[i][j])
                        f[i][j] = max(f[i][j], f[k][j + 1] + 1);
        int ans = 0;
        for (int i = 0; i < m; i++)
            ans = max(ans, f[i][0]);
        return ans;
    }
};
```

```go [sol2-Go]
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for j := n - 2; j >= 0; j-- {
		for i, row := range grid {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if grid[k][j+1] > row[j] {
					f[i][j] = max(f[i][j], f[k][j+1]+1)
				}
			}
		}
	}
	for _, r := range f {
		ans = max(ans, r[0])
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。

> 注：利用滚动数组，空间可以进一步优化至 $O(m)$。

### 相似题目：网格图 DP

- [62. 不同路径](https://leetcode.cn/problems/unique-paths/)
- [63. 不同路径 II](https://leetcode.cn/problems/unique-paths-ii/)
- [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)
- [120. 三角形最小路径和](https://leetcode.cn/problems/triangle/)
- [931. 下降路径最小和](https://leetcode.cn/problems/minimum-falling-path-sum/)
- [2435. 矩阵中和能被 K 整除的路径](https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/)

# 方法二：BFS

也可以用 BFS 做，每一轮向右搜索一列。

```py [sol3-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        q = range(m)
        vis = [-1] * m
        for j in range(n - 1):
            tmp = q
            q = []
            for i in tmp:
                for k in i - 1, i, i + 1:
                    if 0 <= k < m and vis[k] != j and grid[k][j + 1] > grid[i][j]:
                        vis[k] = j  # 时间戳标记，避免重复创建 vis 数组
                        q.append(k)
            if len(q) == 0:
                return j
        return n - 1
```

```go [sol3-Go]
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vis := make([]int, m)
	q := make([]int, m)
	for i := range q {
		q[i] = i
	}
	for j := 0; j < n-1; j++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if vis[k] != j+1 && grid[k][j+1] > grid[i][j] {
					vis[k] = j + 1 // 时间戳标记，避免重复创建 vis 数组
					q = append(q, k)
				}
			}
		}
		if q == nil {
			return j
		}
	}
	return n - 1
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。
