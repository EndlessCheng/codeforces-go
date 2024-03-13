## 方法一：动态规划

**前置知识**：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)。

从第一列的任一单元格出发，比如从 $(2,0)$ 出发，向右下走一步，到 $(3,1)$，那么问题变成从 $(3,1)$ 出发的最大移动次数。

我们要解决的问题，都形如「从 $(i,j)$ 出发的最大移动次数」，这就是递归函数 $\textit{dfs}(i,j)$ 的定义。

枚举向右上/右/右下三个方向走，如果对应的格子没有出界，且格子值大于 $\textit{grid}[i][j]$，则有

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i-1,j+1)+1,\textit{dfs}(i,j+1)+1,\textit{dfs}(i+1,j+1)+1)
$$

递归边界：$\textit{dfs}(i,n-1)=0$。到达最后一列，无法移动。

递归入口：$\textit{dfs}(i,0)$。枚举 $i=0$ 到 $m-1$，取 $\textit{dfs}(i,0)$ 的最大值，即为答案。

```py [sol-Python3]
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

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]
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
```

```js [sol-JavaScript]

```

```rust [sol-Rust]

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。

### 1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示从 $(i,j)$ 出发的最大移动次数。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \max(f[i-1][j+1]+1,\textit{dfs}(i,j+1)+1,\textit{dfs}(i+1,j+1)+1)
$$

初始值 $f[i][n-1]=0$，翻译自递归边界 $\textit{dfs}(i,n-1)=0$。

答案为 $\max\limits_{i=0}^{m-1}f[i][0]$，翻译自递归入口 $\textit{dfs}(i,0)$。

#### 答疑

**问**：如何思考循环顺序？

**答**：这里有一个通用的做法：盯着状态转移方程，想一想，要计算 $f[i][j]$，必须先把右边这一列 $f[\cdot][j+1]$ 算出来，那么只有 $j$ **从大到小**枚举才能做到，并且要在外层枚举 $j$，因为我们是从右往左一列一列计算的。

对于 $i$ 来说，由于在计算 $f[i][j]$ 的时候，$f[\cdot][j+1]$ 这一列已经全部计算完毕，所以 $i$ 无论是正序还是倒序枚举都可以。

```Python [sol-Python3]
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

```java [sol-Java]
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

```cpp [sol-C++]
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

```go [sol-Go]
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
```

```js [sol-JavaScript]

```

```rust [sol-Rust]

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

### 题单：网格图 DP

#### 练习 1

- [62. 不同路径](https://leetcode.cn/problems/unique-paths/)
- [63. 不同路径 II](https://leetcode.cn/problems/unique-paths-ii/)
- [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)
- [120. 三角形最小路径和](https://leetcode.cn/problems/triangle/)
- [2684. 矩阵中移动的最大次数](https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/) 1626
- [1301. 最大得分的路径数目](https://leetcode.cn/problems/number-of-paths-with-max-score/) 1853

#### 练习 2

- [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)
- [2328. 网格图中递增路径的数目](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/) 2001

#### 练习 3

- [1289. 下降路径最小和 II](https://leetcode.cn/problems/minimum-falling-path-sum-ii/) 1697
- [2435. 矩阵中和能被 K 整除的路径](https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/) 1952
- [741. 摘樱桃](https://leetcode.cn/problems/cherry-pickup/)
- [1463. 摘樱桃 II](https://leetcode.cn/problems/cherry-pickup-ii/) 1957

## 方法二：网格图 BFS

也可以用 BFS 做。一开始把所有 $(i,0)$ 都入队，每一轮循环，遍历队列中的当前这一列的坐标，把右边这一列的能到达的格子坐标加入队列，直到队列为空。返回最远可以到达的列号。

```py [sol-Python3]
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
                        vis[k] = j
                        q.append(k)
            if not q:
                return j
        return n - 1
```

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]
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
					vis[k] = j + 1
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
```

```js [sol-JavaScript]

```

```rust [sol-Rust]

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。

### 题单：网格图搜索

- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
