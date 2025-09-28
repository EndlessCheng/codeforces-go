## 一、寻找子问题

我们要从 $0$ 爬到 $n$。

考虑最后一步爬了多少个台阶：

- 最后一步爬 $1$ 个台阶，问题变成从 $0$ 爬到 $n-1$ 的最小总成本。
- 最后一步爬 $2$ 个台阶，问题变成从 $0$ 爬到 $n-2$ 的最小总成本。
- 最后一步爬 $3$ 个台阶，问题变成从 $0$ 爬到 $n-3$ 的最小总成本。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 注：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。

## 二、状态定义与状态转移方程

根据上面的讨论，定义 $\textit{dfs}(i)$，表示从 $0$ 爬到 $i$ 的最小总成本。

枚举最后一步从 $j$ 爬到 $i$，问题变成从 $0$ 爬到 $j$ 的最小总成本，即 $\textit{dfs}(j)$。加上从 $j$ 到 $i$ 的成本 $(i-j)^2 + \textit{costs}[i]$，更新 $\textit{dfs}(i)$ 的最小值：

$$
\textit{dfs}(i) = \min_{j=\max(i-3,0)}^{i-1}\textit{dfs}(j) + (i-j)^2 + \textit{costs}[i]
$$

**递归边界**：$\textit{dfs}(0)=0$。从 $0$ 爬到 $0$ 无需移动，成本为 $0$。

**递归入口**：$\textit{dfs}(n)$，这是原问题（从 $0$ 爬到 $n$ 的最小总成本），也是答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

⚠**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。但本题成本是正数，$i>0$ 时 $\textit{dfs}(i) > 0$，所以可以把 $\textit{memo}[i]$ 初始化成 $0$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

⚠**注意**：题目说 $\textit{costs}$ 的下标从 $1$ 开始，但传入的 $\textit{costs}$ 的下标是从 $0$ 开始的。访问数组的时候下标要减一。

```py [sol-Python3]
class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int) -> int:
            if i == 0:
                return 0
            return min(dfs(j) + (i - j) * (i - j) for j in range(max(i - 3, 0), i)) + costs[i - 1]
        return dfs(n)
```

```java [sol-Java]
class Solution {
    public int climbStairs(int n, int[] costs) {
        int[] memo = new int[n + 1];
        return dfs(n, costs, memo);
    }

    private int dfs(int i, int[] costs, int[] memo) {
        if (i == 0) {
            return 0;
        }
        if (memo[i] != 0) { // 之前计算过
            return memo[i];
        }
        int res = Integer.MAX_VALUE;
        for (int j = Math.max(i - 3, 0); j < i; j++) {
            res = Math.min(res, dfs(j, costs, memo) + (i - j) * (i - j));
        }
        res += costs[i - 1];
        return memo[i] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int climbStairs(int n, vector<int>& costs) {
        vector<int> memo(n + 1);

        // lambda 递归函数
        auto dfs = [&](this auto&& dfs, int i) -> int {
            if (i == 0) {
                return 0;
            }
            int& res = memo[i]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = INT_MAX;
            for (int j = max(i - 3, 0); j < i; j++) {
                res = min(res, dfs(j) + (i - j) * (i - j));
            }
            res += costs[i - 1];
            return res;
        };

        return dfs(n);
    }
};
```

```go [sol-Go]
func climbStairs(n int, costs []int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != 0 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		for j := max(i-3, 0); j < i; j++ {
			res = min(res, dfs(j)+(i-j)*(i-j))
		}
		res += costs[i-1]
		*p = res // 记忆化
		return res
	}
	return dfs(n)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nK)$，其中 $K=3$ 是最多可以跳的台阶数。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(K)$，所以总的时间复杂度为 $\mathcal{O}(nK)$。
- 空间复杂度：$\mathcal{O}(n)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i]$ 的定义和 $\textit{dfs}(i)$ 的定义是完全一样的，都表示从 $0$ 爬到 $i$ 的最小总成本。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i] = \min_{j=\max(i-3,0)}^{i-1}f[j] + (i-j)^2 + \textit{costs}[i]
$$

初始值 $f[0]=0$，翻译自递归边界 $\textit{dfs}(0)=0$。

答案为 $f[n]$，翻译自递归入口 $\textit{dfs}(n)$。

```py [sol-Python3]
class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        f = [0] * (n + 1)
        for i in range(1, n + 1):
            f[i] = min(f[j] + (i - j) * (i - j) for j in range(max(i - 3, 0), i)) + costs[i - 1]
        return f[n]
