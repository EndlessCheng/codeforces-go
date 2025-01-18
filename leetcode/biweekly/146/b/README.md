**前置知识**：[动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

**前置题目**：[64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/description/)，[讲解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/)。

## 方法一：记忆化搜索

本题由于值域范围小，可以把路径的 XOR 值作为 DP 的第三个参数，即定义 $\textit{dfs}(i,j,x)$ 表示从左上角 $(0,0)$ 到当前位置 $(i,j)$，路径 XOR 值为 $x$ 的方案数。

设从左上角 $(0,0)$ 到 $(i,j-1)$ 的路径 XOR 值为 $y$，那么必须满足

$$
y\oplus \textit{grid}[i][j] = x
$$

即

$$
y = x \oplus \textit{grid}[i][j]
$$

其中 $\oplus$ 表示异或运算。

分类讨论怎么到达 $(i,j)$：

- 如果是从左边过来，根据上文的公式，有 $\textit{dfs}(i,j,x) = \textit{dfs}(i,j-1,x\oplus \textit{grid}[i][j])$。
- 如果是从上边过来，则 $\textit{dfs}(i,j,x) = \textit{dfs}(i-1,j,x\oplus \textit{grid}[i][j])$。

两条路径互斥，根据加法原理，有

$$
\textit{dfs}(i,j,x) = \textit{dfs}(i,j-1,x\oplus \textit{grid}[i][j]) + \textit{dfs}(i-1,j,x\oplus \textit{grid}[i][j])
$$

> 倒序递归是为了方便后面 1:1 翻译成正序的递推。

递归边界：

- $\textit{dfs}(-1,j,x)=\textit{dfs}(i,-1,x)=0$。出界，方案数为 $0$。
- $\textit{dfs}(0,0,\textit{grid}[0][0])=1$，其余 $\textit{dfs}(0,0,x)=0$。左上角 $(0,0)$ 到它自己的方案数是 $1$，路径 XOR 值为 $\textit{grid}[0][0]$。

递归入口：$\textit{dfs}(m-1,n-1,k)$。

### 细节

设 $\textit{mx}$ 为 $\textit{grid}[i][j]$ 的最大值，其二进制长度为 $L$。

那么 $2^L-1$ 就是 XOR 能取到的最大值。

如果 $k > 2^L-1$，那么直接返回 $0$。

> 注：也可以用所有 $\textit{grid}[i][j]$ 的 OR，作为 XOR 可以取到的最大值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ifkqYjEvc/?t=2m19s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countPathsWithXorValue(self, grid: List[List[int]], k: int) -> int:
        MOD = 1_000_000_007
        @cache
        def dfs(i: int, j: int, x: int) -> int:
            if i < 0 or j < 0:
                return 0
            val = grid[i][j]
            if i == 0 and j == 0:
                return 1 if x == val else 0
            return (dfs(i, j - 1, x ^ val) + dfs(i - 1, j, x ^ val)) % MOD
        return dfs(len(grid) - 1, len(grid[0]) - 1, k)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countPathsWithXorValue(int[][] grid, int k) {
        int mx = 0;
        for (int[] row : grid) {
            for (int val : row) {
                mx = Math.max(mx, val);
            }
        }
        int u = 1 << (32 - Integer.numberOfLeadingZeros(mx));
        if (k >= u) {
            return 0;
        }

        int m = grid.length;
        int n = grid[0].length;
        int[][][] memo = new int[m][n][u];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }
        return dfs(grid, m - 1, n - 1, k, memo);
    }

    private int dfs(int[][] grid, int i, int j, int x, int[][][] memo) {
        if (i < 0 || j < 0) {
            return 0;
        }
        int val = grid[i][j];
        if (i == 0 && j == 0) {
            return x == val ? 1 : 0;
        }
        if (memo[i][j][x] != -1) {
            return memo[i][j][x];
        }
        int left = dfs(grid, i, j - 1, x ^ val, memo);
        int up = dfs(grid, i - 1, j, x ^ val, memo);
        return memo[i][j][x] = (left + up) % MOD;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPathsWithXorValue(vector<vector<int>>& grid, int k) {
        const int MOD = 1'000'000'007;
        int mx = 0;
        for (auto& row : grid) {
            mx = max(mx, ranges::max(row));
        }
        int u = 1 << bit_width((unsigned) mx);
        if (k >= u) {
            return 0;
        }

        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector(n, vector<int>(u, -1)));
        auto dfs = [&](this auto&& dfs, int i, int j, int x) -> int {
            if (i < 0 || j < 0) {
                return 0;
            }
            int val = grid[i][j];
            if (i == 0 && j == 0) {
                return x == val;
            }
            int& res = memo[i][j][x]; // 注意这里是引用
            if (res != -1) {
                return res;
            }
            return res = (dfs(i, j - 1, x ^ val) + dfs(i - 1, j, x ^ val)) % MOD;
        };
        return dfs(m - 1, n - 1, k);
    }
};
```

```go [sol-Go]
func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	u := 1 << bits.Len(uint(mx))
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, u)
			for x := range memo[i][j] {
				memo[i][j][x] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, x int) int {
		if i < 0 || j < 0 {
			return 0
		}
		val := grid[i][j]
		if i == 0 && j == 0 {
			if x == val {
				return 1
			}
			return 0
		}
		p := &memo[i][j][x]
		if *p != -1 {
			return *p
		}
		*p = (dfs(i, j-1, x^val) + dfs(i-1, j, x^val)) % mod
		return *p
	}
	return dfs(m-1, n-1, k)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U=\max(\textit{grid}[i][j])$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mnU)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mnU)$。
