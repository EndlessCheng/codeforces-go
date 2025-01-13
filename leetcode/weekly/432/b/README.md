请先完成不允许感化的版本：[64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/description/)，[讲解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/)。

本题相当于可以不选路径上的至多 $2$ 个数。

**多一个约束，就多一个参数。**

额外增加一个参数 $k$，定义 $\textit{dfs}(i,j,k)$ 表示从 $(i,j)$ 走到 $(0,0)$，在剩余不选次数为 $k$ 的情况下，可以获得的最大金币数。

用「选或不选」分类讨论：

- 选：$\textit{dfs}(i,j,k) = \max(\textit{dfs}(i - 1, j, k), \textit{dfs}(i, j - 1, k)) + \textit{coins}[i][j]$。
- 不选：如果 $k>0$ 且 $\textit{coins}[i][j]<0$，则可以不选，$\textit{dfs}(i,j,k) = \max(\textit{dfs}(i - 1, j, k-1), \textit{dfs}(i, j - 1, k-1))$。

两种情况取最大值。

**递归边界**：

- $\textit{dfs}(-1,j,k)=\textit{dfs}(i,-1,k)=-\infty$。用 $-\infty$ 表示不合法的状态，从而保证 $\max$ 不会取到不合法的状态。
- $\textit{dfs}(0,0,0)=\textit{coins}[0][0]$。
- $\textit{dfs}(0,0,k>0)=\max(\textit{coins}[0][0],0)$。

**递归入口**：$\textit{dfs}(m-1,n-1,2)$，这是原问题，也是答案。

⚠**注意**：由于答案可能是负数，所以记忆化数组的初始值不能是 $-1$，可以初始化成 $-\infty$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1HKcue9ETm/?t=3m51s)，欢迎点赞关注~

## 一、记忆化搜索

```py [sol-Python3]
class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, k: int) -> int:
            if i < 0 or j < 0:
                return -inf
            x = coins[i][j]
            if i == 0 and j == 0:
                return max(x, 0) if k else x
            res = max(dfs(i - 1, j, k), dfs(i, j - 1, k)) + x  # 选
            if k and x < 0:
                res = max(res, dfs(i - 1, j, k - 1), dfs(i, j - 1, k - 1))  # 不选
            return res
        return dfs(len(coins) - 1, len(coins[0]) - 1, 2)
```

```java [sol-Java]
class Solution {
    public int maximumAmount(int[][] coins) {
        int m = coins.length;
        int n = coins[0].length;
        int[][][] memo = new int[m][n][3];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, Integer.MIN_VALUE);
            }
        }
        return dfs(m - 1, n - 1, 2, coins, memo);
    }

    private int dfs(int i, int j, int k, int[][] coins, int[][][] memo) {
        if (i < 0 || j < 0) {
            return Integer.MIN_VALUE;
        }
        int x = coins[i][j];
        if (i == 0 && j == 0) {
            return k > 0 ? Math.max(x, 0) : x;
        }
        if (memo[i][j][k] != Integer.MIN_VALUE) { // 之前计算过
            return memo[i][j][k];
        }
        int res = Math.max(dfs(i - 1, j, k, coins, memo), dfs(i, j - 1, k, coins, memo)) + x; // 选
        if (k > 0 && x < 0) {
            res = Math.max(res, Math.max(dfs(i - 1, j, k - 1, coins, memo), dfs(i, j - 1, k - 1, coins, memo))); // 不选
        }
        return memo[i][j][k] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumAmount(vector<vector<int>>& coins) {
        int m = coins.size(), n = coins[0].size();
        vector memo(m, vector(n, array<int, 3>{INT_MIN, INT_MIN, INT_MIN}));
        auto dfs = [&](this auto&& dfs, int i, int j, int k) -> int {
            if (i < 0 || j < 0) {
                return INT_MIN;
            }
            int x = coins[i][j];
            if (i == 0 && j == 0) {
                return memo[i][j][k] = k ? max(x, 0) : x;
            }
            int& res = memo[i][j][k]; // 注意这里是引用
            if (res != INT_MIN) { // 之前计算过
                return res;
            }
            res = max(dfs(i - 1, j, k), dfs(i, j - 1, k)) + x; // 选
            if (k && x < 0) {
                res = max({res, dfs(i - 1, j, k - 1), dfs(i, j - 1, k - 1)}); // 不选
            }
            return res;
        };
        return dfs(m - 1, n - 1, 2);
    }
};
```

