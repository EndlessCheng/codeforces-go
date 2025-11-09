做法类似 [3418. 机器人可以获得的最大金币数](https://leetcode.cn/problems/maximum-amount-of-money-robot-can-earn/)，[我的题解](https://leetcode.cn/problems/maximum-amount-of-money-robot-can-earn/solutions/3045103/wang-ge-tu-dp-by-endlesscheng-g96j/)。

和 3418 题一样，定义 $\textit{dfs}(i,j,k)$ 表示从 $(0,0)$ 走到 $(i,j)$，在剩余金额为 $k$ 的情况下，可以获得的最大分数。

- 设 $x = \textit{grid}[i][j]$。
- 首先，如果 $x>0$，把 $k$ 减少一。设新的 $k$ 为 $k'$。
- 如果最后一步从 $(i-1,j)$ 走到 $(i,j)$，那么问题变成从 $(0,0)$ 走到 $(i-1,j)$，在剩余金额为 $k'$ 的情况下，可以获得的最大分数，即 $\textit{dfs}(i-1, j, k')$。所以有 $\textit{dfs}(i,j,k) = \textit{dfs}(i-1, j, k') + x$。
- 如果最后一步从 $(i,j-1)$ 走到 $(i,j)$，那么问题变成从 $(0,0)$ 走到 $(i,j-1)$，在剩余金额为 $k'$ 的情况下，可以获得的最大分数，即 $\textit{dfs}(i, j-1, k')$。所以有 $\textit{dfs}(i,j,k) = \textit{dfs}(i, j-1, k') + x$。

两种情况取最大值，得

$$
\textit{dfs}(i,j,k) = \max(\textit{dfs}(i-1, j, k'), \textit{dfs}(i, j-1, k')) + x
$$

**递归边界**：

- 如果 $i,j,k$ 中的任意一个数小于 $0$，不合法，返回 $-\infty$，从而保证 $\max$ 不会取到不合法的状态。
- $\textit{dfs}(0,0,k)=0$。注意题目保证 $\textit{grid}[0][0] = 0$。

**递归入口**：$\textit{dfs}(m-1,n-1,k)$，这是原问题，也是答案。

## 记忆化搜索

原理见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1oskQBLEsY/?t=4m3s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxPathScore(self, grid: List[List[int]], k: int) -> int:
        @cache
        def dfs(i: int, j: int, k: int) -> int:
            if i < 0 or j < 0 or k < 0:  # 出界或者总花费超了
                return -inf
            if i == 0 and j == 0:
                return 0  # 题目保证 grid[0][0] = 0
            x = grid[i][j]
            if x > 0:
                k -= 1
            return max(dfs(i - 1, j, k), dfs(i, j - 1, k)) + x

        ans = dfs(len(grid) - 1, len(grid[0]) - 1, k)
        return -1 if ans < 0 else ans
```

```java [sol-Java]
class Solution {
    public int maxPathScore(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] memo = new int[m][n][k + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }
        int ans = dfs(m - 1, n - 1, k, grid, memo);
        return ans < 0 ? -1 : ans;
    }

    private int dfs(int i, int j, int k, int[][] grid, int[][][] memo) {
        if (i < 0 || j < 0 || k < 0) { // 出界或者总花费超了
            return Integer.MIN_VALUE;
        }
        if (i == 0 && j == 0) {
            return 0; // 题目保证 grid[0][0] = 0
        }
        if (memo[i][j][k] != -1) {
            return memo[i][j][k];
        }
        int x = grid[i][j];
        int newK = x > 0 ? k - 1 : k;
        return memo[i][j][k] = Math.max(dfs(i - 1, j, newK, grid, memo), dfs(i, j - 1, newK, grid, memo)) + x;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPathScore(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector(n, vector<int>(k + 1, -1)));

        auto dfs = [&](this auto&& dfs, int i, int j, int k) -> int {
            if (i < 0 || j < 0 || k < 0) { // 出界或者总花费超了
                return INT_MIN;
            }
            if (i == 0 && j == 0) {
                return 0; // 题目保证 grid[0][0] = 0
            }
            int& res = memo[i][j][k];
            if (res != -1) {
                return res;
            }
            int x = grid[i][j];
            if (x > 0) {
                k--;
            }
            return res = max(dfs(i - 1, j, k), dfs(i, j - 1, k)) + x;
        };

        int ans = dfs(m - 1, n - 1, k);
        return ans < 0 ? -1 : ans;
    }
};
```

```go [sol-Go]
func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 || k < 0 { // 出界或者总花费超了
			return math.MinInt
		}
		if i == 0 && j == 0 {
			return 0 // 题目保证 grid[0][0] = 0
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		x := grid[i][j]
		if x > 0 {
			k--
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x
		*p = res
		return res
	}

	ans := dfs(m-1, n-1, k)
	if ans < 0 {
		return -1
	}
	return ans
}
```

## 递推

把 $f[0][1]$（或者 $f[1][0]$）除了首项都初始化成 $0$，这样 $f[1][1]$ 可以用递推式计算，无需特判。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxPathScore(self, grid: List[List[int]], K: int) -> int:
        m, n = len(grid), len(grid[0])
        f = [[[-inf] * (K + 2) for _ in range(n + 1)] for _ in range(m + 1)]
        f[0][1][1:] = [0] * (K + 1)

        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                for k in range(K + 1):
                    new_k = k - 1 if x else k
                    f[i + 1][j + 1][k + 1] = max(f[i][j + 1][new_k + 1], f[i + 1][j][new_k + 1]) + x

        ans = f[m][n][-1]
        return -1 if ans < 0 else ans
```

