如果没有 $\textit{word}_2$，那么本题就是 [115. 不同的子序列](https://leetcode.cn/problems/distinct-subsequences/)，请看 [我的题解](https://leetcode.cn/problems/distinct-subsequences/solutions/3060706/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-9va6/)。

我们要在 115 题的状态定义的基础上，增加一个维度。

定义 $\textit{dfs}(i,j,k)$ 表示在 $\textit{word}_1$ 的前缀 $[0,j]$ 和 $\textit{word}_2$ 的前缀 $[0,k]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i]$ 的方案数。这里允许其中一个子序列为空。

讨论 $\textit{word}_1[j]$ 和 $\textit{word}_2[k]$ **选或不选**（是否在 $\textit{target}$ 中），有四种情况：

1. 都不选。
2. 选 $\textit{word}_1[j]$，不选 $\textit{word}_2[k]$。
3. 不选 $\textit{word}_1[j]$，选 $\textit{word}_2[k]$。
4. 都选。意思是 $\textit{word}_1[j]$ 和 $\textit{word}_2[k]$ 都在 $\textit{target}$ 中。

如果不选 $\textit{word}_1[j]$，问题变成在 $\textit{word}_1$ 的前缀 $[0,j-1]$ 和 $\textit{word}_2$ 的前缀 $[0,k]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i]$ 的方案数，即 $\textit{dfs}(i,j-1,k)$。注意 $\textit{dfs}(i,j-1,k)$ 包含了情况一和情况三。

如果不选 $\textit{word}_2[k]$，问题变成在 $\textit{word}_1$ 的前缀 $[0,j]$ 和 $\textit{word}_2$ 的前缀 $[0,k-1]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i]$ 的方案数，即 $\textit{dfs}(i,j,k-1)$。注意 $\textit{dfs}(i,j,k-1)$ 包含了情况一和情况二。

所以 $\textit{dfs}(i,j-1,k)$ 和 $\textit{dfs}(i,j,k-1)$ 的**交集**是情况一，也就是在 $\textit{word}_1$ 的前缀 $[0,j-1]$ 和 $\textit{word}_2$ 的前缀 $[0,k-1]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i]$ 的方案数，即 $\textit{dfs}(i,j-1,k-1)$。

根据**容斥原理**，前三种情况的方案数之和为

$$
\textit{dfs}(i,j-1,k) + \textit{dfs}(i,j,k-1) - \textit{dfs}(i,j-1,k-1)
$$

下面来算情况四（$\textit{word}_1[j]$ 和 $\textit{word}_2[k]$ 都在 $\textit{target}$ 中）的方案数。

此时要么 $\textit{word}_1[j]$ 与 $\textit{target}[i]$ 匹配，要么 $\textit{word}_2[k]$ 与 $\textit{target}[i]$ 匹配。**解释**：如果两个字符都不与 $\textit{target}[i]$ 匹配，而是和 $\textit{target}$ 中更靠前的字符匹配，那么 $\textit{target}[i]$ 匹配谁呢？只能匹配 $\textit{word}_1$ 或 $\textit{word}_2$ 中的一个更靠前的字符，但这样我们选出的下标就不是严格递增的了。

> 顺带一提，对于情况二，只保证 $\textit{word}_1[j]$ 在 $\textit{target}$ 中，并不能说明 $\textit{word}_1[j]$ 一定要与 $\textit{target}[i]$ 匹配。情况三同理。

如果 $\textit{word}_1[j] = \textit{target}[i]$，那么 $\textit{word}_1[j]$ 可以直接选，此时我们需要计算：

- 在 $\textit{word}_1$ 的前缀 $[0,j-1]$ 和 $\textit{word}_2$ 的前缀 $[0,k]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i-1]$ 的方案数，其中 $\textit{word}_2[k]$ 必须选。

这也可以用容斥计算。

- 首先，算出在 $\textit{word}_1$ 的前缀 $[0,j-1]$ 和 $\textit{word}_2$ 的前缀 $[0,k]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i-1]$ 的方案数，即 $\textit{dfs}(i-1,j-1,k)$。
- 然后，减去其中不选 $\textit{word}_2[k]$ 的方案数，也就是在 $\textit{word}_1$ 的前缀 $[0,j-1]$ 和 $\textit{word}_2$ 的前缀 $[0,k-1]$ 中选择字符（选择子序列），组成 $\textit{target}$ 的前缀 $[0,i-1]$ 的方案数，即 $\textit{dfs}(i-1,j-1,k-1)$。

得到

$$
\textit{dfs}(i-1,j-1,k) - \textit{dfs}(i-1,j-1,k-1)
$$

