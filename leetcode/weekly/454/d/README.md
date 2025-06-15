**前置知识**：最近公共祖先（LCA），请看我的 [讲解](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)。

定义：

- 询问中的两个节点分别为 $x$ 和 $y$。
- $\textit{lca}$ 为 $x$ 和 $y$ 的最近公共祖先。
- $\textit{disXY}$ 为 $x$ 到 $y$ 的距离。
- $\textit{half}$ 为 $\textit{disXY}$ 的至少一半，即 $\left\lceil\dfrac{\textit{disXY}}{2}\right\rceil$。

分类讨论：

- 如果 $x$ 到 $\textit{lca}$ 的距离 $\ge \textit{half}$，那么答案在 $x$ 到 $\textit{lca}$ 的路径中，我们需要计算从 $x$ 往上跳**至少** $\textit{half}$ 的节点。
- 否则，答案在 $y$ 到 $\textit{lca}$ 的路径中，我们需要计算从 $y$ 往上跳**至多** $\textit{disXY} - \textit{half}$ 的节点，就是从 $x$ 出发，距离 $x$ **至少** $\textit{half}$ 的节点。

往上跳**至少** $\textit{half}$ 可以转化成：往上跳**至多** $\textit{half}-1$，然后再往上跳一个节点。

现在，两种情况统一为：

- 计算从 $x$ 出发，往上跳**至多** $d$，所能到达的最远节点。

这和 LCA 模板中的 $\texttt{uptoDep}$ 函数类似：

1. 枚举 $i=\textit{mx}-1,\textit{mx}-2,\ldots,0$。
2. 设 $p$ 为 $x$ 的 $2^i$ 级祖先。
3. 如果 $x$ 到 $p$ 的距离 $\le d$，那么可以跳到 $p$，更新 $x$ 为 $p$。
4. 继续枚举下一个 $i$。
5. 最终返回跳到的节点。

注意 $x=y$ 的情况，可以直接把 $x$ 加到答案中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1qsMxz6EEd/?t=28m32s)，欢迎点赞关注~

```py [sol-Python3]
class LcaBinaryLifting:
    def __init__(self, edges: List[List[int]]):
        n = len(edges) + 1
        self.m = m = n.bit_length()
        g = [[] for _ in range(n)]
        for x, y, w in edges:
            g[x].append((y, w))
            g[y].append((x, w))

        depth = [0] * n
        dis = [0] * n
        pa = [[-1] * m for _ in range(n)]

        def dfs(x: int, fa: int) -> None:
            pa[x][0] = fa
            for y, w in g[x]:
                if y != fa:
                    depth[y] = depth[x] + 1
                    dis[y] = dis[x] + w
                    dfs(y, x)

        dfs(0, -1)

        for i in range(m - 1):
            for x in range(n):
                if (p := pa[x][i]) != -1:
                    pa[x][i + 1] = pa[p][i]

        self.depth = depth
        self.dis = dis
        self.pa = pa

    def get_kth_ancestor(self, node: int, k: int) -> int:
        for i in range(k.bit_length()):
            if k >> i & 1:
                node = self.pa[node][i]
        return node

    # 返回 x 和 y 的最近公共祖先
    def get_lca(self, x: int, y: int) -> int:
        if self.depth[x] > self.depth[y]:
            x, y = y, x
        # 使 y 和 x 在同一深度
        y = self.get_kth_ancestor(y, self.depth[y] - self.depth[x])
        if y == x:
            return x
        for i in range(self.m - 1, -1, -1):
            px, py = self.pa[x][i], self.pa[y][i]
            if px != py:
                x, y = px, py  # 同时往上跳 2**i 步
        return self.pa[x][0]

    # 返回 x 到 y 的距离（最短路长度）
    def get_dis(self, x: int, y: int) -> int:
        return self.dis[x] + self.dis[y] - self.dis[self.get_lca(x, y)] * 2

    # 从 x 往上跳【至多】d 距离，返回最远能到达的节点
    def upto_dis(self, x: int, d: int) -> int:
        dx = self.dis[x]
        for i in range(self.m - 1, -1, -1):
            p = self.pa[x][i]
            if p != -1 and dx - self.dis[p] <= d:  # 可以跳至多 d
                x = p
        return x


class Solution:
    def findMedian(self, n: int, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        g = LcaBinaryLifting(edges)
        ans = []
        for x, y in queries:
            if x == y:
                ans.append(x)
                continue
            lca = g.get_lca(x, y)
            dis_xy = g.dis[x] + g.dis[y] - g.dis[lca] * 2
            half = (dis_xy + 1) // 2
            if g.dis[x] - g.dis[lca] >= half:  # 答案在 x-lca 路径中
                # 先往上跳至多 half-1，然后再跳一步，就是至少 half
                to = g.upto_dis(x, half - 1)
                res = g.pa[to][0]  # 再跳一步
            else:  # 答案在 y-lca 路径中
                # 从 y 出发至多 dis_xy-half，就是从 x 出发至少 half
                res = g.upto_dis(y, dis_xy - half)
            ans.append(res)
        return ans
```

