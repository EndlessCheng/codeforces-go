## 题意

从 $\textit{nums}_1$ 和 $\textit{nums}_2$ 中各选一个长度恰好为 $k$ 的子序列，计算 [1458. 两个子序列的最大点积](https://leetcode.cn/problems/max-dot-product-of-two-subsequences/)。

## 思路

根据数据范围，我们可以写一个 $\mathcal{O}(knm)$ 的算法，额外用一个参数 $k$ 表示剩余需要选的下标对数。

定义 $\textit{dfs}(k,i,j)$ 表示从 $\textit{nums}_1$ 的 $[0,i]$ 和 $\textit{nums}_2$ 的 $[0,j]$ 中各选一个长度恰好为 $k$ 的子序列，这两个子序列的最大点积。

用「选或不选」思考：

- 不选 $\textit{nums}_1[i]$，问题变成从 $\textit{nums}_1$ 的 $[0,i-1]$ 和 $\textit{nums}_2$ 的 $[0,j]$ 中各选一个长度恰好为 $k$ 的子序列，这两个子序列的最大点积，即 $\textit{dfs}(k,i-1,j)$。
- 不选 $\textit{nums}_2[j]$，问题变成从 $\textit{nums}_1$ 的 $[0,i]$ 和 $\textit{nums}_2$ 的 $[0,j-1]$ 中各选一个长度恰好为 $k$ 的子序列，这两个子序列的最大点积，即 $\textit{dfs}(k,i,j-1)$。
- 选 $\textit{nums}_1[i]$ 和 $\textit{nums}_2[j]$，问题变成从 $\textit{nums}_1$ 的 $[0,i-1]$ 和 $\textit{nums}_2$ 的 $[0,j-1]$ 中各选一个长度恰好为 $k-1$ 的子序列，这两个子序列的最大点积，再加上 $\textit{nums}_1[i]\cdot\textit{nums}_2[j]$，即 $\textit{dfs}(k-1,i-1,j-1) + \textit{nums}_1[i]\cdot\textit{nums}_2[j]$。

三种情况取最大值，得

$$
\textit{dfs}(k,i,j) = \max(\textit{dfs}(k,i-1,j),\textit{dfs}(k,i,j-1),\textit{dfs}(k-1,i-1,j-1) + \textit{nums}_1[i]\cdot\textit{nums}_2[j])
$$

**递归边界**：

- $\textit{dfs}(0,i,j)=0$。没有数对要选，点积为 $0$。
- 如果 $i+1<k$ 或者 $j+1<k$，那么剩余元素不足 $k$ 个，无法满足要求，返回 $-\infty$。这样计算 $\max$ 的时候，一定会取到合法的情况。

**递归入口**：$\textit{dfs}(k,n-1,m-1)$，这是原问题，也是答案。

## 递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

⚠**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(k,i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。但本题有负数，可以初始化成 $-\infty$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1idFoz3Efi/?t=18m12s)，欢迎点赞关注~

## 写法一：记忆化搜索

```py [sol-Python3]
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(k: int, i: int, j: int) -> int:
            if k == 0:  # 选完了
                return 0
            if i + 1 < k or j + 1 < k:  # 剩余元素不足 k 个
                return -inf  # 下面计算 max 不会取到 -inf
            res1 = dfs(k, i - 1, j)  # 不选 nums1[i]
            res2 = dfs(k, i, j - 1)  # 不选 nums2[j]
            res3 = dfs(k - 1, i - 1, j - 1) + nums1[i] * nums2[j]  # 选 nums1[i] 和 nums2[j]
            return max(res1, res2, res3)

        return dfs(k, len(nums1) - 1, len(nums2) - 1)
```

```java [sol-Java]
class Solution {
    public long maxScore(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        int m = nums2.length;

        long[][][] memo = new long[k + 1][n][m];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, Long.MIN_VALUE);
            }
        }

        return dfs(k, n - 1, m - 1, nums1, nums2, memo);
    }

    private long dfs(int k, int i, int j, int[] nums1, int[] nums2, long[][][] memo) {
        if (k == 0) { // 选完了
            return 0;
        }
        if (i + 1 < k || j + 1 < k) { // 剩余元素不足 k 个
            return Long.MIN_VALUE; // 下面计算 max 不会取到 MIN_VALUE
        }
        if (memo[k][i][j] != Long.MIN_VALUE) { // 之前计算过
            return memo[k][i][j];
        }

        long res1 = dfs(k, i - 1, j, nums1, nums2, memo); // 不选 nums1[i]
        long res2 = dfs(k, i, j - 1, nums1, nums2, memo); // 不选 nums2[j]
        long res3 = dfs(k - 1, i - 1, j - 1, nums1, nums2, memo) + (long) nums1[i] * nums2[j]; // 选 nums1[i] 和 nums2[j]
        return memo[k][i][j] = Math.max(Math.max(res1, res2), res3); // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(vector<int>& nums1, vector<int>& nums2, int k) {
        int n = nums1.size(), m = nums2.size();
        vector memo(k + 1, vector(n, vector<long long>(m, LLONG_MIN)));

        auto dfs = [&](this auto&& dfs, int k, int i, int j) -> long long {
            if (k == 0) { // 选完了
                return 0;
            }
            if (i + 1 < k || j + 1 < k) { // 剩余元素不足 k 个
                return LLONG_MIN; // 下面计算 max 不会取到 LLONG_MIN
            }

            long long& res = memo[k][i][j]; // 注意这里是引用
            if (res != LLONG_MIN) { // 之前计算过
                return res;
            }

            long long res1 = dfs(k, i - 1, j); // 不选 nums1[i]
            long long res2 = dfs(k, i, j - 1); // 不选 nums2[j]
            long long res3 = dfs(k - 1, i - 1, j - 1) + 1LL * nums1[i] * nums2[j]; // 选 nums1[i] 和 nums2[j]
            res = max(max(res1, res2), res3);
            return res;
        };

        return dfs(k, n - 1, m - 1);
    }
};
```

```go [sol-Go]
func maxScore(nums1, nums2 []int, k int) int64 {
	n, m := len(nums1), len(nums2)
	memo := make([][][]int, k+1)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, m)
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt // MinInt 表示没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(k, i, j int) int {
		if k == 0 { // 选完了
			return 0
		}
		if i+1 < k || j+1 < k { // 剩余元素不足 k 个
			return math.MinInt // 下面计算 max 不会取到 MinInt
		}

		p := &memo[k][i][j]
		if *p != math.MinInt { // 之前计算过
			return *p
		}

		res1 := dfs(k, i-1, j)                         // 不选 nums1[i]
		res2 := dfs(k, i, j-1)                         // 不选 nums2[j]
		res3 := dfs(k-1, i-1, j-1) + nums1[i]*nums2[j] // 选 nums1[i] 和 nums2[j]
		res := max(res1, res2, res3)

		*p = res // 记忆化
		return res
	}
	return int64(dfs(k, n-1, m-1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(knm)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(knm)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(knm)$。
- 空间复杂度：$\mathcal{O}(knm)$。保存多少状态，就需要多少空间。

## 写法二：递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[k][i+1][j+1]$ 的定义和 $\textit{dfs}(k,i,j)$ 的定义是一样的，都表示从 $\textit{nums}_1$ 的 $[0,i]$ 和 $\textit{nums}_2$ 的 $[0,j]$ 中各选一个长度恰好为 $k$ 的子序列，这两个子序列的最大点积。这里 $+1$ 是为了把 $\textit{dfs}(k,-1,j)$ 和 $\textit{dfs}(k,i,-1)$ 这种状态也翻译过来，这样我们可以把 $f[k][0][j]$ 和 $f[k][i][0]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[k][i+1][j+1] = \max(f[k][i][j+1],f[k][i+1][j],f[k-1][i][j] + \textit{nums}_1[i]\cdot\textit{nums}_2[j])
$$

由于 $i+1<k$ 和 $j+1<k$ 均不合法，所以 $i$ 和 $j$ 可以从 $k-1$ 开始循环。

**初始值**：$f[0][i][j] = 0$，其余为 $-\infty$。

**答案**：$f[k][n][m]$。

```py [sol-Python3]
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], K: int) -> int:
        n, m = len(nums1), len(nums2)
        f = [[[-inf] * (m + 1) for _ in range(n + 1)] for _ in range(K + 1)]
        f[0] = [[0] * (m + 1) for _ in range(n + 1)]
        for k in range(1, K + 1):
            for i in range(k - 1, n):
                for j in range(k - 1, m):
                    f[k][i + 1][j + 1] = max(f[k][i][j + 1], f[k][i + 1][j], f[k - 1][i][j] + nums1[i] * nums2[j])
        return f[K][n][m]
