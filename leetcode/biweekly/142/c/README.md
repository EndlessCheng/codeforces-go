## 一、寻找子问题

假设一开始（第 $0$ 天）我们在城市 $1$，分类讨论：

- 留在当前城市。接下来需要解决的问题为：第 $1$ 天到第 $k-1$ 天，从城市 $1$ 开始旅游，可以获得的最多点数。
- 前往另外一座城市。枚举前往城市 $0,1,2,\ldots,n-1$，假如前往城市 $0$，那么接下来需要解决的问题为：第 $1$ 天到第 $k-1$ 天，从城市 $0$ 开始旅游，可以获得的最多点数。⚠**注意**：题目保证 $\textit{travelScore}[i][i]=0$，所以前往当前城市也可以，这一定不如留在当前城市优。如果题目不保证这一点，那么必须枚举不等于当前城市的其他城市。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 为了方便遍历 $\textit{travelScore}$，本文使用从前往后递归，从后往前递推的写法。反过来写也是可以的。

## 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前在第 $i$ 天。
- $j$：当前在城市 $j$。

因此，定义状态为 $\textit{dfs}(i,j)$，表示第 $i$ 天到第 $k-1$ 天，从城市 $j$ 开始旅游，可以获得的最多点数。

分类讨论：

- 留在当前城市。接下来需要解决的问题为：第 $i+1$ 天到第 $k-1$ 天，从城市 $j$ 开始旅游，可以获得的最多点数，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1,j) + \textit{stayScore}[i][j]$。
- 前往另外一座城市。枚举前往城市 $d=0,1,2,\ldots,n-1$，接下来需要解决的问题为：第 $i+1$ 天到第 $k-1$ 天，从城市 $d$ 开始旅游，可以获得的最多点数，即 $\textit{dfs}(i,j) = \textit{dfs}(i+1,d) + \textit{travelScore}[j][d]$。

所有情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i+1,j) + \textit{stayScore}[i][j], \max\limits_{d=0}^{n-1} \{\textit{dfs}(i+1,d) + \textit{travelScore}[j][d]\})
$$

**递归边界**：$\textit{dfs}(k,j)=0$。$k$ 天的旅程结束了。

**递归入口**：$\textit{dfs}(0,j)$。取最大值 $\max\limits_{j=0}^{n-1}\textit{dfs}(0,j)$ 作为答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。本题数据范围保证记忆化的值不等于 $0$，所以可以初始化成 $0$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

