## 位运算恒等式

根据 [从集合论到位运算](https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/)，我们可以把二进制数看成集合，那么异或运算即集合的**对称差**，二进制中的 $1$ 的个数即**集合大小**。

对于集合 $A$ 和集合 $B$，我们有

$$
|A\ \Delta\ B| = |A| + |B| - 2|A\cap B|
$$

所以

$$
|A\ \Delta\ B|\bmod 2 = (|A| + |B| - 2|A\cap B|)\bmod 2 = (|A| + |B|)\bmod 2
$$

设 $p(x)$ 是 $x$ 的二进制中的 $1$ 的个数。从位运算角度来看，上式为

$$
p(a\oplus b)\bmod 2 = (p(a) + p(b)) \bmod 2
$$

上式可以推广到多个数的情况。

所以我们只需关心点权的二进制中的 $1$ 的个数的**奇偶性**，每个点权 $\textit{values}[x]$ 都可以简化成 $p(\textit{values}[x])\bmod 2$，这只有 $0$ 或者 $1$，大幅缩小了问题的规模。

## 方法一：树形 DP

这里介绍一个通用做法，把题干中的「最小」改成「最大」也能做。

对于树上路径 $a\text{-}b$ 的点权异或和，常用套路是转化成如下三部分的异或和：

1. 从根到 $a$ 的点权异或和。
2. 从根到 $b$ 的点权异或和。
3. $a$ 和 $b$ 的最近公共祖先（LCA）的点权异或了两次，抵消了，要再异或一次。

对于树上路径 $a\text{-}b$ 的边权和，转化成如下两部分的边权和：

1. 从 LCA 到 $a$ 的边权和。
2. 从 LCA 到 $b$ 的边权和。