```

```java [sol-Java]
class Solution {
    public long maxScore(int[] nums1, int[] nums2, int K) {
        int n = nums1.length;
        int m = nums2.length;
        long[][][] f = new long[K + 1][n + 1][m + 1];
        for (int k = 1; k <= K; k++) {
            for (long[] row : f[k]) {
                Arrays.fill(row, Long.MIN_VALUE);
            }
        }
        for (int k = 1; k <= K; k++) {
            for (int i = k - 1; i < n; i++) {
                for (int j = k - 1; j < m; j++) {
                    f[k][i + 1][j + 1] = Math.max(Math.max(f[k][i][j + 1], f[k][i + 1][j]),
                            f[k - 1][i][j] + (long) nums1[i] * nums2[j]);
                }
            }
        }
        return f[K][n][m];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxScore(vector<int>& nums1, vector<int>& nums2, int K) {
        int n = nums1.size(), m = nums2.size();
        vector f(K + 1, vector(n + 1, vector<long long>(m + 1, LLONG_MIN)));
        for (auto& row : f[0]) {
            ranges::fill(row, 0);
        }
        for (int k = 1; k <= K; k++) {
            for (int i = k - 1; i < n; i++) {
                for (int j = k - 1; j < m; j++) {
                    f[k][i + 1][j + 1] = max(max(f[k][i][j + 1], f[k][i + 1][j]), 
                                             f[k - 1][i][j] + 1LL * nums1[i] * nums2[j]);
                }
            }
        }
        return f[K][n][m];
    }
};
```

```go [sol-Go]
func maxScore(nums1, nums2 []int, K int) int64 {
	n, m := len(nums1), len(nums2)
	f := make([][][]int, K+1)
	for k := range f {
		f[k] = make([][]int, n+1)
		for i := range f[k] {
			f[k][i] = make([]int, m+1)
			if k > 0 {
				for j := range f[k][i] {
					f[k][i][j] = math.MinInt
				}
			}
		}
	}
	for k := 1; k <= K; k++ {
		for i := k - 1; i < n; i++ {
			for j := k - 1; j < m; j++ {
				f[k][i+1][j+1] = max(f[k][i][j+1], f[k][i+1][j], f[k-1][i][j]+nums1[i]*nums2[j])
			}
		}
	}
	return int64(f[K][n][m])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(knm)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(knm)$。

**注**：利用滚动数组，可以把空间复杂度优化到 $\mathcal{O}(nm)$。

## 专题训练

见下面动态规划题单的「**§4.1 最长公共子序列（LCS）**」。

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
