推荐先完成本题的简单版本 [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)，[我的题解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/)。

本题需要遵循交替模式。由于我们从奇数秒（第 $1$ 秒）开始，所以下一步一定进入相邻单元格，且处于偶数秒，必须等待。等待 $1$ 秒后变成奇数秒，继续移动。

所以本质上来说，题意相当于每个单元格的值为 $\textit{waitCost}[i][j] + (i+1)\cdot (j+1)$，在此基础上求最小路径和。这样就转化成 64 题了。

⚠**注意**：起点和终点无需等待，只需计算进入成本。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

## 一、记忆化搜索

```py [sol-Python3]
class Solution:
    def minCost(self, m: int, n: int, waitCost: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i < 0 or j < 0:
                return inf
            if i == 0 and j == 0:
                return 1  # 起点只有进入成本，不需要等待
            return min(dfs(i, j - 1), dfs(i - 1, j)) + waitCost[i][j] + (i + 1) * (j + 1)
        return dfs(m - 1, n - 1) - waitCost[-1][-1]  # 终点不需要等待
```

```java [sol-Java]
class Solution {
    public long minCost(int m, int n, int[][] waitCost) {
        long[][] memo = new long[m][n];
        return dfs(m - 1, n - 1, waitCost, memo) - waitCost[m - 1][n - 1]; // 终点不需要等待
    }

    private long dfs(int i, int j, int[][] waitCost, long[][] memo) {
        if (i < 0 || j < 0) {
            return Long.MAX_VALUE;
        }
        if (i == 0 && j == 0) {
            return 1; // 起点只有进入成本，不需要等待
        }
        if (memo[i][j] != 0) { // 之前计算过
            return memo[i][j];
        }
        long cost = waitCost[i][j] + (long) (i + 1) * (j + 1);
        return memo[i][j] = Math.min(dfs(i, j - 1, waitCost, memo), dfs(i - 1, j, waitCost, memo)) + cost;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(int m, int n, vector<vector<int>>& waitCost) {
        vector memo(m, vector<long long>(n));
        auto dfs = [&](this auto&& dfs, int i, int j) -> long long {
            if (i < 0 || j < 0) {
                return LLONG_MAX;
            }
            if (i == 0 && j == 0) {
                return 1; // 起点只有进入成本，不需要等待
            }
            long long& res = memo[i][j]; // 注意这里是引用
            if (res != 0) { // 之前计算过
                return res;
            }
            long long cost = waitCost[i][j] + 1LL * (i + 1) * (j + 1);
            return res = min(dfs(i, j - 1), dfs(i - 1, j)) + cost;
        };
        return dfs(m - 1, n - 1) - waitCost[m - 1][n - 1]; // 终点不需要等待
    }
};
```

```go [sol-Go]
func minCost(m, n int, waitCost [][]int) int64 {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return 1 // 起点只有进入成本，不需要等待
		}
		p := &memo[i][j]
		if *p == 0 {
			*p = min(dfs(i, j-1), dfs(i-1, j)) + waitCost[i][j] + (i+1)*(j+1)
		}
		return *p
	}
	return int64(dfs(m-1, n-1) - waitCost[m-1][n-1]) // 终点不需要等待
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

## 二、1:1 翻译成递推

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minCost(self, m: int, n: int, waitCost: List[List[int]]) -> int:
        f = [[inf] * (n + 1) for _ in range(m + 1)]
        f[0][1] = -waitCost[0][0]  # 计算 f[1][1] 的时候抵消掉
        for i, row in enumerate(waitCost):
            for j, c in enumerate(row):
                f[i + 1][j + 1] = min(f[i + 1][j], f[i][j + 1]) + c + (i + 1) * (j + 1)
        return f[m][n] - waitCost[-1][-1]
```

