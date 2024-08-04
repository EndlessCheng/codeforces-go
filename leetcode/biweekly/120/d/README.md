[本题视频讲解](https://www.bilibili.com/video/BV1jg4y1y7PA/)

首先来看如下问题：

给你一个长度 $\ge 3$ 的数组，返回**三数之积**的最大值，如果最大值是负数，返回 $0$。

- 如果只考虑正数，显然最大的三个数的乘积是最大的。
- 如果把负数也考虑进来，那么最小的两个负数乘最大的正数也可能是最大值。

这两种情况再和 $0$ 取最大值，即为返回值。

由上述讨论可知，只需要知道 $5$ 个数，就能算出三数之积的最大值。每棵子树只需要返回它最小的 $2$ 个 $\textit{cost}$ 值和最大的 $3$ 个 $\textit{cost}$ 值就行，其余数字可以舍弃。

对于一棵子树，把它的所有儿子子树的返回值与当前节点的 $\textit{cost}$ 值排序后，按照三数之积的方法，就得到了当前节点的答案。

```py [sol-Python3]
class Solution:
    def placedCoins(self, edges: List[List[int]], cost: List[int]) -> List[int]:
        n = len(cost)
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        ans = [1] * n
        def dfs(x: int, fa: int) -> List[int]:
            a = [cost[x]]
            for y in g[x]:
                if y != fa:
                    a.extend(dfs(y, x))
            a.sort()
            m = len(a)
            if m >= 3:
                ans[x] = max(a[-3] * a[-2] * a[-1], a[0] * a[1] * a[-1], 0)
            if m > 5:
                a = a[:2] + a[-3:]
            return a
        dfs(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] placedCoins(int[][] edges, int[] cost) {
        int n = cost.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        long[] ans = new long[n];
        dfs(0, -1, cost, g, ans);
        return ans;
    }

    private List<Integer> dfs(int x, int fa, int[] cost, List<Integer>[] g, long[] ans) {
        List<Integer> a = new ArrayList<>();
        a.add(cost[x]);
        for (int y : g[x]) {
            if (y != fa) {
                a.addAll(dfs(y, x, cost, g, ans));
            }
        }
        Collections.sort(a);
        int m = a.size();
        if (m < 3) {
            ans[x] = 1;
        } else {
            ans[x] = Math.max((long) a.get(m - 3) * a.get(m - 2) * a.get(m - 1),
                    Math.max((long) a.get(0) * a.get(1) * a.get(m - 1), 0));
        }
        if (m > 5) {
            a = List.of(a.get(0), a.get(1), a.get(m - 3), a.get(m - 2), a.get(m - 1));
        }
        return a;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> placedCoins(vector<vector<int>>& edges, vector<int>& cost) {
        int n = cost.size();
        vector<vector<int>> g(n);
        for (auto& e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<long long> ans(n, 1);
        auto dfs = [&](auto&& dfs, int x, int fa) -> vector<int> {
            vector<int> a = {cost[x]};
            for (int y: g[x]) {
                if (y != fa) {
                    auto res = dfs(dfs, y, x);
                    a.insert(a.end(), res.begin(), res.end());
                }
            }
            ranges::sort(a);
            int m = a.size();
            if (m >= 3) {
                ans[x] = max(max((long long) a[m - 3] * a[m - 2] * a[m - 1], (long long) a[0] * a[1] * a[m - 1]), 0LL);
            }
            if (m > 5) {
                a = {a[0], a[1], a[m - 3], a[m - 2], a[m - 1]};
            }
            return a;
        };
        dfs(dfs, 0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int64, n)
	var dfs func(int, int) []int
	dfs = func(x, fa int) []int {
		a := []int{cost[x]}
		for _, y := range g[x] {
			if y != fa {
				a = append(a, dfs(y, x)...)
			}
		}

		slices.Sort(a)
		m := len(a)
		if m < 3 {
			ans[x] = 1
		} else {
			ans[x] = int64(max(a[m-3]*a[m-2]*a[m-1], a[0]*a[1]*a[m-1], 0))
		}
		if m > 5 {
			a = append(a[:2], a[m-3:]...)
		}
		return a
	}
	dfs(0, -1)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上，如果手动维护最小的两个数和最大的三个数可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

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
