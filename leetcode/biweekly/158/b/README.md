如果只有普通交易，那么本题就是 [188. 买卖股票的最佳时机 IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/)。

回顾一下，188 题中的状态定义为：

- $\textit{dfs}(i,j,0)$ 表示在 $\textit{prices}[0]$ 到 $\textit{prices}[i]$ 中完成至多 $j$ 笔交易，第 $i$ 天结束时**未持有股票**的最大利润。
- $\textit{dfs}(i,j,1)$ 表示在 $\textit{prices}[0]$ 到 $\textit{prices}[i]$ 中完成至多 $j$ 笔交易，第 $i$ 天结束时**持有股票**的最大利润。

本题可以做空交易，也就是可以**先加再减**（普通交易是先减再加）。我们需要额外增加一类状态：

- $\textit{dfs}(i,j,2)$ 表示在 $\textit{prices}[0]$ 到 $\textit{prices}[i]$ 中完成至多 $j$ 笔交易，第 $i$ 天结束时**处于做空中**（空头状态）的最大利润。

同样地，转移方程增加一条：

$$
\textit{dfs}(i,j,2) = \max(\textit{dfs}(i-1, j, 2), \textit{dfs}(i-1, j-1, 0)+\textit{prices}[i])
$$

分别对应继续处于做空中，或者在第 $i$ 天卖出股票，开始做空（开空）。

此外，对于 $\textit{dfs}(i,j,0)$，在计算最大值时额外考虑在第 $i$ 天买回股票（平空）的情况，即 $\textit{dfs}(i,j,2) - \textit{prices}[i]$。

注意本题可能算出负数，$\textit{memo}$ 数组可以初始化成 $\infty$ 或者 $-\infty$。

