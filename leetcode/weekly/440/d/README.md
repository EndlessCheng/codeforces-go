## 不删除冲突对

先考虑不删除冲突对，要怎么做。

既然是统计子数组数量，我们可以**枚举子数组左端点，看看有多少个合法的右端点**。（当然，你想枚举右端点也可以）

下面假设冲突对中的 $a\le b$（如果 $a>b$ 则交换 $a$ 和 $b$）。

例如子数组左端点为 $i=2$，满足 $a\ge i$ 的冲突对有 $[2,6],[3,5],[4,7]$，那么子数组的右端点可以是 $2,3,4$。右端点不能比 $4$ 大，因为这会导致子数组包含冲突对 $[3,5]$。

从这个例子可以发现，我们需要维护满足 $a\ge i$ 的冲突对中的 $b$ 的最小值，记作 $b_0$。

那么子数组左端点**固定**为 $i$ 的时候，右端点可以是

$$
i,i+1,i+2,\ldots,b_0 - 1
$$

所以有

$$
b_0 - i
$$

个左端点在 $i$ 的合法子数组。

累加上式，即为不删除冲突对时的答案。

如何维护 $b_0$？

把所有冲突对按照 $a$ 分组，得到 $n$ 个列表，第 $a$ 个列表就保存着相同的 $a$ 对应的所有 $b$。

倒着枚举 $a=n,n-1,n-2,\ldots,1$。用第 $a$ 个列表中的最小的 $b$，更新 $b_0$ 的最小值。

## 删除一个冲突对

讨论删除冲突对 $[a,b]$ 时，会产生什么影响。

如果 $b$ 恰好等于 $b_0$，那么我们需要知道在 $i$ 右边的**第二小**的 $b$，记作 $b_1$。

那么子数组左端点**固定**为 $i$ 的时候，右端点可以是

$$
i,i+1,i+2,\ldots,b_1 - 1
$$

所以有

$$
b_1 - i
$$

个左端点在 $i$ 的合法子数组。

换句话说，删除冲突对 $[a,b]$ 会**额外增加**

$$
(b_1 - i) - (b_0 - i) = b_1 - b_0
$$

个左端点在 $i$ 的合法子数组。

把这个增量累加到 $\textit{extra}[b_0]$ 中。注意：对于多个不同的左端点 $i$，如果对应的 $b_0$ 都相同，那么把增量 $b_1-b_0$ 都累加到 $\textit{extra}[b_0]$ 中，表示删除 $b=b_0$ 的冲突对会让这些左端点都受益。

最终答案就是不删除冲突对时的合法子数组个数，加上 $\max(\textit{extra})$。

代码实现时，可以初始化 $b_0=b_1=n+1$，这样在右边没有冲突对的时候，我们也能用上述公式计算合法子数组个数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15gRaYZE5o/?t=44m20s)，欢迎点赞关注~

### 答疑

**问**：如果有多个相同的冲突对，代码能算出正确答案吗？

**答**：如果要删除多个相同的冲突对中的一个，因为有多个相同的 $b$，导致最小和次小的 $b$ 相等，所以 $b_1-b_0 = 0$，说明删不删都一样，这符合删除多个相同的冲突对中的一个的情况。

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
        b = [n + 1, n + 1]  # 维护最小 b 和次小 b
        for a in range(n, 0, -1):
            b = sorted(b + groups[a])[:2]
            ans += b[0] - a
            extra[b[0]] += b[1] - b[0]

        return ans + max(extra)
```

```py [sol-Python3 更快写法]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        groups = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            groups[a].append(b)

        ans = 0
        extra = [0] * (n + 2)
        b = [n + 1, n + 1]  # 维护最小 b 和次小 b
        for a in range(n, 0, -1):
            list_b = groups[a]
            if list_b:
                b.extend(list_b)
                b.sort()
                b = b[:2]
            ans += b[0] - a
            extra[b[0]] += b[1] - b[0]

        return ans + max(extra)
