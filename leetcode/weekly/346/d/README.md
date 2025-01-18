## 视频讲解

见[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第四题，欢迎点赞投币！

## 什么时候无解？

题目要求把边权为 $-1$ 的至少修改为 $1$。如果都修改成 $1$，跑最短路（Dijkstra），发现从起点到终点的最短路长度大于 $\textit{target}$，那么由于边权变大，最短路不可能变小，所以此时无解。

另外一种无解的情况是，如果都修改成无穷大（本题限制为 $2\cdot 10^9$），发现从起点到终点的最短路长度小于 $\textit{target}$，那么由于边权变小，最短路不可能变大，所以此时也无解。

## 一个错误的思路

先把 $-1$ 都修改成 $1$，然后跑 Dijkstra。设从起点到终点的最短路长度为 $d$。

如果 $d<\textit{target}$，**看上去**把其中一个 $1$ 加上 $\textit{target}-d$，就可以使从起点到终点的最短路长度恰好为 $\textit{target}$ 了。

这是不对的，因为在增加边权后，最短路可能就走别的边了，不走刚才修改的这条边了。（具体见视频中画的例子。）

只修改一条边不行，那么修改两条边呢？也是不行的，最短路仍然可以走别的边。

## 正确思路

先把 $-1$ 都修改成 $1$，然后跑第一遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,0}$。

如果从起点到终点的最短路长度不超过 $\textit{target}$（否则无解，返回空数组），那我们就来尝试修改某些边权，使得从起点到终点的最短路长度恰好等于 $\textit{target}$。

这会引出两个问题：

1. 要按照什么样的**顺序**修改这些边？
2. 修改成多少合适？

正所谓「牵一发而动全身」，从上面对错误思路的分析可知，仅仅修改一个边权，就可能影响很多最短路的值。

那不妨再跑一遍 Dijkstra，由于 Dijkstra 算法保证每次拿到的点的最短路就是最终的最短路，所以按照 Dijkstra 算法遍历点/边的顺序去修改，就不会对**已确定的**最短路产生影响。

对于第二遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,1}$。

对于一条可以修改的边 $x-y$，假设要把它的边权改为 $w$，那么 $\textit{source}-x-y-\textit{destination}$ 这条路径由三部分组成：

1. 从 $\textit{source}$ 到 $x$ 的最短路，这是第二遍 Dijkstra 算出来的，即 $d_{x,1}$。
2. 从 $x$ 到 $y$，即 $w$。
3. 从 $y$ 到 $\textit{destination}$ 的最短路，由于后面的边还没有修改，这个最短路是第一遍 Dijkstra 算出来的，即 $d_{\textit{destination},0} - d_{y,0}$。

这三部分之和需要等于 $\textit{target}$，所以有

$$
d_{x,1} + w + d_{\textit{destination},0} - d_{y,0} = \textit{target}
$$

解得

$$
w = \textit{target} - d_{\textit{destination},0} + d_{y,0} - d_{x,1}
$$

> 注意上式中的 $\textit{target} - d_{\textit{destination},0}$ 是一个定值，代码中用 $\textit{delta}$ 表示。

根据「什么时候无解」中的分析，如果第二遍 Dijkstra 跑完后，从起点到终点的最短路仍然小于 $\textit{target}$，那么就说明无法修改，返回空数组。

否则，答案就是我们在第二遍 Dijkstra 中做出的修改。注意第二遍 Dijkstra 跑完后可能还有些边是 $-1$（因为在 $w=1$ 的时候没有修改，或者有些边不影响最短路），把这些边都改成 $1$ 就行。

代码实现时，为了修改边权，需要在邻接表中额外记录边的编号。

此外，由于输入最坏是稠密图，可以用朴素版的 Dijkstra 算法。

