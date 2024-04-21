**前置知识**：[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)

首先用 Dijkstra 算法（堆优化版本）计算出起点 $0$ 到所有节点的最短路长度 $\textit{dis}$。

如果 $\textit{dis}[n-1]=\infty$，说明无法从起点 $0$ 到终点 $n-1$，答案全为 $\textit{false}$。

否则，我们可以从终点 $n-1$ 出发，倒着 DFS 或 BFS，设当前在点 $y$，邻居为 $x$，边权为 $w$，如果满足

$$
\textit{dis}[x] + w = \textit{dis}[y]
$$

则说明 $x\textit{-}y$ 这条边在从 $0$ 到 $n-1$ 的最短路上。

### 答疑

**问**：为什么在求出最短路后，不能从起点 $0$ 出发去寻找在最短路上的边？

**答**：从起点 $0$ 出发，当发现 $\textit{dis}[x] + w = \textit{dis}[y]$ 时，这仅仅意味着 $x\textit{-}y$ 这条边在从 $0$ 出发的最短路上，但这条最短路不一定通向终点 $n-1$（比如图是一棵树，这条边不通往 $n-1$）。而从终点 $n-1$ 出发倒着寻找，就能保证符合等式的边在通向终点的最短路上。

请看 [视频讲解](https://www.bilibili.com/video/BV1gu4m1F7B8/) 第四题，欢迎点赞关注！

## 写法一：Dijkstra 算法 + DFS

```py [sol-Python3]
class Solution:
    def findAnswer(self, n: int, edges: List[List[int]]) -> List[bool]:
        g = [[] for _ in range(n)]
        for i, (x, y, w) in enumerate(edges):
            g[x].append((y, w, i))
            g[y].append((x, w, i))

        # Dijkstra 算法模板
        dis = [inf] * n
        dis[0] = 0
        h = [(0, 0)]
        while h:
            dx, x = heappop(h)
            if dx > dis[x]:
                continue
            for y, w, _ in g[x]:
                new_dis = dx + w
                if new_dis < dis[y]:
                    dis[y] = new_dis
                    heappush(h, (new_dis, y))

        ans = [False] * len(edges)
        # 图不连通
        if dis[-1] == inf:
            return ans

        # 从终点出发 DFS
        vis = [False] * n
        def dfs(y: int) -> None:
            vis[y] = True
            for x, w, i in g[y]:
                if dis[x] + w != dis[y]:
                    continue
                ans[i] = True
                if not vis[x]:
                    dfs(x)
        dfs(n - 1)
        return ans
```

```java [sol-Java]
class Solution {
    public boolean[] findAnswer(int n, int[][] edges) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 0; i < edges.length; i++) {
            int[] e = edges[i];
            int x = e[0], y = e[1], w = e[2];
            g[x].add(new int[]{y, w, i});
            g[y].add(new int[]{x, w, i});
        }

        long[] dis = new long[n];
        Arrays.fill(dis, Long.MAX_VALUE);
        dis[0] = 0;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        pq.offer(new long[]{0, 0});
        while (!pq.isEmpty()) {
            long[] dxPair = pq.poll();
            long dx = dxPair[0];
            int x = (int) dxPair[1];
            if (dx > dis[x]) {
                continue;
            }
            for (int[] t : g[x]) {
                int y = t[0];
                int w = t[1];
                long newDis = dx + w;
                if (newDis < dis[y]) {
                    dis[y] = newDis;
                    pq.offer(new long[]{newDis, y});
                }
            }
        }

        boolean[] ans = new boolean[edges.length];
        // 图不连通
        if (dis[n - 1] == Long.MAX_VALUE) {
            return ans;
        }

        // 从终点出发 DFS
        boolean[] vis = new boolean[n];
        dfs(n - 1, g, dis, ans, vis);
        return ans;
    }

    private void dfs(int y, List<int[]>[] g, long[] dis, boolean[] ans, boolean[] vis) {
        vis[y] = true;
        for (int[] t : g[y]) {
            int x = t[0];
            int w = t[1];
            int i = t[2];
            if (dis[x] + w != dis[y]) {
                continue;
            }
            ans[i] = true;
            if (!vis[x]) {
                dfs(x, g, dis, ans, vis);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> findAnswer(int n, vector<vector<int>>& edges) {
        vector<vector<tuple<int, int, int>>> g(n);
        for (int i = 0; i < edges.size(); i++) {
            auto& e = edges[i];
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w, i);
            g[y].emplace_back(x, w, i);
        }

        vector<long long> dis(n, LLONG_MAX);
        dis[0] = 0;
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        pq.emplace(0, 0);
        while (!pq.empty()) {
            auto [dx, x] = pq.top();
            pq.pop();
            if (dx > dis[x]) {
                continue;
            }
            for (auto [y, w, _] : g[x]) {
                int new_dis = dx + w;
                if (new_dis < dis[y]) {
                    dis[y] = new_dis;
                    pq.emplace(new_dis, y);
                }
            }
        }

        vector<bool> ans(edges.size());
        // 图不连通
        if (dis[n - 1] == LLONG_MAX) {
            return ans;
        }

        // 从终点出发 DFS
        vector<int> vis(n);
        function<void(int)> dfs = [&](int y) {
            vis[y] = true;
            for (auto [x, w, i] : g[y]) {
                if (dis[x] + w != dis[y]) {
                    continue;
                }
                ans[i] = true;
                if (!vis[x]) {
                    dfs(x);
                }
            }
        };
        dfs(n - 1);
        return ans;
    }
};
```

```go [sol-Go]
func findAnswer(n int, edges [][]int) []bool {
	type edge struct{ to, w, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w, i})
		g[y] = append(g[y], edge{x, w, i})
	}

	// Dijkstra 算法模板
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := p.dis + e.w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := make([]bool, len(edges))
	// 图不连通
	if dis[n-1] == math.MaxInt {
		return ans
	}

	// 从终点出发 DFS
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(y int) {
		vis[y] = true
		for _, e := range g[y] {
			x := e.to
			if dis[x]+e.w != dis[y] {
				continue
			}
			ans[e.i] = true
			if !vis[x] {
				dfs(x)
			}
		}
	}
	dfs(n - 1)
	return ans
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 写法二：Dijkstra 算法 + BFS

```py [sol-Python3]
class Solution:
    def findAnswer(self, n: int, edges: List[List[int]]) -> List[bool]:
        g = [[] for _ in range(n)]
        for i, (x, y, w) in enumerate(edges):
            g[x].append((y, w, i))
            g[y].append((x, w, i))

        # Dijkstra 算法模板
        dis = [inf] * n
        dis[0] = 0
        h = [(0, 0)]
        while h:
            dx, x = heappop(h)
            if dx > dis[x]:
                continue
            for y, w, _ in g[x]:
                new_dis = dx + w
                if new_dis < dis[y]:
                    dis[y] = new_dis
                    heappush(h, (new_dis, y))

        ans = [False] * len(edges)
        # 图不连通
        if dis[-1] == inf:
            return ans

        # 从终点出发 BFS
        vis = [False] * n
        vis[-1] = True
        q = deque([n - 1])
        while q:
            y = q.popleft()
            for x, w, i in g[y]:
                if dis[x] + w != dis[y]:
                    continue
                ans[i] = True
                if not vis[x]:
                    vis[x] = True
                    q.append(x)
        return ans
```

```java [sol-Java]
class Solution {
    public boolean[] findAnswer(int n, int[][] edges) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 0; i < edges.length; i++) {
            int[] e = edges[i];
            int x = e[0], y = e[1], w = e[2];
            g[x].add(new int[]{y, w, i});
            g[y].add(new int[]{x, w, i});
        }

        long[] dis = new long[n];
        Arrays.fill(dis, Long.MAX_VALUE);
        dis[0] = 0;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        pq.offer(new long[]{0, 0});
        while (!pq.isEmpty()) {
            long[] dxPair = pq.poll();
            long dx = dxPair[0];
            int x = (int) dxPair[1];
            if (dx > dis[x]) {
                continue;
            }
            for (int[] t : g[x]) {
                int y = t[0];
                int w = t[1];
                long newDis = dx + w;
                if (newDis < dis[y]) {
                    dis[y] = newDis;
                    pq.offer(new long[]{newDis, y});
                }
            }
        }

        boolean[] ans = new boolean[edges.length];
        // 图不连通
        if (dis[n - 1] == Long.MAX_VALUE) {
            return ans;
        }

        // 从终点出发 BFS
        boolean[] vis = new boolean[n];
        vis[n - 1] = true;
        Queue<Integer> q = new ArrayDeque<>();
        q.offer(n - 1);
        while (!q.isEmpty()) {
            int y = q.poll();
            for (int[] t : g[y]) {
                int x = t[0];
                int w = t[1];
                int i = t[2];
                if (dis[x] + w != dis[y]) {
                    continue;
                }
                ans[i] = true;
                if (!vis[x]) {
                    vis[x] = true;
                    q.offer(x);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> findAnswer(int n, vector<vector<int>>& edges) {
        vector<vector<tuple<int, int, int>>> g(n);
        for (int i = 0; i < edges.size(); i++) {
            auto& e = edges[i];
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w, i);
            g[y].emplace_back(x, w, i);
        }

        vector<long long> dis(n, LLONG_MAX);
        dis[0] = 0;
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        pq.emplace(0, 0);
        while (!pq.empty()) {
            auto [dx, x] = pq.top();
            pq.pop();
            if (dx > dis[x]) {
                continue;
            }
            for (auto [y, w, _] : g[x]) {
                int new_dis = dx + w;
                if (new_dis < dis[y]) {
                    dis[y] = new_dis;
                    pq.emplace(new_dis, y);
                }
            }
        }

        vector<bool> ans(edges.size());
        // 图不连通
        if (dis[n - 1] == LLONG_MAX) {
            return ans;
        }

        // 从终点出发 BFS
        vector<int> vis(n);
        vis[n - 1] = true;
        queue<int> q;
        q.push(n - 1);
        while (!q.empty()) {
            int y = q.front();
            q.pop();
            for (auto [x, w, i] : g[y]) {
                if (dis[x] + w != dis[y]) {
                    continue;
                }
                ans[i] = true;
                if (!vis[x]) {
                    vis[x] = true;
                    q.push(x);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findAnswer(n int, edges [][]int) []bool {
	type edge struct{ to, w, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w, i})
		g[y] = append(g[y], edge{x, w, i})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := p.dis + e.w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := make([]bool, len(edges))
	// 图不连通
	if dis[n-1] == math.MaxInt {
		return ans
	}

	// 从终点出发 BFS
	vis := make([]bool, n)
	vis[n-1] = true
	q := []int{n - 1}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, e := range g[y] {
			x := e.to
			if dis[x]+e.w != dis[y] {
				continue
			}
			ans[e.i] = true
			if !vis[x] {
				vis[x] = true
				q = append(q, x)
			}
		}
	}
	return ans
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