```java [sol-Java]
class LcaBinaryLifting {
    private final int[] depth;
    public final long[] dis; // 如果是无权树（边权为 1），dis 可以去掉，用 depth 代替
    public final int[][] pa;

    public LcaBinaryLifting(int[][] edges) {
        int n = edges.length + 1;
        int m = 32 - Integer.numberOfLeadingZeros(n); // n 的二进制长度
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].add(new int[]{y, w});
            g[y].add(new int[]{x, w});
        }

        depth = new int[n];
        dis = new long[n];
        pa = new int[n][m];

        dfs(g, 0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                int p = pa[x][i];
                pa[x][i + 1] = p < 0 ? -1 : pa[p][i];
            }
        }
    }

    private void dfs(List<int[]>[] g, int x, int fa) {
        pa[x][0] = fa;
        for (int[] e : g[x]) {
            int y = e[0];
            if (y != fa) {
                depth[y] = depth[x] + 1;
                dis[y] = dis[x] + e[1];
                dfs(g, y, x);
            }
        }
    }

    public int getKthAncestor(int node, int k) {
        for (; k > 0; k &= k - 1) {
            node = pa[node][Integer.numberOfTrailingZeros(k)];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    public int getLCA(int x, int y) {
        if (depth[x] > depth[y]) {
            int tmp = y;
            y = x;
            x = tmp;
        }
        // 使 y 和 x 在同一深度
        y = getKthAncestor(y, depth[y] - depth[x]);
        if (y == x) {
            return x;
        }
        for (int i = pa[x].length - 1; i >= 0; i--) {
            int px = pa[x][i], py = pa[y][i];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[x][0];
    }

    // 返回 x 到 y 的距离（最短路长度）
    public long getDis(int x, int y) {
        return dis[x] + dis[y] - dis[getLCA(x, y)] * 2;
    }

    // 从 x 往上跳【至多】d 距离，返回最远能到达的节点
    public int uptoDis(int x, long d) {
        long dx = dis[x];
        for (int i = pa[x].length - 1; i >= 0; i--) {
            int p = pa[x][i];
            if (p != -1 && dx - dis[p] <= d) { // 可以跳至多 d
                x = p;
            }
        }
        return x;
    }
}

class Solution {
    public int[] findMedian(int n, int[][] edges, int[][] queries) {
        LcaBinaryLifting g = new LcaBinaryLifting(edges);

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int x = queries[i][0];
            int y = queries[i][1];
            if (x == y) {
                ans[i] = x;
                continue;
            }
            int lca = g.getLCA(x, y);
            long disXY = g.dis[x] + g.dis[y] - g.dis[lca] * 2;
            long half = (disXY + 1) / 2;
            if (g.dis[x] - g.dis[lca] >= half) { // 答案在 x-lca 路径中
                // 先往上跳至多 half-1，然后再跳一步，就是至少 half
                int to = g.uptoDis(x, half - 1);
                ans[i] = g.pa[to][0];
            } else { // 答案在 y-lca 路径中
                // 从 y 出发至多 dis_xy-half，就是从 x 出发至少 half
                ans[i] = g.uptoDis(y, disXY - half);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class LcaBinaryLifting {
public:
    vector<int> depth;
    vector<long long> dis;
    vector<vector<int>> pa;

    LcaBinaryLifting(vector<vector<int>>& edges) {
        int n = edges.size() + 1;
        int m = bit_width((unsigned) n); // n 的二进制长度
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }

        depth.resize(n);
        dis.resize(n);
        pa.resize(n, vector<int>(m, -1));

        auto dfs = [&](this auto&& dfs, int x, int fa) -> void {
            pa[x][0] = fa;
            for (auto& [y, w] : g[x]) {
                if (y != fa) {
                    depth[y] = depth[x] + 1;
                    dis[y] = dis[x] + w;
                    dfs(y, x);
                }
            }
        };
        dfs(0, -1);

        for (int i = 0; i < m - 1; i++) {
            for (int x = 0; x < n; x++) {
                if (int p = pa[x][i]; p != -1) {
                    pa[x][i + 1] = pa[p][i];
                }
            }
        }
    }

    int get_kth_ancestor(int node, int k) {
        for (; k; k &= k - 1) {
            node = pa[node][countr_zero((unsigned) k)];
        }
        return node;
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    int get_lca(int x, int y) {
        if (depth[x] > depth[y]) {
            swap(x, y);
        }
        y = get_kth_ancestor(y, depth[y] - depth[x]); // 使 y 和 x 在同一深度
        if (y == x) {
            return x;
        }
        for (int i = pa[x].size() - 1; i >= 0; i--) {
            int px = pa[x][i], py = pa[y][i];
            if (px != py) {
                x = px;
                y = py; // 同时往上跳 2^i 步
            }
        }
        return pa[x][0];
    }

    // 返回 x 到 y 的距离（最短路长度）
    long long get_dis(int x, int y) {
        return dis[x] + dis[y] - dis[get_lca(x, y)] * 2;
    }

    // 从 x 往上跳【至多】d 距离，返回最远能到达的节点
    int upto_dis(int x, long long d) {
        long long dx = dis[x];
        for (int i = pa[x].size() - 1; i >= 0; i--) {
            int p = pa[x][i];
            if (p != -1 && dx - dis[p] <= d) { // 可以跳至多 d
                x = p;
            }
        }
        return x;
    };
};

class Solution {
public:
    vector<int> findMedian(int n, vector<vector<int>>& edges, vector<vector<int>>& queries) {
        LcaBinaryLifting g(edges);

        vector<int> ans;
        ans.reserve(queries.size());
        for (auto& q : queries) {
            int x = q[0], y = q[1];
            if (x == y) {
                ans.push_back(x);
                continue;
            }
            int lca = g.get_lca(x, y);
            long long dis_xy = g.dis[x] + g.dis[y] - g.dis[lca] * 2;
            long long half = (dis_xy + 1) / 2;
            int res;
            if (g.dis[x] - g.dis[lca] >= half) { // 答案在 x-lca 路径中
                // 先往上跳至多 half-1，然后再跳一步，就是至少 half
                int to = g.upto_dis(x, half - 1);
                res = g.pa[to][0];
            } else { // 答案在 y-lca 路径中
                // 从 y 出发至多 dis_xy-half，就是从 x 出发至少 half
                res = g.upto_dis(y, dis_xy - half);
            }
            ans.push_back(res);
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMedian(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	// 17 可以替换成 bits.Len(uint(n))，但数组内存连续性更好
	pa := make([][17]int, n) 
	dep := make([]int, n)
	dis := make([]int, n)

	var dfs func(int, int)
	dfs = func(x, p int) {
		pa[x][0] = p
		for _, e := range g[x] {
			y := e.to
			if y == p {
				continue
			}
			dep[y] = dep[x] + 1
			dis[y] = dis[x] + e.wt
			dfs(y, x)
		}
	}
	dfs(0, -1)

	mx := bits.Len(uint(n))
	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}

	uptoDep := func(x, d int) int {
		for k := uint(dep[x] - d); k > 0; k &= k - 1 {
			x = pa[x][bits.TrailingZeros(k)]
		}
		return x
	}

	// 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
	getLCA := func(x, y int) int {
		if dep[x] > dep[y] {
			x, y = y, x
		}
		y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
		if y == x {
			return x
		}
		for i := mx - 1; i >= 0; i-- {
			px, py := pa[x][i], pa[y][i]
			if px != py {
				x, y = px, py // 同时往上跳 2^i 步
			}
		}
		return pa[x][0]
	}

	// 从 x 往上跳【至多】d 距离，返回最远能到达的节点 
	uptoDis := func(x, d int) int {
		dx := dis[x]
		for i := mx - 1; i >= 0; i-- {
			p := pa[x][i]
			if p != -1 && dx-dis[p] <= d { // 可以跳至多 d
				x = p
			}
		}
		return x
	}

	// 以上是 LCA 模板

	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		if x == y {
			ans[i] = x
			continue
		}
		lca := getLCA(x, y)
		disXY := dis[x] + dis[y] - dis[lca]*2
		half := (disXY + 1) / 2
		if dis[x]-dis[lca] >= half { // 答案在 x-lca 路径中
			// 先往上跳至多 half-1，然后再跳一步，就是至少 half
			to := uptoDis(x, half-1)
			ans[i] = pa[to][0] // 再跳一步
		} else { // 答案在 y-lca 路径中
			// 从 y 出发至多 disXY-half，就是从 x 出发至少 half
			ans[i] = uptoDis(y, disXY-half) 
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。返回值不计入。

## 相似题目

见下面树题单的「**§3.8 倍增算法、最近公共祖先（LCA）**」。

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
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