```py [sol-Python3]
class Solution:
    def modifiedGraphEdges(self, n: int, edges: List[List[int]], source: int, destination: int, target: int) -> List[List[int]]:
        g = [[] for _ in range(n)]
        for i, (x, y, _) in enumerate(edges):
            g[x].append((y, i))
            g[y].append((x, i))  # 建图，额外保存边的编号

        dis = [[inf, inf] for _ in range(n)]
        dis[source] = [0, 0]

        def dijkstra(k: int) -> None:  # 这里 k 表示第一次/第二次
            vis = [False] * n
            while True:
                # 找到当前最短路，去更新它的邻居的最短路
                # 根据数学归纳法，dis[x][k] 一定是最短路长度
                x = -1
                for y, (b, d) in enumerate(zip(vis, dis)):
                    if not b and (x < 0 or d[k] < dis[x][k]):
                        x = y
                if x == destination:  # 起点 source 到终点 destination 的最短路已确定
                    return
                vis[x] = True  # 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
                for y, eid in g[x]:
                    wt = edges[eid][2]
                    if wt == -1:
                        wt = 1  # -1 改成 1
                    if k == 1 and edges[eid][2] == -1:
                        # 第二次 Dijkstra，改成 w
                        w = delta + dis[y][0] - dis[x][1]
                        if w > wt:
                            edges[eid][2] = wt = w  # 直接在 edges 上修改
                    # 更新最短路
                    dis[y][k] = min(dis[y][k], dis[x][k] + wt)

        dijkstra(0)
        delta = target - dis[destination][0]
        if delta < 0:  # -1 全改为 1 时，最短路比 target 还大
            return []

        dijkstra(1)
        if dis[destination][1] < target:  # 最短路无法再变大，无法达到 target
            return []

        for e in edges:
            if e[2] == -1:  # 剩余没修改的边全部改成 1
                e[2] = 1
        return edges
```

```java [sol-Java]
class Solution {
    public int[][] modifiedGraphEdges(int n, int[][] edges, int source, int destination, int target) {
        List<int[]> g[] = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (int i = 0; i < edges.length; i++) {
            int x = edges[i][0], y = edges[i][1];
            g[x].add(new int[]{y, i});
            g[y].add(new int[]{x, i}); // 建图，额外记录边的编号
        }

        var dis = new int[n][2];
        for (int i = 0; i < n; i++)
            if (i != source)
                dis[i][0] = dis[i][1] = Integer.MAX_VALUE / 2;

        dijkstra(g, edges, destination, dis, 0, 0);
        int delta = target - dis[destination][0];
        if (delta < 0) // -1 全改为 1 时，最短路比 target 还大
            return new int[][]{};

        dijkstra(g, edges, destination, dis, delta, 1);
        if (dis[destination][1] < target) // 最短路无法再变大，无法达到 target
            return new int[][]{};

        for (int[] e : edges)
            if (e[2] == -1) // 剩余没修改的边全部改成 1
                e[2] = 1;
        return edges;
    }

    // 朴素 Dijkstra 算法
    // 这里 k 表示第一次/第二次
    private void dijkstra(List<int[]> g[], int[][] edges, int destination, int[][] dis, int delta, int k) {
        int n = g.length;
        boolean[] vis = new boolean[n];
        for (; ; ) {
            // 找到当前最短路，去更新它的邻居的最短路
            // 根据数学归纳法，dis[x][k] 一定是最短路长度
            int x = -1;
            for (int i = 0; i < n; ++i)
                if (!vis[i] && (x < 0 || dis[i][k] < dis[x][k]))
                    x = i;
            if (x == destination) // 起点 source 到终点 destination 的最短路已确定
                return;
            vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for (int[] e : g[x]) {
                int y = e[0], eid = e[1];
                int wt = edges[eid][2];
                if (wt == -1)
                    wt = 1; // -1 改成 1
                if (k == 1 && edges[eid][2] == -1) {
                    // 第二次 Dijkstra，改成 w
                    int w = delta + dis[y][0] - dis[x][1];
                    if (w > wt)
                        edges[eid][2] = wt = w; // 直接在 edges 上修改
                }
                // 更新最短路
                dis[y][k] = Math.min(dis[y][k], dis[x][k] + wt);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> modifiedGraphEdges(int n, vector<vector<int>> &edges, int source, int destination, int target) {
        vector<pair<int, int>> g[n];
        for (int i = 0; i < edges.size(); i++) {
            int x = edges[i][0], y = edges[i][1];
            g[x].emplace_back(y, i);
            g[y].emplace_back(x, i); // 建图，额外记录边的编号
        }

        int dis[n][2], delta, vis[n];
        memset(dis, 0x3f, sizeof(dis));
        dis[source][0] = dis[source][1] = 0;
        auto dijkstra = [&](int k) { // 这里 k 表示第一次/第二次
            memset(vis, 0, sizeof(vis));
            for (;;) {
                // 找到当前最短路，去更新它的邻居的最短路
                // 根据数学归纳法，dis[x][k] 一定是最短路长度
                int x = -1;
                for (int i = 0; i < n; ++i)
                    if (!vis[i] && (x < 0 || dis[i][k] < dis[x][k]))
                        x = i;
                if (x == destination) // 起点 source 到终点 destination 的最短路已确定
                    return;
                vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
                for (auto [y, eid]: g[x]) {
                    int wt = edges[eid][2];
                    if (wt == -1)
                        wt = 1; // -1 改成 1
                    if (k == 1 && edges[eid][2] == -1) {
                        // 第二次 Dijkstra，改成 w
                        int w = delta + dis[y][0] - dis[x][1];
                        if (w > wt)
                            edges[eid][2] = wt = w; // 直接在 edges 上修改
                    }
                    // 更新最短路
                    dis[y][k] = min(dis[y][k], dis[x][k] + wt);
                }
            }
        };

        dijkstra(0);
        delta = target - dis[destination][0];
        if (delta < 0) // -1 全改为 1 时，最短路比 target 还大
            return {};

        dijkstra(1);
        if (dis[destination][1] < target) // 最短路无法再变大，无法达到 target
            return {};

        for (auto &e: edges)
            if (e[2] == -1) // 剩余没修改的边全部改成 1
                e[2] = 1;
        return edges;
    }
};
```