```go [sol-Go]
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][3]int, m)
	for i := range memo {
		memo[i] = make([][3]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		x := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return x
			}
			return max(x, 0)
		}
		p := &memo[i][j][k]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x // 选
		if x < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1)) // 不选
		}
		*p = res // 记忆化
		return res
	}
	return dfs(m-1, n-1, 2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{coins}$ 的行数和列数。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

## 二、1:1 翻译成递推

1:1 地把记忆化搜索翻译成递推，见 [讲解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/)。

代码实现时，可以把 $f[0][1][k]$ 初始化成 $0$，这样我们无需单独计算 $f[1][1]$。

```py [sol-Python3]
class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        m, n = len(coins), len(coins[0])
        f = [[[-inf] * 3 for _ in range(n + 1)] for _ in range(m + 1)]
        f[0][1] = [0] * 3
        for i, row in enumerate(coins):
            for j, x in enumerate(row):
                f[i + 1][j + 1][0] = max(f[i + 1][j][0], f[i][j + 1][0]) + x
                f[i + 1][j + 1][1] = max(f[i + 1][j][1] + x, f[i][j + 1][1] + x,
                                         f[i + 1][j][0], f[i][j + 1][0])
                f[i + 1][j + 1][2] = max(f[i + 1][j][2] + x, f[i][j + 1][2] + x,
                                         f[i + 1][j][1], f[i][j + 1][1])
        return f[m][n][2]