接下来的做法类似 [树的直径【基础算法精讲 23】](https://www.bilibili.com/video/BV17o4y187h1/)，枚举路径拐点 $x$ 作为 LCA，在遍历 $x$ 的子树的过程中，维护：

1. 从根到子树节点 $a$ 的点权异或和的二进制中的 $1$ 的个数是**偶数**时，从 $x$ 到 $a$ 的边权和的最小值 $\textit{minSum}[0]$。
2. 从根到子树节点 $a$ 的点权异或和的二进制中的 $1$ 的个数是**奇数**时，从 $x$ 到 $a$ 的边权和的最小值 $\textit{minSum}[1]$。

把两棵子树中的两条最优路径拼起来，即为拐点为 $x$ 的路径边权和。讨论上文「三部分的异或和」中的 LCA 的值：

- 如果 $\textit{values}[x]$ 的二进制中的 $1$ 的个数是偶数，那么把两棵子树的 $\textit{minSum}[0]$ 相加，或者两棵子树的 $\textit{minSum}[1]$ 相加。
- 如果 $\textit{values}[x]$ 的二进制中的 $1$ 的个数是奇数，那么把一棵子树的 $\textit{minSum}[0]$ 与另一棵子树的 $\textit{minSum}[1]$ 相加。

[本题视频讲解](https://www.bilibili.com/video/BV1kfxW63EpG/?t=30m7s)，额外讲解了其他位运算恒等式，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minActivationCost(self, edges: list[list[int]], values: list[int]) -> int:
        g = [[] for _ in values]
        for x, y, c in edges:
            g[x].append((y, c))
            g[y].append((x, c))

        # p_ones 为根到 x 的点权异或和的二进制中的 1 的个数的奇偶性
        # 返回从 x 到其后代的最小边权和（按照 p_ones 的奇偶性分别计算）
        def dfs(x: int, fa: int, p_ones: int) -> list[int]:
            nonlocal ans
            p_val = values[x].bit_count() % 2
            p_ones ^= p_val
            min_sum = [inf, inf]
            min_sum[p_ones] = 0  # x 作为路径的端点

            for y, c in g[x]:
                if y == fa:
                    continue
                mn = dfs(y, x, p_ones)
                mn[0] += c
                mn[1] += c
                # 现在 mn 是从 x 到子树 y 中节点的最小边权和
                # 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
                ans = min(ans, min_sum[0] + mn[p_val], min_sum[1] + mn[p_val ^ 1])
                min_sum[0] = min(min_sum[0], mn[0])
                min_sum[1] = min(min_sum[1], mn[1])

            return min_sum

        ans = inf
        dfs(0, -1, 0)
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    private static final long INF = Long.MAX_VALUE / 4;
    private long ans = INF;

    public int minActivationCost(int[][] edges, int[] values) {
        List<int[]>[] g = new ArrayList[values.length];
        Arrays.setAll(g, _ -> new ArrayList<>());

        for (int[] e : edges) {
            int x = e[0], y = e[1], c = e[2];
            g[x].add(new int[]{y, c});
            g[y].add(new int[]{x, c});
        }

        dfs(0, -1, 0, g, values);
        return ans < INF ? (int) ans : -1;
    }

    // pOnes 为根到 x 的点权异或和的二进制中的 1 的个数的奇偶性
    // 返回从 x 到其后代的最小边权和（按照 pOnes 的奇偶性分别计算）
    private long[] dfs(int x, int fa, int pOnes, List<int[]>[] g, int[] values) {
        int pVal = Integer.bitCount(values[x]) % 2;
        pOnes ^= pVal;
        long[] minSum = {INF, INF};
        minSum[pOnes] = 0; // x 作为路径的端点

        for (int[] e : g[x]) {
            int y = e[0], c = e[1];
            if (y == fa) {
                continue;
            }
            long[] mn = dfs(y, x, pOnes, g, values);
            mn[0] += c;
            mn[1] += c;
            // 现在 mn 是从 x 到子树 y 中节点的最小边权和
            // 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
            ans = Math.min(ans, Math.min(minSum[0] + mn[pVal], minSum[1] + mn[pVal ^ 1]));
            minSum[0] = Math.min(minSum[0], mn[0]);
            minSum[1] = Math.min(minSum[1], mn[1]);
        }

        return minSum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minActivationCost(vector<vector<int>>& edges, vector<int>& values) {
        int n = values.size();
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1], c = e[2];
            g[x].emplace_back(y, c);
            g[y].emplace_back(x, c);
        }

        constexpr long long INF = LLONG_MAX / 4;
        vector<array<long long, 2>> f(n, {INF, INF}); // 在 dfs 外面创建 array 更快一些，可以对比 C++ 写法二
        long long ans = INF;

        // p_ones 为根到 x 的点权异或和的二进制中的 1 的个数的奇偶性
        // 返回从 x 到其后代的最小边权和（按照 p_ones 的奇偶性分别计算）
        auto dfs = [&](this auto&& dfs, int x, int fa, int p_ones) -> void {
            int p_val = popcount(1u * values[x]) % 2;
            p_ones ^= p_val;
            auto& min_sum = f[x];
            min_sum[p_ones] = 0; // x 作为路径的端点

            for (auto& [y, c] : g[x]) {
                if (y == fa) {
                    continue;
                }
                dfs(y, x, p_ones);
                auto& mn = f[y];
                mn[0] += c;
                mn[1] += c;
                // 现在 mn 是从 x 到子树 y 中节点的最小边权和
                // 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
                ans = min(ans, min(min_sum[0] + mn[p_val], min_sum[1] + mn[p_val ^ 1]));
                min_sum[0] = min(min_sum[0], mn[0]);
                min_sum[1] = min(min_sum[1], mn[1]);
            }
        };

        dfs(0, -1, 0);
        return ans < INF ? ans : -1;
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
public:
    int minActivationCost(vector<vector<int>>& edges, vector<int>& values) {
        vector<vector<pair<int, int>>> g(values.size());
        for (auto& e : edges) {
            int x = e[0], y = e[1], c = e[2];
            g[x].emplace_back(y, c);
            g[y].emplace_back(x, c);
        }

        constexpr long long INF = LLONG_MAX / 4;
        long long ans = INF;

        // p_ones 为根到 x 的点权异或和的二进制中的 1 的个数的奇偶性
        // 返回从 x 到其后代的最小边权和（按照 p_ones 的奇偶性分别计算）
        auto dfs = [&](this auto&& dfs, int x, int fa, int p_ones) -> array<long long, 2> {
            int p_val = popcount(1u * values[x]) % 2;
            p_ones ^= p_val;
            // ⚠注意：测试发现在 dfs 中创建 array 很慢，改成在 dfs 外面创建会快很多
            array<long long, 2> min_sum = {INF, INF};
            min_sum[p_ones] = 0; // x 作为路径的端点

            for (auto& [y, c] : g[x]) {
                if (y == fa) {
                    continue;
                }
                auto mn = dfs(y, x, p_ones);
                mn[0] += c;
                mn[1] += c;
                // 现在 mn 是从 x 到子树 y 中节点的最小边权和
                // 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
                ans = min(ans, min(min_sum[0] + mn[p_val], min_sum[1] + mn[p_val ^ 1]));
                min_sum[0] = min(min_sum[0], mn[0]);
                min_sum[1] = min(min_sum[1], mn[1]);
            }

            return min_sum;
        };

        dfs(0, -1, 0);
        return ans < INF ? ans : -1;
    }
};
```

```go [sol-Go]
func minActivationCost(edges [][]int, values []int) int {
	type edge struct{ to, c int }
	g := make([][]edge, len(values))
	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, c})
		g[y] = append(g[y], edge{x, c})
	}

	const inf = math.MaxInt / 4
	ans := inf

	// pOnes 为根到 x 的点权异或和的二进制中的 1 的个数的奇偶性
	// 返回从 x 到其后代的最小边权和（按照 pOnes 的奇偶性分别计算）
	var dfs func(int, int, int) [2]int
	dfs = func(x, fa, pOnes int) [2]int {
		pVal := bits.OnesCount(uint(values[x])) % 2
		pOnes ^= pVal
		minSum := [2]int{inf, inf}
		minSum[pOnes] = 0 // x 作为路径的端点

		for _, e := range g[x] {
			y := e.to
			if y == fa {
				continue
			}
			mn := dfs(y, x, pOnes)
			mn[0] += e.c
			mn[1] += e.c
			// 现在 mn 是从 x 到子树 y 中节点的最小边权和
			// 拼接从 x 到两棵子树中的路径（拼接后，路径在 x 拐弯）
			ans = min(ans, minSum[0]+mn[pVal], minSum[1]+mn[pVal^1])
			minSum[0] = min(minSum[0], mn[0])
			minSum[1] = min(minSum[1], mn[1])
		}

		return minSum
	}
	dfs(0, -1, 0)

	if ans < inf {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{values}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：脑筋急转弯

假设存在节点个数 $\ge 4$ 的最优路径，那么：

- 如果路径中有两个相邻节点的点权（已简化成 $0$ 和 $1$）相同，那么这两个节点组成了平衡异或路径。由于本题边权为正，所以这两个节点组成的路径的边权和更小，矛盾。
- 否则，路径中的点权是 $0$ 和 $1$ 交替出现的。无论是 $0,1,0,1,\ldots$ 还是 $1,0,1,0,\ldots$，都包含点权为 $1,0,1$ 的子路径，这是平衡异或路径，边权和更小，矛盾。

所以用反证法可以证明如下结论：

- **如果存在最优路径，那么其节点个数至多为** $3$。

所以只需遍历边权以及相邻的边权（求和），计算其中的最小值。

### 优化前

```py [sol-Python3]
class Solution:
    def minActivationCost(self, edges: list[list[int]], values: list[int]) -> int:
        ans = inf
        groups = [[] for _ in values]

        for x, y, c in edges:
            vx = values[x].bit_count() % 2
            vy = values[y].bit_count() % 2
            if vx == vy:
                ans = min(ans, c)
            elif vx == 0:  # 此时 vy = 1
                # 三节点路径的点权必须是 1,0,1，中间的点权必须是 0
                groups[x].append(c)
            else:  # 此时 vy = 0 且 vx = 1
                groups[y].append(c)

        # 枚举三节点路径的中间节点
        for cs in groups:
            # 枚举相邻边权，计算两数之和的最小值（枚举右维护左）
            mn = inf
            for c in cs:
                ans = min(ans, mn + c)
                mn = min(mn, c)

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minActivationCost(int[][] edges, int[] values) {
        int ans = Integer.MAX_VALUE;
        List<Integer>[] groups = new ArrayList[values.length];
        Arrays.setAll(groups, _ -> new ArrayList<>());

        for (int[] e : edges) {
            int x = e[0], y = e[1], c = e[2];
            int vx = Integer.bitCount(values[x]) % 2;
            int vy = Integer.bitCount(values[y]) % 2;
            if (vx == vy) {
                ans = Math.min(ans, c);
            } else if (vx == 0) { // 此时 vy = 1
                // 三节点路径的点权必须是 1,0,1，中间的点权必须是 0
                groups[x].add(c);
            } else { // 此时 vy = 0 且 vx = 1
                groups[y].add(c);
            }
        }

        // 枚举三节点路径的中间节点
        for (List<Integer> cs : groups) {
            // 枚举相邻边权，计算两数之和的最小值（枚举右维护左）
            int mn = Integer.MAX_VALUE;
            for (int c : cs) {
                ans = (int) Math.min(ans, (long) mn + c);
                mn = Math.min(mn, c);
            }
        }

        return ans < Integer.MAX_VALUE ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minActivationCost(vector<vector<int>>& edges, vector<int>& values) {
        int ans = INT_MAX;
        vector<vector<int>> groups(values.size());

        for (auto& e : edges) {
            int x = e[0], y = e[1], c = e[2];
            int vx = popcount(1u * values[x]) % 2;
            int vy = popcount(1u * values[y]) % 2;
            if (vx == vy) {
                ans = min(ans, c);
            } else if (vx == 0) { // 此时 vy = 1
                // 三节点路径的点权必须是 1,0,1，中间的点权必须是 0
                groups[x].push_back(c);
            } else { // 此时 vy = 0 且 vx = 1
                groups[y].push_back(c);
            }
        }

        // 枚举三节点路径的中间节点
        for (auto& cs : groups) {
            // 枚举相邻边权，计算两数之和的最小值（枚举右维护左）
            int mn = INT_MAX;
            for (int c : cs) {
                ans = min(1LL * ans, 1LL * mn + c);
                mn = min(mn, c);
            }
        }

        return ans < INT_MAX ? ans : -1;
    }
};
```

```go [sol-Go]
func minActivationCost(edges [][]int, values []int) int {
	const inf = math.MaxInt / 2
	ans := inf
	groups := make([][]int, len(values))

	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		vx := bits.OnesCount(uint(values[x])) % 2
		vy := bits.OnesCount(uint(values[y])) % 2
		if vx == vy {
			ans = min(ans, c)
		} else if vx == 0 { // 此时 vy = 1
			// 三节点路径的点权必须是 1,0,1，中间的点权必须是 0
			groups[x] = append(groups[x], c)
		} else { // 此时 vy = 0 且 vx = 1
			groups[y] = append(groups[y], c)
		}
	}

	// 枚举三节点路径的中间节点
	for _, cs := range groups {
		// 枚举相邻边权，计算两数之和的最小值（枚举右维护左）
		mn := inf
		for _, c := range cs {
			ans = min(ans, mn+c)
			mn = min(mn, c)
		}
	}

	if ans < inf {
		return ans
	}
	return -1
}
```

### 优化：一次遍历

把上面代码的两个循环合并为一个循环。

```py [sol-Python3]
class Solution:
    def minActivationCost(self, edges: list[list[int]], values: list[int]) -> int:
        ans = inf
        mn = [inf] * len(values)

        for x, y, c in edges:
            vx = values[x].bit_count() % 2
            vy = values[y].bit_count() % 2
            if vx == vy:
                ans = min(ans, c)
            else:
                center = x if vy else y
                ans = min(ans, mn[center] + c)
                mn[center] = min(mn[center], c)

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minActivationCost(int[][] edges, int[] values) {
        int ans = Integer.MAX_VALUE;
        int[] mn = new int[values.length];
        Arrays.fill(mn, Integer.MAX_VALUE);

        for (int[] e : edges) {
            int x = e[0], y = e[1], c = e[2];
            int vx = Integer.bitCount(values[x]) & 1;
            int vy = Integer.bitCount(values[y]) & 1;
            if (vx == vy) {
                ans = Math.min(ans, c);
            } else {
                int center = vx == 0 ? x : y;
                ans = (int) Math.min(ans, (long) mn[center] + c);
                mn[center] = Math.min(mn[center], c);
            }
        }

        return ans < Integer.MAX_VALUE ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minActivationCost(vector<vector<int>>& edges, vector<int>& values) {
        int ans = INT_MAX;
        vector<int> mn(values.size(), INT_MAX);

        for (auto& e : edges) {
            int x = e[0], y = e[1], c = e[2];
            int vx = popcount(1u * values[x]) & 1;
            int vy = popcount(1u * values[y]) & 1;
            if (vx == vy) {
                ans = min(ans, c);
            } else {
                int center = vy ? x : y;
                ans = min(1LL * ans, 1LL * mn[center] + c);
                mn[center] = min(mn[center], c);
            }
        }

        return ans < INT_MAX ? ans : -1;
    }
};
```

```go [sol-Go]
func minActivationCost(edges [][]int, values []int) int {
	const inf = math.MaxInt / 2
	ans := inf
	mn := make([]int, len(values))
	for i := range mn {
		mn[i] = inf
	}

	for _, e := range edges {
		x, y, c := e[0], e[1], e[2]
		vx := bits.OnesCount(uint(values[x])) & 1
		vy := bits.OnesCount(uint(values[y])) & 1
		if vx == vy {
			ans = min(ans, c)
		} else if vx == 0 { // 此时 vy = 1
			ans = min(ans, mn[x]+c)
			mn[x] = min(mn[x], c)
		} else { // 此时 vy = 0 且 vx = 1
			ans = min(ans, mn[y]+c)
			mn[y] = min(mn[y], c)
		}
	}

	if ans < inf {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{values}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面树题单的「**§3.5 树的直径**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
