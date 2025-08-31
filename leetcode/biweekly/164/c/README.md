如果没有镜子，本题就是 [62. 不同路径](https://leetcode.cn/problems/unique-paths/)。

请先完成 62 题，本文接着 [我的题解](https://leetcode.cn/problems/unique-paths/solutions/3062432/liang-chong-fang-fa-dong-tai-gui-hua-zu-o5k32/) 继续讲。

和 62 题一样，从终点 $(m-1,n-1)$ 倒着回到起点 $(0,0)$。

如果一个格子有镜子，根据题意，这个镜子是 $\texttt{\textbackslash}$ 摆放的。那么：

- 如果从右边过来，会被反射到上边。
- 如果从下边过来，会被反射到左边。

所以当我们到达格子 $(i,j)$ 时，还需要知道是从哪边过来的，从而决定下一步去哪。

因此，在 62 题的基础上，增加一个参数 $k$：

- $k=0$ 表示从右边来到 $(i,j)$。
- $k=1$ 表示从下边来到 $(i,j)$。

定义 $\textit{dfs}(i,j,k)$，表示在从方向 $k$ 来到 $(i,j)$ 的情况下，从 $(i,j)$ 倒着回到起点 $(0,0)$ 的方案数。

分类讨论：

- 如果 $\textit{grid}[i][j] = 0$，没有镜子，那么：
  - 可以往左到达 $(i,j-1)$，问题变成在从方向 $k=0$ 来到 $(i,j-1)$ 的情况下，从 $(i,j-1)$ 倒着回到起点 $(0,0)$ 的方案数，即 $\textit{dfs}(i,j-1,0)$。
  - 可以往上到达 $(i-1,j)$，问题变成在从方向 $k=1$ 来到 $(i-1,j)$ 的情况下，从 $(i-1,j)$ 倒着回到起点 $(0,0)$ 的方案数，即 $\textit{dfs}(i-1,j,1)$。
  - 这两种情况互斥，根据**加法原理**相加。
- 如果 $\textit{grid}[i][j] = 1$，有镜子，那么：
  - 如果 $k=0$，那么反射到上边，即 $\textit{dfs}(i-1,j,1)$。
  - 如果 $k=1$，那么反射到左边，即 $\textit{dfs}(i,j-1,0)$。

综上所述，我们有

$$
\textit{dfs}(i,j,0) =
\begin{cases}
\textit{dfs}(i,j-1,0) + \textit{dfs}(i-1,j,1), & \textit{grid}[i][j] = 0     \\
\textit{dfs}(i-1,j,1), & \textit{grid}[i][j] = 1     \\
\end{cases}
$$

以及

$$
\textit{dfs}(i,j,1) =
\begin{cases}
\textit{dfs}(i,j-1,0) + \textit{dfs}(i-1,j,1), & \textit{grid}[i][j] = 0     \\
\textit{dfs}(i,j-1,0), & \textit{grid}[i][j] = 1     \\
\end{cases}
$$

**递归边界**：

- $\textit{dfs}(-1,j)=\textit{dfs}(i,-1)=0$。出界，不合法。
- $\textit{dfs}(0,0)=1$。起点到它自己有一条路径，即原地不动。

**递归入口**：$\textit{dfs}(m-1,n-1,0)$。注意题目保证 $\textit{grid}[m-1][n-1] = 0$，所以无需在意 $k$ 的值，$k=0$ 还是 $k=1$ 都可以。

注意取模。为什么可以在计算过程中取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1aCaGzWEm4/?t=8m26s)，欢迎点赞关注~

## 写法一：记忆化搜索

```py [sol-Python3]
class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int, k: int) -> int:
            if i < 0 or j < 0:  # 出界
                return 0
            if i == 0 and j == 0:  # 到达起点
                return 1
            if grid[i][j] == 0:  # 没有镜子，随便走
                return (dfs(i, j - 1, 0) + dfs(i - 1, j, 1)) % 1_000_000_007
            if k == 0:  # 从下边过来
                return dfs(i - 1, j, 1)  # 反射到左边
            # 从右边过来
            return dfs(i, j - 1, 0)  # 反射到上边

        m, n = len(grid), len(grid[0])
        return dfs(m - 1, n - 1, 0)  # 从终点出发
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int uniquePaths(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] memo = new int[m][n][2];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1); // -1 表示没有计算过
            }
        }
        return dfs(m - 1, n - 1, 0, memo, grid); // 从终点出发
    }

    private int dfs(int i, int j, int k, int[][][] memo, int[][] grid) {
        if (i < 0 || j < 0) { // 出界
            return 0;
        }
        if (i == 0 && j == 0) { // 到达起点
            return 1;
        }
        if (memo[i][j][k] != -1) { // 之前计算过
            return memo[i][j][k];
        }
        int res;
        if (grid[i][j] == 0) { // 没有镜子，随便走
            res = (dfs(i, j - 1, 0, memo, grid) + dfs(i - 1, j, 1, memo, grid)) % MOD;
        } else if (k == 0) { // 从下边过来
            res = dfs(i - 1, j, 1, memo, grid); // 反射到左边
        } else { // 从右边过来
            res = dfs(i, j - 1, 0, memo, grid); // 反射到上边
        }
        return memo[i][j][k] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int uniquePaths(vector<vector<int>>& grid) {
        const int MOD = 1'000'000'007;
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector(n, array<int, 2>{-1, -1})); // -1 表示没有计算过

        auto dfs = [&](this auto&& dfs, int i, int j, int k) -> int {
            if (i < 0 || j < 0) { // 出界
                return 0;
            }
            if (i == 0 && j == 0) { // 到达起点
                return 1;
            }
            int& res = memo[i][j][k]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            if (grid[i][j] == 0) { // 没有镜子，随便走
                res = (dfs(i, j - 1, 0) + dfs(i - 1, j, 1)) % MOD;
            } else if (k == 0) { // 从下边过来
                res = dfs(i - 1, j, 1); // 反射到左边
            } else { // 从右边过来
                res = dfs(i, j - 1, 0); // 反射到上边
            }
            return res;
        };

        return dfs(m - 1, n - 1, 0); // 从终点出发
    }
};
```

```go [sol-Go]
func uniquePaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	memo := make([][][2]int, m)
	for i := range memo {
		memo[i] = make([][2]int, n)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i < 0 || j < 0 { // 出界
			return 0
		}
		if i == 0 && j == 0 { // 到达起点
			return 1
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if grid[i][j] == 0 { // 没有镜子，随便走
			return (dfs(i, j-1, 0) + dfs(i-1, j, 1)) % mod
		}
		if k == 0 { // 从下边过来
			return dfs(i-1, j, 1) // 反射到左边
		}
		// 从右边过来
		return dfs(i, j-1, 0) // 反射到上边
	}
	return dfs(m-1, n-1, 0) // 从终点出发
}
```

## 写法二：递推

把记忆化搜索 1:1 翻译成递推，原理见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

```py [sol-Python3]
class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        MOD = 1_000_000_007
        m, n = len(grid), len(grid[0])
        f = [[[0, 0] for _ in range(n + 1)] for _ in range(m + 1)]
        f[0][1] = [1, 1]  # 原理见 62 题我的题解
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x == 0:
                    f[i + 1][j + 1][0] = (f[i + 1][j][0] + f[i][j + 1][1]) % MOD
                    f[i + 1][j + 1][1] = f[i + 1][j + 1][0]
                else:
                    f[i + 1][j + 1][0] = f[i][j + 1][1]
                    f[i + 1][j + 1][1] = f[i + 1][j][0]
        return f[m][n][0]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int uniquePaths(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] f = new int[m + 1][n + 1][2];
        f[0][1][0] = f[0][1][1] = 1; // 原理见 62 题我的题解
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 0) {
                    f[i + 1][j + 1][0] = (f[i + 1][j][0] + f[i][j + 1][1]) % MOD;
                    f[i + 1][j + 1][1] = f[i + 1][j + 1][0];
                } else {
                    f[i + 1][j + 1][0] = f[i][j + 1][1];
                    f[i + 1][j + 1][1] = f[i + 1][j][0];
                }
            }
        }
        return f[m][n][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int uniquePaths(vector<vector<int>>& grid) {
        const int MOD = 1'000'000'007;
        int m = grid.size(), n = grid[0].size();
        vector f(m + 1, vector<array<int, 2>>(n + 1));
        f[0][1] = {1, 1}; // 原理见 62 题我的题解
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 0) {
                    f[i + 1][j + 1][0] = (f[i + 1][j][0] + f[i][j + 1][1]) % MOD;
                    f[i + 1][j + 1][1] = f[i + 1][j + 1][0];
                } else {
                    f[i + 1][j + 1][0] = f[i][j + 1][1];
                    f[i + 1][j + 1][1] = f[i + 1][j][0];
                }
            }
        }
        return f[m][n][0];
    }
};
```

```go [sol-Go]
func uniquePaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	f := make([][][2]int, m+1)
	for i := range f {
		f[i] = make([][2]int, n+1)
	}
	f[0][1] = [2]int{1, 1} // 原理见 62 题我的题解
	for i, row := range grid {
		for j, x := range row {
			if x == 0 {
				f[i+1][j+1][0] = (f[i+1][j][0] + f[i][j+1][1]) % mod
				f[i+1][j+1][1] = f[i+1][j+1][0]
			} else {
				f[i+1][j+1][0] = f[i][j+1][1]
				f[i+1][j+1][1] = f[i+1][j][0]
			}
		}
	}
	return f[m][n][0]
}
```

## 写法三：空间优化

去掉 $f$ 数组的第一个维度。

```py [sol-Python3]
class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        MOD = 1_000_000_007
        n = len(grid[0])
        f = [[0, 0] for _ in range(n + 1)]
        f[1] = [1, 1]
        for row in grid:
            for j, x in enumerate(row):
                if x == 0:
                    f[j + 1][0] = (f[j][0] + f[j + 1][1]) % MOD
                    f[j + 1][1] = f[j + 1][0]
                else:
                    f[j + 1][0] = f[j + 1][1]
                    f[j + 1][1] = f[j][0]
        return f[n][0]
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int uniquePaths(int[][] grid) {
        int n = grid[0].length;
        int[][] f = new int[n + 1][2];
        f[1][0] = f[1][1] = 1;
        for (int[] row : grid) {
            for (int j = 0; j < n; j++) {
                if (row[j] == 0) {
                    f[j + 1][0] = (f[j][0] + f[j + 1][1]) % MOD;
                    f[j + 1][1] = f[j + 1][0];
                } else {
                    f[j + 1][0] = f[j + 1][1];
                    f[j + 1][1] = f[j][0];
                }
            }
        }
        return f[n][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int uniquePaths(vector<vector<int>>& grid) {
        const int MOD = 1'000'000'007;
        int n = grid[0].size();
        vector<array<int, 2>> f(n + 1);
        f[1] = {1, 1};
        for (auto& row : grid) {
            for (int j = 0; j < n; j++) {
                if (row[j] == 0) {
                    f[j + 1][0] = (f[j][0] + f[j + 1][1]) % MOD;
                    f[j + 1][1] = f[j + 1][0];
                } else {
                    f[j + 1][0] = f[j + 1][1];
                    f[j + 1][1] = f[j][0];
                }
            }
        }
        return f[n][0];
    }
};
```

```go [sol-Go]
func uniquePaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	n := len(grid[0])
	f := make([][2]int, n+1)
	f[1] = [2]int{1, 1}
	for _, row := range grid {
		for j, x := range row {
			if x == 0 {
				f[j+1][0] = (f[j][0] + f[j+1][1]) % mod
				f[j+1][1] = f[j+1][0]
			} else {
				f[j+1][0] = f[j+1][1]
				f[j+1][1] = f[j][0]
			}
		}
	}
	return f[n][0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**二、网格图 DP**」和「**六、状态机 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
