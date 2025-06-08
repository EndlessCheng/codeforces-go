题目要求子集中 $0$ 到 $9$ 每个数字在所有数的数位中最多出现一次。

把 $\textit{vals}[i]$ 的数位保存到集合 $V_i$ 中（不考虑有重复数位的数字），题目要求转化成：

- 从子树中选择若干没有交集的集合，对应的 $\textit{vals}[i]$ 之和越大越好。

定义 $f_x[S]$ 表示从 $x$ 子树中选择若干没有交集的集合，这些集合的并集为 $S$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值。

枚举 $S$ 的非空真子集 $T$，把集合 $S$ 视作 $T$ 和 $\complement_ST$ 的并集，那么：

- 集合的并集为 $T$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值，即 $f_x[T]$。
- 集合的并集为 $\complement_ST$ 的情况下，对应的 $\textit{vals}[i]$ 之和的最大值，即 $f_x[\complement_ST]$。

二者相加，更新 $f_x[S]$ 的最大值，即

$$
f_x[S] = \max_{T\subseteq S} f_x[T] + f_x[\complement_ST]
$$

初始值：

- 选 $\textit{vals}[x]$，初始化 $f_x[V_x] = \textit{vals}[x]$。
- 枚举 $x$ 的儿子 $y$，由于同一个集合至多选一个（否则就有交集了），所以取最大值得 $f_x[S] = \max\limits_y f_y[S]$。

$\max(f_x)$ 就是题目要求的 $\textit{maxScore}[x]$，加到答案中。

**代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看** [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
fmax = lambda a, b: b if b > a else a

class Solution:
    def goodSubtreeSum(self, vals: List[int], par: List[int]) -> int:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        ans = 0

        def dfs(x: int) -> Dict[int, int]:
            f = defaultdict(int)

            # 计算 vals[x] 的 mask
            mask = 0
            v = vals[x]
            while v:
                v, d = divmod(v, 10)
                if mask >> d & 1:
                    break
                mask |= 1 << d
            else:  # 没有中途 break
                f[mask] = vals[x]

            for y in g[x]:
                fy = dfs(y)
                nf = f.copy()
                for msk, s in fy.items():
                    # 同一个 mask 至多选一个，直接取 max
                    nf[msk] = max(nf[msk], s)
                    # 求两个 mask 的并集，刷表转移
                    for msk2, s2 in f.items():
                        if msk & msk2:
                            continue
                        new_mask = msk | msk2
                        nf[new_mask] = max(nf[new_mask], s + s2)
                f = nf

            if f:
                nonlocal ans
                ans += max(f.values())
            return f

        dfs(0)
        return ans % 1_000_000_007
```

```go [sol-Go]
func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) [1 << D]int
	dfs = func(x int) (f [1 << D]int) {
		// 计算 vals[x] 的数位集合 mask
		mask := 0
		for v := vals[x]; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 { // d 在集合 mask 中
				mask = 0 // 不符合要求
				break
			}
			mask |= 1 << d // 把 d 加到集合 mask 中
		}

		if mask > 0 {
			f[mask] = vals[x]
		}

		// 同一个集合 i 至多选一个，直接取 max
		for _, y := range g[x] {
			fy := dfs(y)
			for i, sum := range fy {
				f[i] = max(f[i], sum)
			}
		}

		for i := range f {
			// 枚举集合 i 的非空真子集 sub
			for sub := i & (i - 1); sub > 0; sub = (sub - 1) & i {
				f[i] = max(f[i], f[sub]+f[i^sub])
			}
		}

		ans += slices.Max(f[:])
		return
	}
	dfs(0)
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\cdot 3^D)$，其中 $n$ 是 $\textit{vals}$ 的长度，$D=10$。大小为 $D$ 的集合的大小为 $m$ 的子集有 $\binom D m$ 个，子集的子集有 $2^m$ 个，根据二项式定理，$\sum\limits_{m=0}^D \binom D m 2^m = (2+1)^D = 3^D$，所以二重循环的时间复杂度为 $\mathcal{O}(3^D)$。
- 空间复杂度：$\mathcal{O}(n\cdot 2^D)$。

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