```

```java [sol-Java]
class Solution {
    public int maximumAmount(int[][] coins) {
        int m = coins.length;
        int n = coins[0].length;
        int[][][] f = new int[m + 1][n + 1][3];
        for (int[] row : f[0]) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        Arrays.fill(f[0][1], 0);
        for (int i = 0; i < m; i++) {
            Arrays.fill(f[i + 1][0], Integer.MIN_VALUE);
            for (int j = 0; j < n; j++) {
                int x = coins[i][j];
                f[i + 1][j + 1][0] = Math.max(f[i + 1][j][0], f[i][j + 1][0]) + x;
                f[i + 1][j + 1][1] = Math.max(
                        Math.max(f[i + 1][j][1], f[i][j + 1][1]) + x,
                        Math.max(f[i + 1][j][0], f[i][j + 1][0])
                );
                f[i + 1][j + 1][2] = Math.max(
                        Math.max(f[i + 1][j][2], f[i][j + 1][2]) + x,
                        Math.max(f[i + 1][j][1], f[i][j + 1][1])
                );
            }
        }
        return f[m][n][2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumAmount(vector<vector<int>>& coins) {
        int m = coins.size(), n = coins[0].size();
        vector f(m + 1, vector(n + 1, array<int, 3>{INT_MIN / 2, INT_MIN / 2, INT_MIN / 2}));
        f[0][1] = {0, 0, 0};
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x = coins[i][j];
                f[i + 1][j + 1][0] = max(f[i + 1][j][0], f[i][j + 1][0]) + x;
                f[i + 1][j + 1][1] = max({f[i + 1][j][1] + x, f[i][j + 1][1] + x,
                                          f[i + 1][j][0], f[i][j + 1][0]});
                f[i + 1][j + 1][2] = max({f[i + 1][j][2] + x, f[i][j + 1][2] + x,
                                          f[i + 1][j][1], f[i][j + 1][1]});
            }
        }
        return f[m][n][2];
    }
};
```

```go [sol-Go]
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	f := make([][][3]int, m+1)
	for i := range f {
		f[i] = make([][3]int, n+1)
	}
	for j := range f[0] {
		f[0][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[0][1] = [3]int{}
	for i, row := range coins {
		f[i+1][0] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		for j, x := range row {
			f[i+1][j+1][0] = max(f[i+1][j][0], f[i][j+1][0]) + x
			f[i+1][j+1][1] = max(f[i+1][j][1]+x, f[i][j+1][1]+x, f[i+1][j][0], f[i][j+1][0])
			f[i+1][j+1][2] = max(f[i+1][j][2]+x, f[i][j+1][2]+x, f[i+1][j][1], f[i][j+1][1])
		}
	}
	return f[m][n][2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{coins}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 三、空间优化

举个例子，在计算 $f[1][1]$ 时，会用到 $f[0][1]$，但是之后就不再用到了。那么干脆把 $f[1][1]$ 记到 $f[0][1]$ 中，这样对于 $f[1][2]$ 来说，它需要的数据就在 $f[0][1]$ 和 $f[0][2]$ 中。$f[1][2]$ 算完后也可以同样记到 $f[0][2]$ 中。

所以第一个维度可以去掉。

具体可以看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)中的讲解。本题的转移方程类似完全背包，故整体采用正序遍历（但内部的 $k$ 要倒序）。

```py [sol-Python3]
class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        n = len(coins[0])
        f = [[-inf] * 3 for _ in range(n + 1)]
        f[1] = [0] * 3
        for row in coins:
            for j, x in enumerate(row):
                f[j + 1][2] = max(f[j][2] + x, f[j + 1][2] + x, f[j][1], f[j + 1][1])
                f[j + 1][1] = max(f[j][1] + x, f[j + 1][1] + x, f[j][0], f[j + 1][0])
                f[j + 1][0] = max(f[j][0], f[j + 1][0]) + x
        return f[n][2]
```

```py [sol-Python3 手写 max]
class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        max = lambda a, b: a if a > b else b
        n = len(coins[0])
        f = [[-inf] * 3 for _ in range(n + 1)]
        f[1] = [0] * 3
        for row in coins:
            for j, x in enumerate(row):
                f[j + 1][2] = max(max(f[j][2], f[j + 1][2]) + x, max(f[j][1], f[j + 1][1]))
                f[j + 1][1] = max(max(f[j][1], f[j + 1][1]) + x, max(f[j][0], f[j + 1][0]))
                f[j + 1][0] = max(f[j][0], f[j + 1][0]) + x
        return f[n][2]
```

```java [sol-Java]
class Solution {
    public int maximumAmount(int[][] coins) {
        int n = coins[0].length;
        int[][] f = new int[n + 1][3];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        Arrays.fill(f[1], 0);
        for (int[] row : coins) {
            for (int j = 0; j < n; j++) {
                int x = row[j];
                f[j + 1][2] = Math.max(
                        Math.max(f[j][2], f[j + 1][2]) + x,
                        Math.max(f[j][1], f[j + 1][1])
                );
                f[j + 1][1] = Math.max(
                        Math.max(f[j][1], f[j + 1][1]) + x,
                        Math.max(f[j][0], f[j + 1][0])
                );
                f[j + 1][0] = Math.max(f[j][0], f[j + 1][0]) + x;
            }
        }
        return f[n][2];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumAmount(vector<vector<int>>& coins) {
        int n = coins[0].size();
        vector f(n + 1, array<int, 3>{INT_MIN / 2, INT_MIN / 2, INT_MIN / 2});
        f[1] = {0, 0, 0};
        for (auto& row : coins) {
            for (int j = 0; j < n; j++) {
                int x = row[j];
                f[j + 1][2] = max({f[j][2] + x, f[j + 1][2] + x, f[j][1], f[j + 1][1]});
                f[j + 1][1] = max({f[j][1] + x, f[j + 1][1] + x, f[j][0], f[j + 1][0]});
                f[j + 1][0] = max(f[j][0], f[j + 1][0]) + x;
            }
        }
        return f[n][2];
    }
};
```

```go [sol-Go]
func maximumAmount(coins [][]int) int {
	n := len(coins[0])
	f := make([][3]int, n+1)
	for j := range f {
		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[1] = [3]int{}
	for _, row := range coins {
		for j, x := range row {
			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x, f[j][1], f[j+1][1])
			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x, f[j][0], f[j+1][0])
			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
		}
	}
	return f[n][2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{coins}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单中的「**二、网格图 DP**」。

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