```go [sol-Go]
func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i}) // 建图，额外记录边的编号
	}

	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i][0] = math.MaxInt32 / 2
		dis[i][1] = math.MaxInt32 / 2
	}
	dis[source] = [2]int{}
	dijkstra := func(k int) { // 这里 k 表示第一次/第二次
		vis := make([]bool, n)
		for {
			// 找到当前最短路，去更新它的邻居的最短路
			// 根据数学归纳法，dis[x][k] 一定是最短路长度
			x := -1
			for y, b := range vis {
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination { // 起点 source 到终点 destination 的最短路已确定
				return
			}
			vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
			for _, e := range g[x] {
				y, wt := e.to, edges[e.eid][2]
				if wt == -1 {
					wt = 1 // -1 改成 1
				}
				if k == 1 && edges[e.eid][2] == -1 {
					// 第二次 Dijkstra，改成 w
					w := delta + dis[y][0] - dis[x][1]
					if w > wt {
						wt = w
						edges[e.eid][2] = w // 直接在 edges 上修改
					}
				}
				// 更新最短路
				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
			}
		}
	}

	dijkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大
		return nil
	}

	dijkstra(1)
	if dis[destination][1] < target { // 最短路无法再变大，无法达到 target
		return nil
	}

	for _, e := range edges {
		if e[2] == -1 { // 剩余没修改的边全部改成 1
			e[2] = 1
		}
	}
	return edges
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$。在稠密图（本题最坏情况）中，算法的时间复杂度与边的数量 $m=\mathcal{O}(n^2)$ 成正比。
- 空间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{edges}$ 的长度。注意输入是连通图，$m$ 至少为 $n-1$，所以 $\mathcal{O}(n+m)=\mathcal{O}(m)$

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
