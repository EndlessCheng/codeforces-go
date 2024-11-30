从特殊到一般。想一想，如果这棵树是一条链，且 $k=1$，要怎么选？

由于不能同时选两条相邻的边，所以问题变成：

- 给你一个长为 $n-1$ 的 $w$ 数组，你需要从中选择若干元素，且不能选相邻的元素。你选的元素之和的最大值是多少？

这就是 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)。

> 既然特殊情况都只能用 DP 解决，那就全力往 DP 思考吧。

本题是树，考虑节点 $x$ 和它的儿子 $y$ 的这条边（$x\text{-}y$）**选或不选**：

- 不选：那么在节点 $y$ 及其儿子的边中，至多选 $k$ 条边。
- 选：那么在节点 $y$ 及其儿子的边中，至多选 $k-1$ 条边。

假设节点 $x$ 有三个儿子，不选和选计算出的结果分别记作 $(\textit{nc}_1, c_1),(\textit{nc}_2, c_2),(\textit{nc}_3, c_3)$。

假设要从中选两条边，选哪两条边最优呢？

- 先考虑都不选，也就是 $\textit{nc}_1+\textit{nc}_2+\textit{nc}_3$。
- 然后把其中两个 $\textit{nc}_i$ 替换成 $c_i$，那么选「增量」最大的两个 $c_i - \textit{nc}_i$。

所以本题不仅是 DP，还是贪心。我们需要把 $c_i - \textit{nc}_i$ 保存到一个数组 $\textit{inc}$ 中（非正数不需要保存），然后把数组从大到小排序，取最大的 $k$ 个或者 $k-1$ 个。

**优化**：如果不删除边也满足要求，即所有点的度数都 $\le k$，则直接返回所有边权之和。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1fFB4YGEZY/?t=26m45s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximizeSumOfWeights(self, edges: List[List[int]], k: int) -> int:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y, wt in edges:
            g[x].append((y, wt))
            g[y].append((x, wt))

        # 优化
        if all(len(to) <= k for to in g):
            return sum(e[2] for e in edges)

        def dfs(x: int, fa: int) -> Tuple[int, int]:
            not_choose = 0
            inc = []
            for y, wt in g[x]:
                if y == fa:
                    continue
                nc, c = dfs(y, x)
                not_choose += nc  # 先都不选
                if (d := c + wt - nc) > 0:
                    inc.append(d)
            inc.sort(reverse=True)
            # 再选增量最大的 k 个或者 k-1 个
            return not_choose + sum(inc[:k]), not_choose + sum(inc[:k - 1])
        return dfs(0, -1)[0]  # not_choose >= choose
```

```java [sol-Java]
class Solution {
    public long maximizeSumOfWeights(int[][] edges, int k) {
        List<int[]>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        long sumWt = 0;
        for (int[] e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].add(new int[]{y, wt});
            g[y].add(new int[]{x, wt});
            sumWt += wt;
        }

        // 优化
        boolean simple = true;
        for (List<int[]> to : g) {
            if (to.size() > k) {
                simple = false;
                break;
            }
        }
        if (simple) {
            return sumWt;
        }

        return dfs(0, -1, g, k)[0]; // notChoose >= choose
    }

    private long[] dfs(int x, int fa, List<int[]>[] g, int k) {
        long notChoose = 0;
        List<Integer> inc = new ArrayList<>();
        for (int[] e : g[x]) {
            int y = e[0];
            if (y == fa) {
                continue;
            }
            long[] res = dfs(y, x, g, k);
            notChoose += res[0]; // 先都不选
            int d = (int) (res[1] - res[0]) + e[1];
            if (d > 0) {
                inc.add(d);
            }
        }

        // 再选增量最大的 k 个或者 k-1 个
        inc.sort(Collections.reverseOrder());
        for (int i = 0; i < Math.min(inc.size(), k - 1); i++) {
            notChoose += inc.get(i);
        }
        long choose = notChoose;
        if (inc.size() >= k) {
            notChoose += inc.get(k - 1);
        }
        return new long[]{notChoose, choose};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximizeSumOfWeights(vector<vector<int>>& edges, int k) {
        vector<vector<pair<int, int>>> g(edges.size() + 1);
        long long sum_wt = 0;
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            g[y].emplace_back(x, wt);
            sum_wt += wt;
        }

        // 优化
        bool simple = true;
        for (auto& to : g) {
            if (to.size() > k) {
                simple = false;
                break;
            }
        }
        if (simple) {
            return sum_wt;
        }

        auto dfs = [&](auto& dfs, int x, int fa) -> pair<long long, long long> {
            long long not_choose = 0;
            vector<int> inc;
            for (auto& [y, wt] : g[x]) {
                if (y == fa) {
                    continue;
                }
                auto [nc, c] = dfs(dfs, y, x);
                not_choose += nc; // 先都不选
                int d = c + wt - nc;
                if (d > 0) {
                    inc.push_back(d);
                }
            }

            // 再选增量最大的 k 个或者 k-1 个
            ranges::sort(inc, greater()); // 从大到小排序
            for (int i = 0; i < min((int) inc.size(), k - 1); i++) {
                not_choose += inc[i];
            }
            long long choose = not_choose;
            if (inc.size() >= k) {
                not_choose += inc[k - 1];
            }
            return {not_choose, choose};
        };
        return dfs(dfs, 0, -1).first; // not_choose >= choose
    }
};
```

```go [sol-Go]
func maximizeSumOfWeights(edges [][]int, k int) int64 {
	type edge struct{ to, wt int }
	g := make([][]edge, len(edges)+1)
	sumWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
		sumWt += wt
	}

	// 优化
	simple := true
	for _, to := range g {
		if len(to) > k {
			simple = false
			break
		}
	}
	if simple {
		return int64(sumWt)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notChoose := 0
		inc := []int{}
		for _, e := range g[x] {
			y := e.to
			if y == fa {
				continue
			}
			nc, c := dfs(y, x)
			notChoose += nc // 先都不选
			if d := c + e.wt - nc; d > 0 {
				inc = append(inc, d)
			}
		}

		// 再选增量最大的 k 个或者 k-1 个
		slices.SortFunc(inc, func(a, b int) int { return b - a })
		for i := range min(len(inc), k-1) {
			notChoose += inc[i]
		}
		choose := notChoose
		if len(inc) >= k {
			notChoose += inc[k-1]
		}
		return notChoose, choose
	}
	nc, _ := dfs(0, -1) // notChoose >= choose
	return int64(nc)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。如果用快速选择，可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

贪心的那部分和 [2611. 老鼠和奶酪](https://leetcode.cn/problems/mice-and-cheese/) 是一模一样的。

更多相似题目，见下面动态规划题单中的「**十二、树形 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
