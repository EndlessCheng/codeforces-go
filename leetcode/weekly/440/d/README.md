## 不删除冲突对

先考虑不删除冲突对，要怎么做。也就是统计合法子数组的个数，对于任意冲突对 $(a,b)$，$a$ 和 $b$ 不能都在子数组中。

既然是统计子数组个数，我们可以**枚举子数组左端点，统计有多少个合法的右端点**。（注：枚举左端点仅仅是为了方便写代码求前二小，枚举右端点也是可以的。）

设冲突对中的 $a\le b$。如果 $a>b$ 则交换 $a$ 和 $b$。

举个例子。现在枚举到子数组左端点为 $i=2$，如果满足 $a\ge i$ 的冲突对有 $(2,6),(3,5),(4,7)$，那么子数组的右端点可以是 $2,3,4$，我们找到了 $3$ 个合法子数组。注意子数组的右端点不能比 $4$ 大，因为这会导致子数组包含冲突对 $(3,5)$。

![lc3480-c.png](https://pic.leetcode.cn/1753405055-svTqwA-lc3480-c.png){:width=600px}

从这个例子可以发现，我们需要知道满足 $a\ge i$ 的冲突对中的 $b$ 的**最小值**，记作 $b_0$。（注：无需考虑 $a<i$ 的冲突对，左端点为 $i$ 的子数组一定不包含这些冲突对。）

那么当子数组的左端点**固定**为 $i$ 时，子数组的右端点可以是

$$
i,i+1,i+2,\ldots,b_0 - 1
$$

这一共有

$$
b_0 - i
$$

个。也就是左端点为 $i$ 时，我们找到了 $b_0 - i$ 个合法子数组。

枚举 $i$，维护 $b_0$，累加 $b_0 - i$，即为不删除冲突对时的答案。

如何维护 $b_0$？

把所有冲突对按照 $a$ 分组，得到 $n$ 个列表，第 $i$ 个列表保存着相同的 $a=i$ 对应的所有 $b$。

倒着枚举 $i=n,n-1,n-2,\ldots,1$，用第 $i$ 个列表中的最小的 $b$，更新 $b_0$ 的最小值。此时 $b_0$ 就是满足 $a\ge i$ 的冲突对中的 $b$ 的最小值。（注：$a$ 的范围是 $[i,n]$，这是个后缀，我们要计算后缀中的最小 $b$，所以要倒着枚举 $i$。）

## 删除一个冲突对

讨论删除冲突对 $(a,b)$ 会产生什么影响。

枚举子数组左端点为 $i$，如果删除的是 $a<i$ 或者 $b>b_0$ 的冲突对，不影响答案。

如果删除的是 $b = b_0$ 的冲突对呢？

![lc3480-1-c.png](https://pic.leetcode.cn/1753405453-dVzDyd-lc3480-1-c.png){:width=600px}

设满足 $a\ge i$ 的冲突对中的 $b$ 的**次小值**为 $b_1$。删除 $b = b_0$ 的冲突对后，满足 $a\ge i$ 的冲突对中的 $b$ 的最小值就变成 $b_1$ 了。（注：如果有多个 $b$ 都等于 $b_0$，那么 $b_1=b_0$。）

上文中的式子 $b_0 - i$ 改成 $b_1 - i$。

换句话说，删除冲突对 $(a,b)$ 会**额外增加**

$$
(b_1 - i) - (b_0 - i) = b_1 - b_0
$$

个左端点为 $i$ 的合法子数组。特别地，如果 $b_1 = b_0$，删除冲突对不会额外增加合法子数组。

**删除哪个冲突对最优？看谁带来的额外增量最大。**

不删除冲突对时的答案，加上删除一个冲突对带来的最大额外增量，就是最终答案。

那么，额外增量的最大值，就是 $b_1 - b_0$ 的最大值吗？

没这么简单。对于多个不同的左端点 $i$，如果这些 $i$ 对应的 $b_0$ 都相同，那么删除 $b=b_0$ 的冲突对会让这些左端点**都受益**。所以我们还得累加 $b_1-b_0$。

具体地，把 $b_0$ 相同的增量，累加到一起。创建一个数组 $\textit{extra}$，把 增量 $b_1-b_0$ 累加到 $\textit{extra}[b_0]$ 中。

最终答案就是不删除冲突对时的合法子数组个数，加上 $\max(\textit{extra})$。

代码实现时，一开始没有冲突对，我们初始化 $b_0=b_1=n+1$，这样也能用上述公式计算合法子数组个数。

## 优化前

```py [sol-Python3]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        groups = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            groups[a].append(b)

        ans = 0
        extra = [0] * (n + 2)
        b = [n + 1, n + 1]
        for i in range(n, 0, -1):
            b = sorted(b + groups[i])[:2]  # 维护最小 b 和次小 b
            ans += b[0] - i
            extra[b[0]] += b[1] - b[0]

        return ans + max(extra)
```

```py [sol-Python3 写法二]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        groups = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            groups[a].append(b)

        ans = 0
        extra = [0] * (n + 2)
        b = [n + 1, n + 1]
        for i in range(n, 0, -1):
            # 维护最小 b 和次小 b
            if groups[i]:
                b += groups[i]
                b.sort()
                b = b[:2]
            ans += b[0] - i
            extra[b[0]] += b[1] - b[0]

        return ans + max(extra)
```

```java [sol-Java]
class Solution {
    public long maxSubarrays(int n, int[][] conflictingPairs) {
        List<Integer>[] groups = new ArrayList[n + 1];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int[] p : conflictingPairs) {
            int a = p[0];
            int b = p[1];
            groups[Math.min(a, b)].add(Math.max(a, b));
        }

        long ans = 0;
        long maxExtra = 0;
        long[] extra = new long[n + 2];
        List<Integer> b = new ArrayList<>();
        b.add(n + 1);
        b.add(n + 1);

        for (int i = n; i > 0; i--) {
            // 维护最小 b 和次小 b
            b.addAll(groups[i]);
            Collections.sort(b);
            b.subList(2, b.size()).clear();

            int b0 = b.get(0);
            ans += b0 - i;
            extra[b0] += b.get(1) - b0;
            maxExtra = Math.max(maxExtra, extra[b0]);
        }

        return ans + maxExtra;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarrays(int n, vector<vector<int>>& conflictingPairs) {
        vector<vector<int>> groups(n + 1);
        for (auto& p : conflictingPairs) {
            int a = p[0], b = p[1];
            if (a > b) {
                swap(a, b);
            }
            groups[a].push_back(b);
        }

        long long ans = 0;
        vector<long long> extra(n + 2);
        vector<int> b = {n + 1, n + 1}; 

        for (int i = n; i > 0; i--) {
            // 维护最小 b 和次小 b
            b.insert(b.end(), groups[i].begin(), groups[i].end());
            ranges::sort(b);
            b.resize(2);

            ans += b[0] - i;
            extra[b[0]] += b[1] - b[0];
        }

        return ans + ranges::max(extra);
    }
};
```

```go [sol-Go]
func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	groups := make([][]int, n+1)
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		groups[a] = append(groups[a], b)
	}

	ans := 0
	extra := make([]int, n+2)
	b := []int{n + 1, n + 1} // 维护最小 b 和次小 b

	for i := n; i > 0; i-- {
		b = append(b, groups[i]...)
		slices.Sort(b)
		b = b[:2]
		ans += b[0] - i
		extra[b[0]] += b[1] - b[0]
	}

	return int64(ans + slices.Max(extra))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 优化

去掉排序，改用 $\texttt{if-else}$ 维护前二小。

此外，$\textit{extra}$ 数组可以优化成一个变量。在枚举 $i$ 的过程中，$b_0$ 要么不变，要么减少，所以相同的 $b_0$ 是连续出现的，只需用一个变量 $\textit{extra}$ 维护连续相同 $b_0$ 对应的 $b_1-b_0$ 之和，同时用另一个变量 $\textit{maxExtra}$ 维护 $\textit{extra}$ 的最大值。

```py [sol-Python3]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        # 更快写法见【Python3 写法二】
        groups = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            groups[a].append(b)

        ans = max_extra = extra = 0
        b0 = b1 = n + 1
        for i in range(n, 0, -1):
            pre_b0 = b0
            for b in groups[i]:
                if b < b0:
                    b0, b1 = b, b0
                elif b < b1:
                    b1 = b

            ans += b0 - i
            if b0 != pre_b0:  # 重新统计连续相同 b0 的 extra
                extra = 0
            extra += b1 - b0
            max_extra = max(max_extra, extra)

        return ans + max_extra
```

```py [sol-Python3 写法二]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        g0 = [n + 1] * (n + 1)
        g1 = [n + 1] * (n + 1)
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            if b < g0[a]:
                g1[a] = g0[a]
                g0[a] = b
            elif b < g1[a]:
                g1[a] = b

        ans = max_extra = extra = 0
        b0 = b1 = n + 1
        for i in range(n, 0, -1):
            pre_b0 = b0

            b, c = g0[i], g1[i]
            if b < b0:
                b1 = min(b0, c)
                b0 = b
            elif b < b1:
                b1 = b
            elif c < b1:
                b1 = c

            ans += b0 - i
            if b0 != pre_b0:  # 重新统计连续相同 b0 的 extra
                extra = 0
            extra += b1 - b0
            max_extra = max(max_extra, extra)

        return ans + max_extra
```

```java [sol-Java]
class Solution {
    public long maxSubarrays(int n, int[][] conflictingPairs) {
        // 更快的写法见【Java 写法二】
        List<Integer>[] groups = new ArrayList[n + 1];
        Arrays.setAll(groups, _ -> new ArrayList<>());
        for (int[] p : conflictingPairs) {
            int a = p[0];
            int b = p[1];
            groups[Math.min(a, b)].add(Math.max(a, b));
        }

        long ans = 0;
        long maxExtra = 0;
        long extra = 0;
        int b0 = n + 1;
        int b1 = n + 1;

        for (int i = n; i > 0; i--) {
            int preB0 = b0;
            for (int b : groups[i]) {
                if (b < b0) {
                    b1 = b0;
                    b0 = b;
                } else if (b < b1) {
                    b1 = b;
                }
            }

            ans += b0 - i;
            if (b0 != preB0) { // 重新统计连续相同 b0 的 extra
                extra = 0;
            }
            extra += b1 - b0;
            maxExtra = Math.max(maxExtra, extra);
        }

        return ans + maxExtra;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public long maxSubarrays(int n, int[][] conflictingPairs) {
        int[] g0 = new int[n + 1];
        int[] g1 = new int[n + 1];
        Arrays.fill(g0, n + 1);
        Arrays.fill(g1, n + 1);

        for (int[] p : conflictingPairs) {
            int a = p[0];
            int b = p[1];
            if (a > b) {
                int tmp = a;
                a = b;
                b = tmp;
            }
            if (b < g0[a]) {
                g1[a] = g0[a];
                g0[a] = b;
            } else if (b < g1[a]) {
                g1[a] = b;
            }
        }

        long ans = 0;
        long maxExtra = 0;
        long extra = 0;
        int b0 = n + 1;
        int b1 = n + 1;

        for (int i = n; i > 0; i--) {
            int preB0 = b0;

            int b = g0[i];
            int c = g1[i];
            if (b < b0) {
                b1 = Math.min(b0, c);
                b0 = b;
            } else if (b < b1) {
                b1 = b;
            } else if (c < b1) {
                b1 = c;
            }

            ans += b0 - i;
            if (b0 != preB0) { // 重新统计连续相同 b0 的 extra
                extra = 0;
            }
            extra += b1 - b0;
            maxExtra = Math.max(maxExtra, extra);
        }

        return ans + maxExtra;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarrays(int n, vector<vector<int>>& conflictingPairs) {
        // vector<array<int, 2>> 比 vector<vector<int>> 快
        vector<array<int, 2>> groups(n + 1, {n + 1, n + 1});
        for (auto& p : conflictingPairs) {
            int a = p[0], b = p[1];
            if (a > b) {
                swap(a, b);
            }
            auto& g = groups[a];
            if (b < g[0]) {
                g[1] = g[0];
                g[0] = b;
            } else if (b < g[1]) {
                g[1] = b;
            }
        }

        long long ans = 0, max_extra = 0, extra = 0;
        int b0 = n + 1, b1 = n + 1;

        for (int i = n; i > 0; i--) {
            int pre_b0 = b0;
            for (int b : groups[i]) {
                if (b < b0) {
                    b1 = b0;
                    b0 = b;
                } else if (b < b1) {
                    b1 = b;
                }
            }

            ans += b0 - i;
            if (b0 != pre_b0) { // 重新统计连续相同 b0 的 extra
                extra = 0;
            }
            extra += b1 - b0;
            max_extra = max(max_extra, extra);
        }

        return ans + max_extra;
    }
};
```

```go [sol-Go]
func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	groups := make([][2]int, n+1) // [][2]int 比 [][]int 快
	for i := range groups {
		groups[i] = [2]int{n + 1, n + 1}
	}
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		g := &groups[a]
		if b < g[0] {
			g[0], g[1] = b, g[0]
		} else if b < g[1] {
			g[1] = b
		}
	}

	var ans, maxExtra, extra int
	b0, b1 := n+1, n+1
	for i := n; i > 0; i-- {
		preB0 := b0
		for _, b := range groups[i] {
			if b < b0 {
				b0, b1 = b, b0
			} else if b < b1 {
				b1 = b
			}
		}

		ans += b0 - i
		if b0 != preB0 { // 重新统计连续相同 b0 的 extra
			extra = 0
		}
		extra += b1 - b0
		maxExtra = max(maxExtra, extra)
	}

	return int64(ans + maxExtra)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果可以删除两个冲突对呢？

欢迎在评论区分享你的思路/代码。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
