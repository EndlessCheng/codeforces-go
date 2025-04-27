如果你没做过状压 DP，请先完成 [526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)，并阅读我的题解 [教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

**关键思路**：把拓扑序理解为先修课关系，在学习课程 $j$ 之前，$j$ 的所有先修课（直接前驱）必须全部学完。

定义 $\textit{dfs}(S)$ 表示在已学课程集合为 $S$ 的情况下，**剩余课程**可以获得的最大利润。

考虑下一门课程学哪个：

- 枚举学习 $j=0,1,2,\ldots,n-1$，要求满足 $j\notin S$ 且 $\textit{pre}[j] \subseteq S$。其中 $\textit{pre}[j]$ 表示 $j$ 的先修课集合。
- 要解决的问题变成：在已学课程集合为 $S \cup \{j\}$ 的情况下，剩余课程可以获得的最大利润，即 $\textit{dfs}(S \cup \{j\})$。

取最大值，得

$$
\textit{dfs}(S) = \max_{j=0}^{n-1} \textit{dfs}(S \cup \{j\}) + \textit{score}[j] \cdot (|S|+1)
$$

其中 $j\notin S$ 且 $\textit{pre}[j] \subseteq S$，$|S|$ 表示集合 $S$ 的大小。

递归边界：$\textit{dfs}(U)=0$，其中全集 $U=\{0,1,2,\ldots,n-1\}$。递归到 $S=U$ 的状态，表示所有课程全部学完，剩余课程可以获得的最大利润为 $0$。

递归入口：$\textit{dfs}(\varnothing)$，也就是答案。其中 $\varnothing$ 表示空集，因为一开始什么课程也没有学。

**代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看** [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

## 写法一：记忆化搜索

关于记忆化搜索的原理，请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def maxProfit(self, n: int, edges: List[List[int]], score: List[int]) -> int:
        # 记录每个节点的先修课（直接前驱）
        pre = [0] * n
        for x, y in edges:
            pre[y] |= 1 << x

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(s: int) -> int:
            res = 0
            i = s.bit_count() + 1  # 已学课程数加一
            # 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for j, p in enumerate(pre):
                if (s >> j & 1) == 0 and (s | p) == s:
                    r = dfs(s | 1 << j) + score[j] * i
                    if r > res:  # 手写 max
                        res = r
            return res

        return dfs(0)
```

```java [sol-Java]
class Solution {
    public int maxProfit(int n, int[][] edges, int[] score) {
        // 记录每个节点的先修课（直接前驱）
        int[] pre = new int[n];
        for (int[] e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        int[] memo = new int[1 << n];
        return dfs(0, pre, score, memo);
    }

    private int dfs(int s, int[] pre, int[] score, int[] memo) {
        if (memo[s] > 0) { // 之前计算过
            return memo[s];
        }
        int res = 0;
        int i = Integer.bitCount(s); // 已学课程数
        // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
        for (int j = 0; j < pre.length; j++) {
            if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                res = Math.max(res, dfs(s | 1 << j, pre, score, memo) + score[j] * (i + 1));
            }
        }
        return memo[s] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProfit(int n, vector<vector<int>>& edges, vector<int>& score) {
        // 记录每个节点的先修课（直接前驱）
        vector<int> pre(n);
        for (auto& e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        vector<int> memo(1 << n);
        auto dfs = [&](this auto&& dfs, int s) -> int {
            int& res = memo[s]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            int i = popcount((uint32_t) s); // 已学课程数
            // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                    res = max(res, dfs(s | 1 << j) + score[j] * (i + 1));
                }
            }
            return res;
        };
        return dfs(0);
    }
};
```

```go [sol-Go]
func maxProfit(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	memo := make([]int, 1<<n)
	var dfs func(s int) int
	dfs = func(s int) (res int) {
		m := &memo[s]
		if *m > 0 { // 之前计算过
			return *m
		}
		defer func() { *m = res }() // 记忆化
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for j, p := range pre {
			if s>>j&1 == 0 && s|p == s {
				res = max(res, dfs(s|1<<j)+score[j]*(i+1))
			}
		}
		return
	}
	return dfs(0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n2^n)$，其中 $m$ 是 $\textit{edges}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以记忆化搜索的时间复杂度为 $\mathcal{O}(n2^n)$。
- 空间复杂度：$\mathcal{O}(2^n)$。保存多少状态，就需要多少空间。

## 写法二：记忆化搜索 1:1 翻译成递推（倒序）

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[S]$ 的定义和 $\textit{dfs}(S)$ 的定义是完全一样的，都表示在已学课程集合为 $S$ 的情况下，**剩余课程**可以获得的最大利润。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[S] = \max_{j=0}^{n-1} f[S \cup \{j\}] + \textit{score}[j] \cdot (|S|+1)
$$

初始值 $f[U]=0$，翻译自递归边界 $\textit{dfs}(U)=0$。

答案为 $f[\varnothing]$，翻译自递归入口 $\textit{dfs}(\varnothing)$。

⚠**注意**：下面的写法超时，请看写法三。

```py [sol-Python3]
# 超时了！请看写法三！
class Solution:
    def maxProfit(self, n: int, edges: List[List[int]], score: List[int]) -> int:
        # 记录每个节点的先修课（直接前驱）
        pre = [0] * n
        for x, y in edges:
            pre[y] |= 1 << x

        f = [0] * (1 << n)
        for s in range((1 << n) - 2, -1, -1):
            res = 0
            i = s.bit_count() + 1  # 已学课程数加一
            # 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for j, p in enumerate(pre):
                if (s >> j & 1) == 0 and (s | p) == s:
                    r = f[s | 1 << j] + score[j] * i
                    if r > res:  # 手写 max
                        res = r
            f[s] = res
        return f[0]
```

```java [sol-Java]
// 超时了！请看写法三！
class Solution {
    public int maxProfit(int n, int[][] edges, int[] score) {
        // 记录每个节点的先修课（直接前驱）
        int[] pre = new int[n];
        for (int[] e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        int u = 1 << n;
        int[] f = new int[u];

        for (int s = u - 2; s >= 0; s--) {
            int i = Integer.bitCount(s); // 已学课程数
            // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                    f[s] = Math.max(f[s], f[s | 1 << j] + score[j] * (i + 1));
                }
            }
        }
        return f[0];
    }
}
```

```cpp [sol-C++]
// 超时了！请看写法三！
class Solution {
public:
    int maxProfit(int n, vector<vector<int>>& edges, vector<int>& score) {
        // 记录每个节点的先修课（直接前驱）
        vector<int> pre(n);
        for (auto& e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        int u = 1 << n;
        vector<int> f(u);

        for (int s = u - 2; s >= 0; s--) {
            int i = popcount((uint32_t) s); // 已学课程数
            // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                    f[s] = max(f[s], f[s | 1 << j] + score[j] * (i + 1));
                }
            }
        }
        return f[0];
    }
};
```

```go [sol-Go]
// 超时了！请看写法三！
func maxProfit(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	u := 1 << n
	f := make([]int, u)

	for s := u - 2; s >= 0; s-- {
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for j, p := range pre {
			if s>>j&1 == 0 && s|p == s {
				f[s] = max(f[s], f[s|1<<j]+score[j]*(i+1))
			}
		}
	}
	return f[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n2^n)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 写法三：刷表法（正序）

在动态规划中，用转移来源更新当前状态叫**查表法**（写法二），用当前状态更新其他状态叫**刷表法**（写法三）。

这样写的好处是，如果一个状态从未被更新过，说明这个状态不合法（比如已经学完后面的课程，但前面的课程还没学），无需执行内层循环。

```py [sol-Python3]
class Solution:
    def maxProfit(self, n: int, edges: List[List[int]], score: List[int]) -> int:
        # 记录每个节点的先修课（直接前驱）
        pre = [0] * n
        for x, y in edges:
            pre[y] |= 1 << x

        f = [-1] * (1 << n)
        f[0] = 0
        for s, fs in enumerate(f):
            if fs < 0:  # 不合法状态，比如已经学完后面的课程，但前面的课程还没学
                continue
            i = s.bit_count() + 1  # 已学课程数加一
            # 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for j, p in enumerate(pre):
                if (s >> j & 1) == 0 and (s | p) == s:
                    new_s = s | 1 << j
                    r = f[s] + score[j] * i
                    if r > f[new_s]:  # 手写 max
                        f[new_s] = r
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int maxProfit(int n, int[][] edges, int[] score) {
        // 记录每个节点的先修课（直接前驱）
        int[] pre = new int[n];
        for (int[] e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        int u = 1 << n;
        int[] f = new int[u];
        Arrays.fill(f, -1);
        f[0] = 0;

        for (int s = 0; s < u - 1; s++) {
            if (f[s] < 0) { // 不合法状态，比如已经学完后面的课程，但前面的课程还没学
                continue;
            }
            int i = Integer.bitCount(s); // 已学课程数
            // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                    int newS = s | 1 << j;
                    f[newS] = Math.max(f[newS], f[s] + score[j] * (i + 1));
                }
            }
        }
        return f[u - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProfit(int n, vector<vector<int>>& edges, vector<int>& score) {
        // 记录每个节点的先修课（直接前驱）
        vector<int> pre(n);
        for (auto& e : edges) {
            pre[e[1]] |= 1 << e[0];
        }

        uint32_t u = 1 << n;
        vector<int> f(u, -1);
        f[0] = 0;

        for (uint32_t s = 0; s < u - 1; s++) {
            if (f[s] < 0) { // 不合法状态，比如已经学完后面的课程，但前面的课程还没学
                continue;
            }
            int i = popcount(s); // 已学课程数
            // 枚举还没学过的课程 j，且 j 的所有先修课都学完了
            for (int j = 0; j < n; j++) {
                if ((s >> j & 1) == 0 && (s | pre[j]) == s) {
                    int new_s = s | 1 << j;
                    f[new_s] = max(f[new_s], f[s] + score[j] * (i + 1));
                }
            }
        }
        return f[u - 1];
    }
};
```

```go [sol-Go]
func maxProfit(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	u := 1 << n
	f := make([]int, u)
	for s := 1; s < u; s++ {
		f[s] = -1
	}

	for s, fs := range f {
		if fs < 0 { // 不合法状态，比如已经学完后面的课程，但前面的课程还没学
			continue
		}
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for j, p := range pre {
			if s>>j&1 == 0 && s|p == s {
				newS := s | 1<<j
				f[newS] = max(f[newS], fs+score[j]*(i+1))
			}
		}
	}
	return f[u-1]
}
```

```go [sol-Go 更快写法]
func maxProfit(n int, edges [][]int, score []int) int {
	// 记录每个节点的先修课（直接前驱）
	pre := make([]int, n)
	for _, e := range edges {
		pre[e[1]] |= 1 << e[0]
	}

	u := 1 << n
	f := make([]int, u)
	for s := 1; s < u; s++ {
		f[s] = -1
	}

	for s, fs := range f {
		if fs < 0 { // 不合法状态，比如已经学完后面的课程，但前面的课程还没学
			continue
		}
		i := bits.OnesCount(uint(s)) // 已学课程数
		// 枚举还没学过的课程 j，且 j 的所有先修课都学完了
		for cus, lb := u-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			j := bits.TrailingZeros(uint(lb))
			if s|pre[j] == s {
				newS := s | lb
				f[newS] = max(f[newS], fs+score[j]*(i+1))
			}
		}
	}
	return f[u-1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n2^n)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(2^n)$。

更多相似题目，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 的「**§9.1 排列型 ① 相邻无关**」。

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
