## 方法一：付费与免费的时间差

用「选或不选」的来思考，即是否付费刷墙。

考虑第 $n-1$ 堵墙是否付费刷：

- 选择付费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，在「付费时间之和为 $\textit{time}[n-1]$，免费时间之和为 $0$」状态下的最少开销。
- 选择不付费，即免费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，在「付费时间之和为 $0$，免费时间之和为 $1$」状态下的最少开销。

这启发我们定义 $\textit{dfs}(i,j,k)$ 表示刷前 $i$ 堵墙，在「付费时间之和为 $j$，免费时间之和为 $k$」状态下的最少开销。

递归到终点时，如果 $j\ge k$，说明这种方案是合法的，否则不合法。

但这样定义的话，状态个数就太多了，需要优化。

由于最后是比较的 $j$ 和 $k$ 的「相对大小」，那么不妨把 $j$ 重新定义为「付费时间之和」减去「免费时间之和」，这样递归到终点时，如果 $j\ge 0$，说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。

分类讨论：

- 选择付费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j+\textit{time}[i])+\textit{cost}[i]$。
- 选择免费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j-1)$。

两种情况取最小值，得

$$
\textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i], \textit{dfs}(i-1,j-1))
$$

递归边界：如果 $j>i$，那么剩余的墙都可以免费刷，即 $\textit{dfs}(i,j) = 0$，否则 $\textit{dfs}(-1,j) = \infty$。

递归入口：$\textit{dfs}(n-1,0)$，即答案。