```java [sol-Java]
class Solution {
    public int maxPathScore(int[][] grid, int K) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] f = new int[m + 1][n + 1][K + 2];
        for (int[][] mat : f) {
            for (int[] row : mat) {
                Arrays.fill(row, Integer.MIN_VALUE);
            }
        }
        Arrays.fill(f[0][1], 1, K + 2, 0);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x = grid[i][j];
                for (int k = 0; k <= K; k++) {
                    int newK = x > 0 ? k - 1 : k;
                    f[i + 1][j + 1][k + 1] = Math.max(f[i][j + 1][newK + 1], f[i + 1][j][newK + 1]) + x;
                }
            }
        }

        int ans = f[m][n][K + 1];
        return ans < 0 ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPathScore(vector<vector<int>>& grid, int K) {
        int m = grid.size(), n = grid[0].size();
        vector f(m + 1, vector(n + 1, vector<int>(K + 2, INT_MIN)));
        ranges::fill(f[0][1].begin() + 1, f[0][1].end(), 0);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x = grid[i][j];
                for (int k = 0; k <= K; k++) {
                    int new_k = k - (x > 0);
                    f[i + 1][j + 1][k + 1] = max(f[i][j + 1][new_k + 1], f[i + 1][j][new_k + 1]) + x;
                }
            }
        }

        int ans = f[m][n][K + 1];
        return ans < 0 ? -1 : ans;
    }
};
```

```go [sol-Go]
func maxPathScore(grid [][]int, K int) int {
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, K+2)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt
			}
		}
	}
	for k := 1; k < K+2; k++ {
		f[0][1][k] = 0
	}

	for i, row := range grid {
		for j, x := range row {
			for k := range K + 1 {
				newK := k
				if x > 0 {
					newK--
				}
				f[i+1][j+1][k+1] = max(f[i][j+1][newK+1], f[i+1][j][newK+1]) + x
			}
		}
	}

	ans := f[m][n][K+1]
	if ans < 0 {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnk)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mnk)$。

## 空间优化

去掉第一个维度。

为了避免覆盖状态 $f[i][j+1][\textit{newK}+1]$，$k$ 要倒序枚举（类似 0-1 背包）。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxPathScore(self, grid: List[List[int]], K: int) -> int:
        n = len(grid[0])
        f = [[-inf] * (K + 2) for _ in range(n + 1)]
        f[1][1:] = [0] * (K + 1)

        for row in grid:
            for j, x in enumerate(row):
                for k in range(K, -1, -1):
                    new_k = k - 1 if x else k
                    f[j + 1][k + 1] = max(f[j + 1][new_k + 1], f[j][new_k + 1]) + x

        ans = f[n][-1]
        return -1 if ans < 0 else ans
```

