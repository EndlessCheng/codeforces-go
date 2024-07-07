[视频讲解](https://www.bilibili.com/video/BV1Ry411q71f/) 讲了为什么字典树做法是 $\mathcal{O}(n^2)$ 的，以及字符串哈希的解题思路。

---

利用**后缀数组**，可以快速计算出每个 $\textit{words}[i]$ 在 $\textit{target}$ 中的出现位置（匹配位置）。

一共有多少个匹配位置？考虑这样一个最坏情况：$\textit{target}$ 全是 $\texttt{a}$，$\textit{words}=[\texttt{a},\texttt{aa},\texttt{aaa},\cdots]$。设 $L$ 是 $\textit{words}$ 中所有字符串的长度之和，在这种情况下，$\textit{words}$ 中有 $\mathcal{O}(\sqrt L)$ 个字符串。每个字符串都会产生 $\mathcal{O}(n)$ 次匹配，所以一共有 $\mathcal{O}(n\sqrt L)$ 个匹配位置。在本题数据范围下，这是可以接受的。

如果 $\textit{words}[i]$ 与 $\textit{target}$ 的下标 $[l,r)$ 匹配，那么把二元组 $(l, \textit{costs}[i])$ 添加到 $\textit{from}[r]$ 中。

定义 $f[i]$ 表示使 $s$ 等于 $\textit{target}$ 的长为 $i$ 的前缀的最小成本。枚举 $\textit{from}[i]$ 中的二元组，假设我们匹配了 $\textit{target}$ 的下标 $[l,i)$ 这一段子串，那么我们需要解决的问题变成：使 $s$ 等于 $\textit{target}$ 的长为 $l$ 的前缀的最小成本。所以有

$$
f[i] = \min_j\{ f[\textit{from}[i][j].l] + \textit{from}[i][j].\textit{cost}  \}
$$

如果 $\textit{from}[i]$ 是空的，则 $f[i]=\infty$。

初始值：$f[0]=0$。

答案：$f[n]$。如果 $f[n]=\infty$ 则返回 $-1$。

细节：$\textit{words}$ 中可能有相同字符串，这些字符串对应的成本应当取最小的。

```go
func minimumCost(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{L})$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。有多少个匹配，就有多少次状态转移。
- 空间复杂度：$\mathcal{O}(n\sqrt{L})$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