[视频讲解](https://www.bilibili.com/video/BV1Hj411D7Tr/) 第四题。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        @cache  # 记忆化搜索
        def dfs(i: int, j: int) -> int:
            if j > i:  # 剩余的墙都可以免费刷
                return 0
            if i < 0:  # 上面 if 不成立，意味着 j < 0，不符合题目要求
                return inf
            return min(dfs(i - 1, j + time[i]) + cost[i], dfs(i - 1, j - 1))
        return dfs(len(cost) - 1, 0)
```

```java [sol-Java]
class Solution {
    public int paintWalls(int[] cost, int[] time) {
        int n = cost.length;
        int[][] memo = new int[n][n * 2];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, 0, cost, time, memo);
    }

    private int dfs(int i, int j, int[] cost, int[] time, int[][] memo) {
        if (j > i) { // 剩余的墙都可以免费刷
            return 0;
        }
        if (i < 0) { // 上面 if 不成立，意味着 j < 0，不符合题目要求
            return Integer.MAX_VALUE / 2; // 防止加法溢出
        }
        int k = j + memo.length; // 加上偏移量，防止出现负数
        if (memo[i][k] != -1) { // 之前计算过
            return memo[i][k];
        }
        int res1 = dfs(i - 1, j + time[i], cost, time, memo) + cost[i];
        int res2 = dfs(i - 1, j - 1, cost, time, memo);
        return memo[i][k] = Math.min(res1, res2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int paintWalls(vector<int>& cost, vector<int>& time) {
        int n = cost.size();
        vector<vector<int>> memo(n, vector<int>(n * 2, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (j > i) { // 剩余的墙都可以免费刷
                return 0;
            }
            if (i < 0) { // 上面 if 不成立，意味着 j < 0，不符合题目要求
                return INT_MAX / 2; // 防止加法溢出
            }
            // 注意 res 是引用
            int& res = memo[i][j + n]; // 加上偏移量 n，防止出现负数
            if (res != -1) { // 之前计算过
                return res;
            }
            return res = min(dfs(dfs, i - 1, j + time[i]) + cost[i], dfs(dfs, i - 1, j - 1));
        };
        return dfs(dfs, n - 1, 0);
    }
};
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j > i { // 剩余的墙都可以免费刷
			return 0
		}
		if i < 0 { // 上面 if 不成立，意味着 j < 0，不符合题目要求
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j+n] // 加上偏移量 n，防止出现负数
		if *p != -1 { // 之前计算过
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, 0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{cost}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：转换成 0-1 背包

**前置知识**：[0-1 背包【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

看着方法一的状态转移方程，你可能会觉得：怎么感觉很像 0-1 背包呢？

没错，这题就是 0-1 背包！

根据题意，付费刷墙个数 $+$ 免费刷墙个数 $=n$。

同时，付费刷墙时间之和必须 $\ge$ 免费刷墙个数。

结合这两个式子，得到：付费刷墙时间之和 $\ge n\ -$ 付费刷墙个数。

移项，得到：「付费刷墙时间$+1$」之和 $\ge n$。（把个数拆分成 $1+1+1+\cdots$）

把 $\textit{time}[i]+1$ 看成物品体积，$\textit{cost}[i]$ 看成物品价值，问题变成：

- 从 $n$ 个物品中选择体积和**至少**为 $n$ 的物品，价值和最小是多少？

这是 0-1 背包的一种「**至少装满**」的变形。我们可以定义 $\textit{dfs}(i,j)$ 表示考虑前 $i$ 个物品，**剩余**还需要凑出 $j$ 的体积，此时的最小价值和。

和 0-1 背包一样，用选或不选思考，可以得到类似的状态转移方程：

$$
\textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j-\textit{time}[i]-1)+\textit{cost}[i], \textit{dfs}(i-1,j))
$$

递归边界：如果 $j\le 0$，不需要再选任何物品了，返回 $0$；如果 $i<0$，返回无穷大。

递归入口：$\textit{dfs}(n-1,n)$，表示体积和至少为 $n$，这正是我们要计算的。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        @cache  # 记忆化搜索
        def dfs(i: int, j: int) -> int:  # j 表示剩余需要的体积
            if j <= 0:  # 没有约束，后面什么也不用选了
                return 0
            if i < 0:  # 此时 j>0，但没有物品可选，不合法
                return inf
            return min(dfs(i - 1, j - time[i] - 1) + cost[i], dfs(i - 1, j))
        n = len(cost)
        return dfs(n - 1, n)
```

```java [sol-Java]
class Solution {
    public int paintWalls(int[] cost, int[] time) {
        int n = cost.length;
        int[][] memo = new int[n][n + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n - 1, n, cost, time, memo);
    }

    // j 表示剩余需要的体积
    private int dfs(int i, int j, int[] cost, int[] time, int[][] memo) {
        if (j <= 0) { // 没有约束，后面什么也不用选了
            return 0;
        }
        if (i < 0) { // 此时 j>0，但没有物品可选，不合法
            return Integer.MAX_VALUE / 2; // 防止加法溢出
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        int res1 = dfs(i - 1, j - time[i] - 1, cost, time, memo) + cost[i];
        int res2 = dfs(i - 1, j, cost, time, memo);
        return memo[i][j] = Math.min(res1, res2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int paintWalls(vector<int>& cost, vector<int>& time) {
        int n = cost.size();
        vector<vector<int>> memo(n, vector<int>(n + 1, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int j) -> int { // j 表示剩余需要的体积
            if (j <= 0) { // 没有约束，后面什么也不用选了
                return 0;
            }
            if (i < 0) { // 此时 j>0，但没有物品可选，不合法
                return INT_MAX / 2; // 防止加法溢出
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            return res = min(dfs(dfs, i - 1, j - time[i] - 1) + cost[i], dfs(dfs, i - 1, j));
        };
        return dfs(dfs, n - 1, n);
    }
};
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int { // j 表示剩余需要的体积
		if j <= 0 { // 没有约束，后面什么也不用选了
			return 0
		}
		if i < 0 { // 此时 j>0，但没有物品可选，不合法
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		*p = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
		return *p
	}
	return dfs(n-1, n)
}
```

然后 1:1 翻译成递推 + 空间优化，原理请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        n = len(cost)
        f = [0] + [inf] * n
        for c, t in zip(cost, time):
            for j in range(n, 0, -1):
                f[j] = min(f[j], f[max(j - t - 1, 0)] + c)
        return f[n]
```

```py [sol-Python3]
class Solution:
    def paintWalls(self, cost: List[int], time: List[int]) -> int:
        # 手写 min max 更快
        min = lambda a, b: b if b < a else a
        max = lambda a, b: b if b > a else a
        n = len(cost)
        f = [0] + [inf] * n
        for c, t in zip(cost, time):
            for j in range(n, 0, -1):
                f[j] = min(f[j], f[max(j - t - 1, 0)] + c)
        return f[n]
```

```java [sol-Java]
class Solution {
    public int paintWalls(int[] cost, int[] time) {
        int n = cost.length;
        int[] f = new int[n + 1];
        Arrays.fill(f, Integer.MAX_VALUE / 2); // 防止加法溢出
        f[0] = 0;
        for (int i = 0; i < n; i++) {
            int c = cost[i];
            int t = time[i] + 1; // 注意这里加一了
            for (int j = n; j > 0; j--) {
                f[j] = Math.min(f[j], f[Math.max(j - t, 0)] + c);
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int paintWalls(vector<int> &cost, vector<int> &time) {
        int n = cost.size();
        vector<int> f(n + 1, INT_MAX / 2); // 防止加法溢出
        f[0] = 0;
        for (int i = 0; i < n; i++) {
            int c = cost[i], t = time[i] + 1; // 注意这里加一了
            for (int j = n; j; j--) {
                f[j] = min(f[j], f[max(j - t, 0)] + c);
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func paintWalls(cost, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{cost}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

1. 如果同一面墙可以反复刷呢？
2. 把 $\textit{cost}$ 去掉，不考虑开销，只考虑时间，如果有 $x$ 位付费油漆匠和 $y$ 位免费油漆匠，它们可以同时工作，刷完所有墙的最小耗时是多少？

欢迎在评论区发表你的思路/代码。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