```java [sol-Java]
class Solution {
    public int maxPathScore(int[][] grid, int K) {
        int n = grid[0].length;
        int[][] f = new int[n + 1][K + 2];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        Arrays.fill(f[1], 1, K + 2, 0);

        for (int[] row : grid) {
            for (int j = 0; j < n; j++) {
                int x = row[j];
                for (int k = K; k >= 0; k--) {
                    int newK = x > 0 ? k - 1 : k;
                    f[j + 1][k + 1] = Math.max(f[j + 1][newK + 1], f[j][newK + 1]) + x;
                }
            }
        }

        int ans = f[n][K + 1];
        return ans < 0 ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPathScore(vector<vector<int>>& grid, int K) {
        int n = grid[0].size();
        vector f(n + 1, vector<int>(K + 2, INT_MIN));
        ranges::fill(f[1].begin() + 1, f[1].end(), 0);

        for (auto& row : grid) {
            for (int j = 0; j < n; j++) {
                int x = row[j];
                for (int k = K; k >= 0; k--) {
                    int new_k = k - (x > 0);
                    f[j + 1][k + 1] = max(f[j + 1][new_k + 1], f[j][new_k + 1]) + x;
                }
            }
        }

        int ans = f[n][K + 1];
        return ans < 0 ? -1 : ans;
    }
};
```

```go [sol-Go]
func maxPathScore(grid [][]int, K int) int {
	n := len(grid[0])
	f := make([][]int, n+1)
	for j := range f {
		f[j] = make([]int, K+2)
		for k := range f[j] {
			f[j][k] = math.MinInt
		}
	}
	for k := 1; k < K+2; k++ {
		f[1][k] = 0
	}

	for _, row := range grid {
		for j, x := range row {
			for k := K; k >= 0; k-- {
				newK := k
				if x > 0 {
					newK--
				}
				f[j+1][k+1] = max(f[j+1][newK+1], f[j][newK+1]) + x
			}
		}
	}

	ans := f[n][K+1]
	if ans < 0 {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnk)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(nk)$。

## 优化循环次数

从 $(0,0)$ 移动到 $(m-1,n-1)$，至多花费 $m+n-2$（注意题目保证 $\textit{grid}[0][0] = 0$）。所以可以把 $k$ 更新为 $\min(k, m+n-2)$。

此外，从 $(0,0)$ 移动到 $(i,j)$ 至多花费 $i+j$，所以最内层循环的 $k$ 最大是 $\min(k,i+j)$。

改成这种写法后，由于 $f$ 的定义是「至多」，$f[i][j][>i+j]$ 的状态本该更新，但没有更新。所以最后返回的是 $\max(f[m][n])$。

也可以把 $f$ 的定义改成「恰好」，这样只需要把 $f[0][1][1]$ 初始化成 $0$，其余均为 $-\infty$。

此外，可以加一个特判，如果从起点到终点的最小花费都大于 $k$，那么不存在有效路径，返回 $-1$。做法类似 [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)，[我的题解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/)。

> **注**：更精细的写法是，写一个额外的 DP，计算起点到每个位置的最大花费。

```py [sol-Python3]
# 手写 min max 更快
fmin = lambda a, b: b if b < a else a
fmax = lambda a, b: b if b > a else a

class Solution:
    # 64. 最小路径和
    def minPathSum(self, grid: List[List[int]]) -> int:
        f = [inf] * (len(grid[0]) + 1)
        f[1] = 0
        for row in grid:
            for j, x in enumerate(row):
                f[j + 1] = fmin(f[j], f[j + 1]) + fmin(x, 1)  # 值大于 0 的单元格花费 1
        return f[-1]

    def maxPathScore(self, grid: List[List[int]], K: int) -> int:
        if self.minPathSum(grid) > K:
            return -1

        m, n = len(grid), len(grid[0])
        K = fmin(K, m + n - 2)  # 至多花费 m+n-2
        f = [[-inf] * (K + 2) for _ in range(n + 1)]
        f[1][1] = 0

        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                for k in range(fmin(K, i + j), -1, -1):  # 从 (0,0) 到 (i,j) 至多花费 i+j
                    new_k = k - 1 if x else k
                    f[j + 1][k + 1] = fmax(f[j + 1][new_k + 1], f[j][new_k + 1]) + x

        return max(f[n])