```

```java [sol-Java]
class Solution {
    public int climbStairs(int n, int[] costs) {
        int[] f = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            int res = Integer.MAX_VALUE;
            for (int j = Math.max(i - 3, 0); j < i; j++) {
                res = Math.min(res, f[j] + (i - j) * (i - j));
            }
            f[i] = res + costs[i - 1];
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int climbStairs(int n, vector<int>& costs) {
        vector<int> f(n + 1);
        for (int i = 1; i <= n; i++) {
            int res = INT_MAX;
            for (int j = max(i - 3, 0); j < i; j++) {
                res = min(res, f[j] + (i - j) * (i - j));
            }
            f[i] = res + costs[i - 1];
        }
        return f[n];
    }
};
```

```go [sol-Go]
func climbStairs(n int, costs []int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res := math.MaxInt
		for j := max(i-3, 0); j < i; j++ {
			res = min(res, f[j]+(i-j)*(i-j))
		}
		f[i] = res + costs[i-1]
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nK)$，其中 $K=3$ 是最多可以跳的台阶数。
- 空间复杂度：$\mathcal{O}(n)$。

## 五、空间优化

类似 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)，用三个变量滚动计算。原理见 [我的题解](https://leetcode.cn/problems/climbing-stairs/solutions/2560716/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-7zm1/)。

```py [sol-Python3]
class Solution:
    def climbStairs(self, _, costs: List[int]) -> int:
        f0 = f1 = f2 = 0
        for c in costs:
            f0, f1, f2 = f1, f2, min(f0 + 9, f1 + 4, f2 + 1) + c
        return f2
```

```py [sol-Python3 手写 min]
class Solution:
    def climbStairs(self, _, costs: List[int]) -> int:
        f0 = f1 = f2 = 0
        for c in costs:
            mn = f0 + 9
            if (t := f1 + 4) < mn: mn = t
            if (t := f2 + 1) < mn: mn = t
            f0, f1, f2 = f1, f2, mn + c
        return f2
```

```java [sol-Java]
class Solution {
    public int climbStairs(int n, int[] costs) {
        int f0 = 0, f1 = 0, f2 = 0;
        for (int c : costs) {
            int newF = Math.min(Math.min(f0 + 9, f1 + 4), f2 + 1) + c;
            f0 = f1;
            f1 = f2;
            f2 = newF;
        }
        return f2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int climbStairs(int, vector<int>& costs) {
        int f0 = 0, f1 = 0, f2 = 0;
        for (int c : costs) {
            int new_f = min({f0 + 9, f1 + 4, f2 + 1}) + c;
            f0 = f1;
            f1 = f2;
            f2 = new_f;
        }
        return f2;
    }
};
```

```go [sol-Go]
func climbStairs(_ int, costs []int) int {
	var f0, f1, f2 int
	for _, c := range costs {
		f0, f1, f2 = f1, f2, min(f0+9, f1+4, f2+1)+c
	}
	return f2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nK)$，其中 $K=3$ 是最多可以跳的台阶数。
- 空间复杂度：$\mathcal{O}(K)$。

本题有两个变形：

1. 去掉 $K=3$ 的约束（可以爬任意多级台阶），这可以用斜率优化，时间复杂度 $\mathcal{O}(n)$。
2. 额外传入一个参数 $K$，这可以用 [李超线段树](https://oi-wiki.org/ds/li-chao-tree/)，时间复杂度 $\mathcal{O}(n\log^2 n)$。

## 专题训练

1. 动态规划题单的「**§1.1 爬楼梯**」。
2. 动态规划题单的「**§11.7 斜率优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