```java [sol-Java]
class Solution {
    public long minCost(int m, int n, int[][] waitCost) {
        long[][] f = new long[m + 1][n + 1];
        Arrays.fill(f[0], Long.MAX_VALUE);
        f[0][1] = -waitCost[0][0]; // 计算 f[1][1] 的时候抵消掉
        for (int i = 0; i < m; i++) {
            f[i + 1][0] = Long.MAX_VALUE;
            for (int j = 0; j < n; j++) {
                f[i + 1][j + 1] = Math.min(f[i + 1][j], f[i][j + 1]) + waitCost[i][j] + (long) (i + 1) * (j + 1);
            }
        }
        return f[m][n] - waitCost[m - 1][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(int m, int n, vector<vector<int>>& waitCost) {
        vector f(m + 1, vector<long long>(n + 1, LLONG_MAX));
        f[0][1] = -waitCost[0][0]; // 计算 f[1][1] 的时候抵消掉
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                f[i + 1][j + 1] = min(f[i + 1][j], f[i][j + 1]) + waitCost[i][j] + 1LL * (i + 1) * (j + 1);
            }
        }
        return f[m][n] - waitCost[m - 1][n - 1];
    }
};
```

```go [sol-Go]
func minCost(m, n int, waitCost [][]int) int64 {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][1] = -waitCost[0][0] // 计算 f[1][1] 的时候抵消掉
	for j := 2; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	for i, row := range waitCost {
		f[i+1][0] = math.MaxInt
		for j, c := range row {
			f[i+1][j+1] = min(f[i+1][j], f[i][j+1]) + c + (i+1)*(j+1)
		}
	}
	return int64(f[m][n] - waitCost[m-1][n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。

## 三、空间优化（一）

去掉 $f$ 的第一维度。

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minCost(self, m: int, n: int, waitCost: List[List[int]]) -> int:
        f = [inf] * (n + 1)
        f[1] = -waitCost[0][0]
        for i, row in enumerate(waitCost):
            for j, c in enumerate(row):
                f[j + 1] = min(f[j], f[j + 1]) + c + (i + 1) * (j + 1)
        return f[n] - waitCost[-1][-1]
```

```java [sol-Java]
class Solution {
    public long minCost(int m, int n, int[][] waitCost) {
        long[] f = new long[n + 1];
        Arrays.fill(f, Long.MAX_VALUE);
        f[1] = -waitCost[0][0];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                f[j + 1] = Math.min(f[j], f[j + 1]) + waitCost[i][j] + (long) (i + 1) * (j + 1);
            }
        }
        return f[n] - waitCost[m - 1][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(int m, int n, vector<vector<int>>& waitCost) {
        vector<long long> f(n + 1, LLONG_MAX);
        f[1] = -waitCost[0][0];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                f[j + 1] = min(f[j], f[j + 1]) + waitCost[i][j] + 1LL * (i + 1) * (j + 1);
            }
        }
        return f[n] - waitCost[m - 1][n - 1];
    }
};
```

```go [sol-Go]
func minCost(m, n int, waitCost [][]int) int64 {
	f := make([]int, n+1)
	for j := range f {
		f[j] = math.MaxInt
	}
	f[1] = -waitCost[0][0]
	for i, row := range waitCost {
		for j, c := range row {
			f[j+1] = min(f[j], f[j+1]) + c + (i+1)*(j+1)
		}
	}
	return int64(f[n] - waitCost[m-1][n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 四、空间优化（二）

直接把 $\textit{waitCost}$ 当作 $f$。

> C++ 和 Java 的数组是 $32$ 位整数，无法实现。

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minCost(self, m: int, n: int, f: List[List[int]]) -> int:
        f[0][0] = 1
        f[-1][-1] = 0
        for j in range(1, n):
            f[0][j] += f[0][j - 1] + j + 1
        for i in range(1, m):
            f[i][0] += f[i - 1][0] + i + 1
            for j in range(1, n):
                f[i][j] += min(f[i][j - 1], f[i - 1][j]) + (i + 1) * (j + 1)
        return f[-1][-1]
```

```go [sol-Go]
func minCost(m, n int, f [][]int) int64 {
	f[0][0] = 1
	f[m-1][n-1] = 0
	for j := 1; j < n; j++ {
		f[0][j] += f[0][j-1] + j + 1
	}
	for i := 1; i < m; i++ {
		f[i][0] += f[i-1][0] + i + 1
		for j := 1; j < n; j++ {
			f[i][j] += min(f[i][j-1], f[i-1][j]) + (i+1)*(j+1)
		}
	}
	return int64(f[m-1][n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