[本题视频讲解](https://www.bilibili.com/video/BV13J1MYwEGM/?t=9m44s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxScore(self, n: int, k: int, stayScore: List[List[int]], travelScore: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, j: int) -> int:
            if i == k:
                return 0
            res1 = dfs(i + 1, j) + stayScore[i][j]  # 留在当前城市 j
            # 注意题目保证 travelScore[j][j]=0，这一定不如留在当前城市优
            res2 = max(dfs(i + 1, d) + s for d, s in enumerate(travelScore[j]))  # 前往另外一座城市 d
            return max(res1, res2)
        return max(dfs(0, j) for j in range(n))  # 枚举城市 j 作为起点
```

```java [sol-Java]
class Solution {
    public int maxScore(int n, int k, int[][] stayScore, int[][] travelScore) {
        int[][] memo = new int[k][n];
        int ans = 0;
        // 枚举城市 j 作为起点
        for (int j = 0; j < n; j++) {
            ans = Math.max(ans, dfs(0, j, stayScore, travelScore, k, memo));
        }
        return ans;
    }

    private int dfs(int i, int j, int[][] stayScore, int[][] travelScore, int k, int[][] memo) {
        if (i == k) {
            return 0;
        }
        // 之前计算过
        if (memo[i][j] > 0) {
            return memo[i][j];
        }
        // 留在当前城市 j
        int res = dfs(i + 1, j, stayScore, travelScore, k, memo) + stayScore[i][j];
        // 前往另外一座城市 d
        for (int d = 0; d < travelScore[j].length; d++) {
            // 注意题目保证 travelScore[j][j]=0，这一定不如留在当前城市优
            res = Math.max(res, dfs(i + 1, d, stayScore, travelScore, k, memo) + travelScore[j][d]);
        }
        // 记忆化
        return memo[i][j] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(int n, int k, vector<vector<int>>& stayScore, vector<vector<int>>& travelScore) {
        vector<vector<int>> memo(k, vector<int>(n));
        auto dfs = [&](auto&& dfs, int i, int j) -> int {
            if (i == k) {
                return 0;
            }
            int& res = memo[i][j]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = dfs(dfs, i + 1, j) + stayScore[i][j]; // 留在当前城市 j
            for (int d = 0; d < travelScore[j].size(); d++) {
                // 注意题目保证 travelScore[j][j]=0，这一定不如留在当前城市优
                res = max(res, dfs(dfs, i + 1, d) + travelScore[j][d]); // 前往另外一座城市 d
            }
            return res;
        };

        int ans = 0;
        for (int j = 0; j < n; j++) { // 枚举城市 j 作为起点
            ans = max(ans, dfs(dfs, 0, j));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxScore(n, k int, stayScore, travelScore [][]int) (ans int) {
	memo := make([][]int, k)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == k {
			return 0
		}
		p := &memo[i][j]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := dfs(i+1, j) + stayScore[i][j] // 留在当前城市 j
		for d, s := range travelScore[j] {
			// 注意题目保证 travelScore[j][j]=0，这一定不如留在当前城市优
			res = max(res, dfs(i+1, d)+s) // 前往另外一座城市 d
		}
		*p = res // 记忆化
		return res
	}
	for j := range n { // 枚举城市 j 作为起点
		ans = max(ans, dfs(0, j))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kn^2)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(kn)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以总的时间复杂度为 $\mathcal{O}(kn^2)$。
- 空间复杂度：$\mathcal{O}(kn)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示第 $i$ 天到第 $k-1$ 天，从城市 $j$ 开始旅游，可以获得的最多点数。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \max(f[i+1][j] + \textit{stayScore}[i][j], \max\limits_{d=0}^{n-1} \{f[i+1][d] + \textit{travelScore}[j][d]\})
$$

初始值 $f[k][j]=0$，翻译自递归边界 $\textit{dfs}(k,j)=0$。

答案为 $\max\limits_{j=0}^{n-1}f[0][j]$，翻译自递归入口 $\textit{dfs}(0,j)$。

#### 答疑

**问**：可以正序枚举吗？

**答**：也可以，但要枚举来到 $j$ 的城市是哪个。（或者使用刷表法。）

```py [sol-Python3]
class Solution:
    def maxScore(self, n: int, k: int, stayScore: List[List[int]], travelScore: List[List[int]]) -> int:
        f = [[0] * n for _ in range(k + 1)]
        for i in range(k - 1, -1, -1):
            for j, ss in enumerate(stayScore[i]):
                f[i][j] = max(f[i + 1][j] + ss,
                              max(fd + ts for fd, ts in zip(f[i + 1], travelScore[j])))
        return max(f[0])
```

```java [sol-Java]
class Solution {
    public int maxScore(int n, int k, int[][] stayScore, int[][] travelScore) {
        int[][] f = new int[k + 1][n];
        for (int i = k - 1; i >= 0; i--) {
            for (int j = 0; j < n; j++) {
                f[i][j] = f[i + 1][j] + stayScore[i][j];
                for (int d = 0; d < n; d++) {
                    f[i][j] = Math.max(f[i][j], f[i + 1][d] + travelScore[j][d]);
                }
            }
        }
        int ans = 0;
        for (int x : f[0]) {
            ans = Math.max(ans, x);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxScore(int n, int k, vector<vector<int>>& stayScore, vector<vector<int>>& travelScore) {
        vector<vector<int>> f(k + 1, vector<int>(n));
        for (int i = k - 1; i >= 0; i--) {
            for (int j = 0; j < n; j++) {
                f[i][j] = f[i + 1][j] + stayScore[i][j];
                for (int d = 0; d < n; d++) {
                    f[i][j] = max(f[i][j], f[i + 1][d] + travelScore[j][d]);
                }
            }
        }
        return ranges::max(f[0]);
    }
};
```

```go [sol-Go]
func maxScore(n, k int, stayScore, travelScore [][]int) int {
	f := make([][]int, k+1)
	f[k] = make([]int, n)
	for i, row := range slices.Backward(stayScore) {
		f[i] = make([]int, n)
		for j, s := range row {
			f[i][j] = f[i+1][j] + s // s=stayScore[i][j]
			for d, ts := range travelScore[j] {
				f[i][j] = max(f[i][j], f[i+1][d]+ts)
			}
		}
	}
	return slices.Max(f[0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kn^2)$。
- 空间复杂度：$\mathcal{O}(kn)$。

注：也可以把 $\textit{stayScore}$ 作为 $f$ 数组，这样可以做到 $\mathcal{O}(1)$ 额外空间。

更多相似题目，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§7.5 多维 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