同理可得，如果 $\textit{word}_2[k] = \textit{target}[i]$，那么「都选」的方案数为

$$
\textit{dfs}(i-1,j,k-1) - \textit{dfs}(i-1,j-1,k-1)
$$

四种情况方案数相加，即 $\textit{dfs}(i,j,k)$。

**递归边界**：

- 如果 $j<-1$ 或者 $k<-1$，出界，返回 $0$。
- 如果 $(j+1)+(k+1)<i+1$，字符个数不足以形成 $\textit{target}$ 的前缀 $[0,i]$，返回 $0$。
- 否则，如果 $i<0$，那么不选字符也是一种方案，返回 $1$（或者说我们找到了一种形成 $\textit{target}$ 的方案）。

**递归入口**：$\textit{dfs}(n-1,m_1-1,m_2-1)$。其中 $n$ 是 $\textit{target}$ 的长度，$m_1$ 是 $\textit{word}_1$ 的长度，$m_2$ 是 $\textit{word}_2$ 的长度。

最后，减去只从 $\textit{word}_1$ 中选字符的方案数，以及只从 $\textit{word}_2$ 中选字符的方案数。这就是 115 题。

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

关于记忆化搜索的原理，请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~ **其他语言，以及递推写法，直播结束后补充**。

```py [sol-Python3]
class Solution:
    def numDistinct(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        if n < m:
            return 0

        MOD = 1_000_000_007
        f = [1] + [0] * m
        for i, x in enumerate(s):
            for j in range(min(i, m - 1), max(m - n + i, 0) - 1, -1):
                if x == t[j]:
                    f[j + 1] = (f[j + 1] + f[j]) % MOD
        return f[m]

    def interleaveCharacters(self, word1: str, word2: str, target: str) -> int:
        MOD = 1_000_000_007
        n, m1, m2 = len(target), len(word1), len(word2)

        @cache
        def dfs(i: int, j: int, k: int) -> int:
            if j < -1 or k < -1 or j + k + 2 < i + 1:
                return 0
            if i < 0:
                return 1

            # 不选 word1[j] 或 word2[k]
            res = dfs(i, j - 1, k) + dfs(i, j, k - 1) - dfs(i, j - 1, k - 1)  # 容斥
            if j >= 0 and word1[j] == target[i]:
                # 选 word1[j]
                res += dfs(i - 1, j - 1, k) - dfs(i - 1, j - 1, k - 1)
            if k >= 0 and word2[k] == target[i]:
                # 选 word2[k]
                res += dfs(i - 1, j, k - 1) - dfs(i - 1, j - 1, k - 1)
            return res % MOD

        return (dfs(n - 1, m1 - 1, m2 - 1) - self.numDistinct(word1, target) - self.numDistinct(word2, target)) % MOD
```

```go [sol-Go]
const mod = 1_000_000_007

// 115. 不同的子序列
func numDistinct(s, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	f := make([]int, m+1)
	f[0] = 1
	for i, x := range s {
		for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
			if byte(x) == t[j] {
				f[j+1] = (f[j+1] + f[j]) % mod
			}
		}
	}
	return f[m]
}

func interleaveCharacters(word1, word2, target string) int {
	n, m1, m2 := len(target), len(word1), len(word2)

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m1+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m2+1)
			for p := range memo[i][j] {
				memo[i][j][p] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if j < -1 || k < -1 || j+k+2 < i+1 {
			return 0
		}
		if i < 0 {
			return 1
		}
		p := &memo[i][j+1][k+1]
		if *p != math.MinInt {
			return *p
		}

		// 不选 word1[j] 或 word2[k]
		res := dfs(i, j-1, k) + dfs(i, j, k-1) - dfs(i, j-1, k-1) // 容斥
		if j >= 0 && word1[j] == target[i] {
			// 选 word1[j]
			res += dfs(i-1, j-1, k) - dfs(i-1, j-1, k-1)
		}
		if k >= 0 && word2[k] == target[i] {
			// 选 word2[k]
			res += dfs(i-1, j, k-1) - dfs(i-1, j-1, k-1)
		}
		res %= mod

		*p = res
		return res
	}

	ans := dfs(n-1, m1-1, m2-1) - numDistinct(word1, target) - numDistinct(word2, target)
	return (ans%mod + mod) % mod // 保证 ans 非负
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm_1m_2)$，其中 $n$ 是 $\textit{target}$ 的长度，$m_1$ 是 $\textit{word}_1$ 的长度，$m_2$ 是 $\textit{word}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(nm_1m_2)$。

## 专题训练

见下面动态规划题单的「**§4.1 最长公共子序列**」。

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