什么是做空？[简单科普一下](https://www.bilibili.com/video/BV1rET9zsEsB/?t=4m32s)~

## 一、记忆化搜索

```py [sol-Python3]
class Solution:
    def maximumProfit(self, prices: List[int], k: int) -> int:
        # 在 [0,i] 中完成至多 j 笔交易，第 i 天【结束时】的状态为 end_state 的情况下的最大收益
        # 0=未持有股票，1=持有股票，2=做空中
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int, end_state: int) -> int:
            if j < 0:
                return -inf
            if i < 0:
                return -inf if end_state else 0
            p = prices[i]
            if end_state == 0:
                return max(dfs(i - 1, j, 0), dfs(i - 1, j, 1) + p, dfs(i - 1, j, 2) - p)
            if end_state == 1:
                return max(dfs(i - 1, j, 1), dfs(i - 1, j - 1, 0) - p)
            return max(dfs(i - 1, j, 2), dfs(i - 1, j - 1, 0) + p)

        ans = dfs(len(prices) - 1, k, 0)
        dfs.cache_clear()  # 防止爆内存（一般来说，状态数达到 1e6 就需要写这个）
        return ans
```

```java [sol-Java]
class Solution {
    private int[] prices;
    private long[][][] memo;

    public long maximumProfit(int[] prices, int k) {
        this.prices = prices;
        int n = prices.length;
        memo = new long[n][k + 1][3];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, Long.MIN_VALUE); // MIN_VALUE 表示还没有计算过
            }
        }
        return dfs(n - 1, k, 0);
    }

    // 在 [0,i] 中完成至多 j 笔交易，第 i 天【结束时】的状态为 endState 的情况下的最大收益
    // 0=未持有股票，1=持有股票，2=做空中
    private long dfs(int i, int j, int endState) {
        if (j < 0) {
            return Long.MIN_VALUE / 2; // 除 2 防止溢出
        }
        if (i < 0) {
            return endState > 0 ? Long.MIN_VALUE / 2 : 0;
        }
        if (memo[i][j][endState] != Long.MIN_VALUE) { // 之前计算过
            return memo[i][j][endState];
        }
        int p = prices[i];
        if (endState == 0) {
            return memo[i][j][endState] = Math.max(dfs(i - 1, j, 0), Math.max(dfs(i - 1, j, 1) + p, dfs(i - 1, j, 2) - p));
        }
        if (endState == 1) {
            return memo[i][j][endState] = Math.max(dfs(i - 1, j, 1), dfs(i - 1, j - 1, 0) - p);
        }
        return memo[i][j][endState] = Math.max(dfs(i - 1, j, 2), dfs(i - 1, j - 1, 0) + p);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumProfit(vector<int>& prices, int k) {
        int n = prices.size();
        vector memo(n, vector<array<long long, 3>>(k + 1, {LLONG_MIN, LLONG_MIN, LLONG_MIN})); // LLONG_MIN 表示还没有计算过

        // 在 [0,i] 中完成至多 j 笔交易，第 i 天【结束时】的状态为 end_state 的情况下的最大收益
        // 0=未持有股票，1=持有股票，2=做空中
        auto dfs = [&](this auto&& dfs, int i, int j, int end_state) -> long long {
            if (j < 0) {
                return LLONG_MIN / 2; // 除 2 防止溢出
            }
            if (i < 0) {
                return end_state ? LLONG_MIN / 2 : 0;
            }
            long long& res = memo[i][j][end_state]; // 注意这里是引用
            if (res != LLONG_MIN) { // 之前计算过
                return res;
            }
            int p = prices[i];
            if (end_state == 0) {
                return res = max({dfs(i - 1, j, 0), dfs(i - 1, j, 1) + p, dfs(i - 1, j, 2) - p});
            }
            if (end_state == 1) {
                return res = max(dfs(i - 1, j, 1), dfs(i - 1, j - 1, 0) - p);
            }
            return res = max(dfs(i - 1, j, 2), dfs(i - 1, j - 1, 0) + p);
        };

        return dfs(n - 1, k, 0);
    }
};
```

```go [sol-Go]
func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{math.MinInt, math.MinInt, math.MinInt} // MinInt 表示还没有计算过
		}
	}

	// 在 [0,i] 中完成至多 j 笔交易，第 i 天【结束时】的状态为 endState 的情况下的最大收益
	// 0=未持有股票，1=持有股票，2=做空中
	var dfs func(int, int, int) int
	dfs = func(i, j, endState int) (res int) {
		if j < 0 {
			return math.MinInt / 2 // 防止溢出
		}
		if i < 0 {
			if endState == 1 {
				return math.MinInt / 2
			}
			return
		}
		ptr := &memo[i][j][endState]
		if *ptr != math.MinInt { // 之前计算过
			return *ptr
		}
		defer func() { *ptr = res }() // 记忆化
		p := prices[i]
		if endState == 0 {
			return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+p, dfs(i-1, j, 2)-p)
		}
		if endState == 1 {
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-p)
		}
		return max(dfs(i-1, j, 2), dfs(i-1, j-1, 0)+p)
	}

	return int64(dfs(n-1, k, 0))
}
```

## 二、1:1 翻译成递推

```py [sol-Python3]
class Solution:
    def maximumProfit(self, prices: List[int], k: int) -> int:
        n = len(prices)
        f = [[[-inf] * 3 for _ in range(k + 2)] for _ in range(n + 1)]
        for j in range(1, k + 2):
            f[0][j][0] = 0
        for i, p in enumerate(prices):
            for j in range(1, k + 2):
                f[i + 1][j][0] = max(f[i][j][0], f[i][j][1] + p, f[i][j][2] - p)
                f[i + 1][j][1] = max(f[i][j][1], f[i][j - 1][0] - p)
                f[i + 1][j][2] = max(f[i][j][2], f[i][j - 1][0] + p)
        return f[-1][-1][0]
```

```java [sol-Java]
class Solution {
    public long maximumProfit(int[] prices, int k) {
        int n = prices.length;
        long[][][] f = new long[n + 1][k + 2][3];
        for (long[][] mat : f) {
            for (long[] row : mat) {
                Arrays.fill(row, Long.MIN_VALUE / 2);
            }
        }
        for (int j = 1; j <= k + 1; j++) {
            f[0][j][0] = 0;
        }
        for (int i = 0; i < n; i++) {
            int p = prices[i];
            for (int j = 1; j <= k + 1; j++) {
                f[i + 1][j][0] = Math.max(f[i][j][0], Math.max(f[i][j][1] + p, f[i][j][2] - p));
                f[i + 1][j][1] = Math.max(f[i][j][1], f[i][j - 1][0] - p);
                f[i + 1][j][2] = Math.max(f[i][j][2], f[i][j - 1][0] + p);
            }
        }
        return f[n][k + 1][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumProfit(vector<int>& prices, int k) {
        int n = prices.size();
        vector f(n + 1, vector<array<long long, 3>>(k + 2, {LLONG_MIN / 2, LLONG_MIN / 2, LLONG_MIN / 2}));
        for (int j = 1; j <= k + 1; j++) {
            f[0][j][0] = 0;
        }
        for (int i = 0; i < n; i++) {
            int p = prices[i];
            for (int j = 1; j <= k + 1; j++) {
                f[i + 1][j][0] = max({f[i][j][0], f[i][j][1] + p, f[i][j][2] - p});
                f[i + 1][j][1] = max(f[i][j][1], f[i][j - 1][0] - p);
                f[i + 1][j][2] = max(f[i][j][2], f[i][j - 1][0] + p);
            }
        }
        return f[n][k + 1][0];
    }
};
```

```go [sol-Go]
func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	f := make([][][3]int, n+1)
	for i := range f {
		f[i] = make([][3]int, k+2)
		for j := range f[i] {
			f[i][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}
	for i, p := range prices {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p, f[i][j][2]-p)
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
			f[i+1][j][2] = max(f[i][j][2], f[i][j-1][0]+p)
		}
	}
	return int64(f[n][k+1][0])
}
```

## 三、空间优化

```py [sol-Python3]
class Solution:
    def maximumProfit(self, prices: List[int], k: int) -> int:
        f = [[-inf] * 3 for _ in range(k + 2)]
        for j in range(1, k + 2):
            f[j][0] = 0
        for p in prices:
            for j in range(k + 1, 0, -1):
                f[j][0] = max(f[j][0], f[j][1] + p, f[j][2] - p)
                f[j][1] = max(f[j][1], f[j - 1][0] - p)
                f[j][2] = max(f[j][2], f[j - 1][0] + p)
        return f[-1][0]
```

```java [sol-Java]
class Solution {
    public long maximumProfit(int[] prices, int k) {
        long[][] f = new long[k + 2][3];
        for (int j = 1; j <= k + 1; j++) {
            f[j][1] = Long.MIN_VALUE / 2; // 防止溢出
        }
        f[0][0] = Long.MIN_VALUE / 2;
        for (int p : prices) {
            for (int j = k + 1; j > 0; j--) {
                f[j][0] = Math.max(f[j][0], Math.max(f[j][1] + p, f[j][2] - p));
                f[j][1] = Math.max(f[j][1], f[j - 1][0] - p);
                f[j][2] = Math.max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumProfit(vector<int>& prices, int k) {
        vector<array<long long, 3>> f(k + 2, {LLONG_MIN / 2, LLONG_MIN / 2, LLONG_MIN / 2});
        for (int j = 1; j <= k + 1; j++) {
            f[j][0] = 0;
        }
        for (int p : prices) {
            for (int j = k + 1; j > 0; j--) {
                // f[j][0] = max({f[j][0], f[j][1] + p, f[j][2] - p}); 这种写法比下面的慢
                f[j][0] = max(f[j][0], max(f[j][1] + p, f[j][2] - p));
                f[j][1] = max(f[j][1], f[j - 1][0] - p);
                f[j][2] = max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }
};
```

```go [sol-Go]
func maximumProfit(prices []int, k int) int64 {
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
		f[j][2] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 专题训练

见下面动态规划题单的「**六、状态机 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