```

```java [sol-Java]
class Solution {
    public long maxSubarrays(int n, int[][] conflictingPairs) {
        List<Integer>[] groups = new ArrayList[n + 1];
        Arrays.setAll(groups, i -> new ArrayList<>());
        for (int[] p : conflictingPairs) {
            int a = p[0];
            int b = p[1];
            groups[Math.min(a, b)].add(Math.max(a, b));
        }

        long ans = 0;
        long[] extra = new long[n + 2];
        long maxExtra = 0;
        // 维护最小 b 和次小 b
        int b0 = n + 1;
        int b1 = n + 1;
        for (int a = n; a > 0; a--) {
            for (int b : groups[a]) {
                if (b < b0) {
                    b1 = b0;
                    b0 = b;
                } else if (b < b1) {
                    b1 = b;
                }
            }
            ans += b0 - a;
            extra[b0] += b1 - b0;
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
        vector<int> b = {n + 1, n + 1}; // 维护最小 b 和次小 b
        for (int a = n; a > 0; a--) {
            auto& list_b = groups[a];
            b.insert(b.end(), list_b.begin(), list_b.end());
            ranges::sort(b);
            b.resize(2);
            ans += b[0] - a;
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
	for a := n; a > 0; a-- {
		b = append(b, groups[a]...)
		slices.Sort(b)
		b = b[:2]
		ans += b[0] - a
		extra[b[0]] += b[1] - b[0]
	}

	return int64(ans + slices.Max(extra))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

时间复杂度的瓶颈在排序上，下面来优化。

## 优化

把排序替换成维护前二小，这样时间复杂度就是 $\mathcal{O}(n)$ 了。

此外，$\textit{extra}$ 数组可以优化掉，因为 $b_0$ 只会减少，不会增加，所以相同的 $b_0$ 必定是连续的，我们只需要一个变量维护连续相同 $b_0$ 对应的 $b_1-b_0$ 之和，同时用另一个变量 $\textit{maxExtra}$ 维护 $\textit{extra}$ 的最大值。

```py [sol-Python3]
class Solution:
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        groups = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            if a > b:
                a, b = b, a
            groups[a].append(b)

        ans = extra = max_extra = 0
        b0 = b1 = n + 1
        for a in range(n, 0, -1):
            pre_b0 = b0
            for b in groups[a]:
                if b < b0:
                    b0, b1 = b, b0
                elif b < b1:
                    b1 = b
            ans += b0 - a
            if b0 != pre_b0:
                extra = 0
            extra += b1 - b0
            max_extra = max(max_extra, extra)  # 这里改成手动 if 会快不少

        return ans + max_extra
```

```java [sol-Java]
class Solution {
    public long maxSubarrays(int n, int[][] conflictingPairs) {
        List<Integer>[] groups = new ArrayList[n + 1];
        Arrays.setAll(groups, i -> new ArrayList<>());
        for (int[] p : conflictingPairs) {
            int a = p[0];
            int b = p[1];
            groups[Math.min(a, b)].add(Math.max(a, b));
        }

        long ans = 0;
        long extra = 0;
        long maxExtra = 0;
        int b0 = n + 1;
        int b1 = n + 1;
        for (int a = n; a > 0; a--) {
            int preB0 = b0;
            for (int b : groups[a]) {
                if (b < b0) {
                    b1 = b0;
                    b0 = b;
                } else if (b < b1) {
                    b1 = b;
                }
            }
            ans += b0 - a;
            if (b0 != preB0) {
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
        // vector<array<int, 2>> 比 vector<vector<int>> 快很多
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

        long long ans = 0, extra = 0, max_extra = 0;
        int b0 = n + 1, b1 = n + 1;
        for (int a = n; a > 0; a--) {
            int pre_b0 = b0;
            for (int b : groups[a]) {
                if (b < b0) {
                    b1 = b0;
                    b0 = b;
                } else if (b < b1) {
                    b1 = b;
                }
            }
            ans += b0 - a;
            if (b0 != pre_b0) {
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
	groups := make([][2]int, n+1) // [][2]int 比 [][]int 快很多
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

	var ans, extra, maxExtra int
	b0, b1 := n+1, n+1
	for a := n; a > 0; a-- {
		preB0 := b0
		for _, b := range groups[a] {
			if b < b0 {
				b0, b1 = b, b0
			} else if b < b1 {
				b1 = b
			}
		}
		ans += b0 - a
		if b0 != preB0 {
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

如果可以删两个呢？

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