```

```java [sol-Java]
class Solution {
    public int maxPathScore(int[][] grid, int K) {
        if (minPathSum(grid) > K) {
            return -1;
        }

        int m = grid.length;
        int n = grid[0].length;
        K = Math.min(K, m + n - 2); // 至多花费 m+n-2
        int[][] f = new int[n + 1][K + 2];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        f[1][1] = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x = grid[i][j];
                for (int k = Math.min(K, i + j); k >= 0; k--) { // 从 (0,0) 到 (i,j) 至多花费 i+j
                    int newK = x > 0 ? k - 1 : k;
                    f[j + 1][k + 1] = Math.max(f[j + 1][newK + 1], f[j][newK + 1]) + x;
                }
            }
        }

        int ans = 0;
        for (int x : f[n]) {
            ans = Math.max(ans, x);
        }
        return ans;
    }

    // 64. 最小路径和
    private int minPathSum(int[][] grid) {
        int n = grid[0].length;
        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE);
        f[1] = 0;
        for (int[] row : grid) {
            for (int j = 0; j < n; j++) {
                f[j + 1] = Math.min(f[j], f[j + 1]) + Math.min(row[j], 1); // 值大于 0 的单元格花费 1
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
    // 64. 最小路径和
    int minPathSum(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> f(n + 1, INT_MAX);
        f[1] = 0;
        for (auto& row : grid) {
            for (int j = 0; j < n; j++) {
                f[j + 1] = min(f[j], f[j + 1]) + min(row[j], 1); // 值大于 0 的单元格花费 1
            }
        }
        return f[n];
    }

public:
    int maxPathScore(vector<vector<int>>& grid, int K) {
        if (minPathSum(grid) > K) {
            return -1;
        }
    
        int m = grid.size(), n = grid[0].size();
        K = min(K, m + n - 2); // 至多花费 m+n-2
        vector f(n + 1, vector<int>(K + 2, INT_MIN));
        f[1][1] = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x = grid[i][j];
                for (int k = min(K, i + j); k >= 0; k--) { // 从 (0,0) 到 (i,j) 至多花费 i+j
                    int new_k = k - (x > 0);
                    f[j + 1][k + 1] = max(f[j + 1][new_k + 1], f[j][new_k + 1]) + x;
                }
            }
        }

        return ranges::max(f[n]);
    }
};
```

```go [sol-Go]
// 64. 最小路径和
func minPathSum(grid [][]int) int {
	n := len(grid[0])
	f := make([]int, n+1)
	for j := range f {
		f[j] = math.MaxInt
	}
	f[1] = 0
	for _, row := range grid {
		for j, x := range row {
			f[j+1] = min(f[j], f[j+1]) + min(x, 1) // 值大于 0 的单元格花费 1
		}
	}
	return f[n]
}

func maxPathScore(grid [][]int, K int) int {
	if minPathSum(grid) > K {
		return -1
	}

	m, n := len(grid), len(grid[0])
	K = min(K, m+n-2) // 至多花费 m+n-2
	f := make([][]int, n+1)
	for j := range f {
		f[j] = make([]int, K+2)
		for k := range f[j] {
			f[j][k] = math.MinInt
		}
	}
	f[1][1] = 0

	for i, row := range grid {
		for j, x := range row {
			for k := min(K, i+j); k >= 0; k-- { // 从 (0,0) 到 (i,j) 至多花费 i+j
				newK := k
				if x > 0 {
					newK--
				}
				f[j+1][k+1] = max(f[j+1][newK+1], f[j][newK+1]) + x
			}
		}
	}

	return slices.Max(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\cdot\min(k,m+n))$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n\cdot\min(k,m+n))$。

## 专题训练

见下面动态规划题单的「**二、网格图 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