- 空间复杂度：$\mathcal{O}(mnU)$。保存多少状态，就需要多少空间。

## 方法二：递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j+1][x]$ 的定义和 $\textit{dfs}(i,j,x)$ 的定义是一样的，都表示从左上角 $(0,0)$ 到当前位置 $(i,j)$，路径 XOR 值为 $x$ 的方案数。这里 $+1$ 是为了把 $\textit{dfs}(-1,j,x)$ 和 $\textit{dfs}(i,-1,x)$ 也翻译过来，这样我们可以把 $f[0][j][x]$ 和 $f[i][0][x]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j+1][x] = f[i+1][j][x\oplus \textit{grid}[i][j]] + f[i][j+1][x\oplus \textit{grid}[i][j]]
$$

初始值：

- $f[1][1][\textit{grid}[0][0]]=1$，翻译自递归边界 $\textit{dfs}(0,0,\textit{grid}[0][0])=1$。
- 其余为 $0$。

答案为 $f[m][n][k]$，翻译自递归入口 $\textit{dfs}(m-1,n-1,k)$。

```py [sol-Python3]
class Solution:
    def countPathsWithXorValue(self, grid: List[List[int]], k: int) -> int:
        MOD = 1_000_000_007
        u = 1 << max(map(max, grid)).bit_length()
        if k >= u:
            return 0

        m, n = len(grid), len(grid[0])
        f = [[[0] * u for _ in range(n + 1)] for _ in range(m + 1)]
        f[1][1][grid[0][0]] = 1
        for i, row in enumerate(grid):
            for j, val in enumerate(row):
                for x in range(u):
                    f[i + 1][j + 1][x] += (f[i + 1][j][x ^ val] + f[i][j + 1][x ^ val]) % MOD
        return f[m][n][k]
```

```java [sol-Java]
class Solution {
    public int countPathsWithXorValue(int[][] grid, int k) {
        final int MOD = 1_000_000_007;
        int mx = 0;
        for (int[] row : grid) {
            for (int val : row) {
                mx = Math.max(mx, val);
            }
        }
        int u = 1 << (32 - Integer.numberOfLeadingZeros(mx));
        if (k >= u) {
            return 0;
        }

        int m = grid.length;
        int n = grid[0].length;
        int[][][] f = new int[m + 1][n + 1][u];
        f[1][1][grid[0][0]] = 1;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int val = grid[i][j];
                for (int x = 0; x < u; x++) {
                    f[i + 1][j + 1][x] += (f[i + 1][j][x ^ val] + f[i][j + 1][x ^ val]) % MOD;
                }
            }
        }
        return f[m][n][k];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPathsWithXorValue(vector<vector<int>>& grid, int k) {
        const int MOD = 1'000'000'007;
        int mx = 0;
        for (auto& row : grid) {
            mx = max(mx, ranges::max(row));
        }
        int u = 1 << bit_width((unsigned) mx);
        if (k >= u) {
            return 0;
        }

        int m = grid.size(), n = grid[0].size();
        vector f(m + 1, vector(n + 1, vector<int>(u)));
        f[1][1][grid[0][0]] = 1;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int val = grid[i][j];
                for (int x = 0; x < u; x++) {
                    f[i + 1][j + 1][x] += (f[i + 1][j][x ^ val] + f[i][j + 1][x ^ val]) % MOD;
                }
            }
        }
        return f[m][n][k];
    }
};
```

```go [sol-Go]
func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	u := 1 << bits.Len(uint(mx))
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, u)
		}
	}
	f[1][1][grid[0][0]] = 1
	for i, row := range grid {
		for j, val := range row {
			for x := range u {
				f[i+1][j+1][x] += (f[i+1][j][x^val] + f[i][j+1][x^val]) % mod
			}
		}
	}
	return f[m][n][k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U=\max(\textit{grid}[i][j])$。
- 空间复杂度：$\mathcal{O}(mnU)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
