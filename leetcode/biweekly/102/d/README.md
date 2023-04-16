下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，**包括 Dijkstra 算法的原理及实现**。

---

Dijkstra 模板题，原理如下。

定义 $\textit{start}$ 为起点，$\textit{dis}[i]$ 表示从 $\textit{start}$ 到 $i$ 的最短路的长度。初始时 $\textit{dis}[\textit{start}]=0$，其余 $\textit{dis}[i]$ 为 $\infty$。

首先，从 $\textit{start}$ 出发，更新邻居的最短路。

下一步，寻找除去 $\textit{start}$ 的 $\textit{dis}$ 的最小值，设这个点为 $x$，那么 $\textit{dis}[x]$ 就已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

证明：由于除去起点的其余 $\textit{dis}[i]$ 都不低于 $\textit{dis}[x]$，且图中边权都非负，那么从另外一个点 $y$ 去更新 $\textit{dis}[x]$ 时，是无法让 $\textit{dis}[x]$ 变得更小的（因为 $\textit{dis}[x]$ 是当前最小），因此 $\textit{dis}[x]$ 已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

由于在寻找 $\textit{dis}$ 的最小值时，需要排除在前面的循环中找到的 $x$（因为已经更新 $x$ 到其它点的最短路了，无需反复更新），可以用一个 $\textit{vis}$ 数组标记这些 $x$。

以上，通过**数学归纳法**，可以证明每次找到的未被标记的 $\textit{dis}$ 的最小值就是最短路。

由于输入的图最坏是**稠密图**，所以采用邻接矩阵实现。

```py [sol1-Python3]
class Graph:
    def __init__(self, n: int, edges: List[List[int]]):
        g = [[inf] * n for _ in range(n)]  # 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        for x, y, w in edges:
            g[x][y] = w  # 添加一条边（输入保证没有重边）
        self.g = g

    def addEdge(self, e: List[int]) -> None:
        self.g[e[0]][e[1]] = e[2]  # 添加一条边（输入保证这条边之前不存在）

    def shortestPath(self, node1: int, node2: int) -> int:
        ans = self._dijkstra(node1)[node2]
        return ans if ans < inf else -1

    # Dijkstra 算法的邻接矩阵版本
    # 返回从 start 出发，到各个点的最短路，如果不存在则为无穷大
    def _dijkstra(self, start: int) -> List[int]:
        n = len(self.g)
        dis = [inf] * n
        dis[start] = 0
        vis = [False] * n
        while True:
            # 找到当前最短路，去更新它的邻居的最短路
            # 根据数学归纳法，dis[x] 一定是最短路长度
            x = -1
            for i, (b, d) in enumerate(zip(vis, dis)):
                if not b and (x < 0 or d < dis[x]):
                    x = i
            if x < 0 or dis[x] == inf:  # 所有从 start 能到达的点都被更新了
                return dis
            vis[x] = True  # 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for y, w in enumerate(self.g[x]):
                if dis[x] + w < dis[y]:
                    dis[y] = dis[x] + w  # 更新最短路长度
```

```java [sol1-Java]
class Graph {
    private int[][] g;

    public Graph(int n, int[][] edges) {
        g = new int[n][n]; // 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        for (int i = 0; i < n; ++i)
            Arrays.fill(g[i], Integer.MAX_VALUE);
        for (var e : edges)
            g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边）
    }

    public void addEdge(int[] e) {
        g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证这条边之前不存在）
    }

    // Dijkstra 算法的邻接矩阵版本
    public int shortestPath(int start, int end) {
        int n = g.length;
        var dis = new int[n];
        Arrays.fill(dis, Integer.MAX_VALUE);
        dis[start] = 0;
        var vis = new boolean[n];
        for (;;) {
            // 找到当前最短路，去更新它的邻居的最短路
            // 根据数学归纳法，dis[x] 一定是最短路长度
            int x = -1;
            for (int i = 0; i < n; ++i)
                if (!vis[i] && (x < 0 || dis[i] < dis[x]))
                    x = i;
            if (x < 0 || dis[x] == Integer.MAX_VALUE) // 所有从 start 能到达的点都被更新了
                return dis[end] < Integer.MAX_VALUE ? dis[end] : -1;
            vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for (int y = 0; y < n; ++y)
                if (g[x][y] < Integer.MAX_VALUE && dis[x] + g[x][y] < dis[y])
                    dis[y] = dis[x] + g[x][y]; // 更新最短路长度
        }
    }
}
```

```cpp [sol1-C++]
class Graph {
    vector<vector<int>> g;
public:
    Graph(int n, vector<vector<int>> &edges) {
        // 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        g = vector<vector<int>>(n, vector<int>(n, INT_MAX));
        for (auto &e: edges)
            g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边）
    }

    void addEdge(vector<int> e) {
        g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证这条边之前不存在）
    }

    // Dijkstra 算法的邻接矩阵版本
    int shortestPath(int start, int end) {
        int n = g.size();
        vector<int> dis(n, INT_MAX), vis(n);
        dis[start] = 0;
        for (;;) {
            // 找到当前最短路，去更新它的邻居的最短路
            // 根据数学归纳法，dis[x] 一定是最短路长度
            int x = -1;
            for (int i = 0; i < n; ++i)
                if (!vis[i] && (x < 0 || dis[i] < dis[x]))
                    x = i;
            if (x < 0 || dis[x] == INT_MAX) // 所有从 start 能到达的点都被更新了
                return dis[end] < INT_MAX ? dis[end] : -1;
            vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for (int y = 0; y < n; ++y)
                if (g[x][y] < INT_MAX && dis[x] + g[x][y] < dis[y])
                    dis[y] = dis[x] + g[x][y]; // 更新最短路长度
        }
    }
};
```

```go [sol1-Go]
type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（输入保证这条边之前不存在）
}

func (g Graph) ShortestPath(node1, node2 int) int {
	ans := g.dijkstra(node1)[node2]
	if ans < math.MaxInt {
		return ans
	}
	return -1
}

// Dijkstra 算法的邻接矩阵版本
// 返回从 start 出发，到各个点的最短路，如果不存在则为无穷大
func (g Graph) dijkstra(start int) []int {
	n := len(g)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == math.MaxInt { // 所有从 start 能到达的点都被更新了
			return dis
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			if w < math.MaxInt && dis[x]+w < dis[y] {
				dis[y] = dis[x] + w // 更新最短路长度
			}
		}
	}
}
```

### 复杂度分析

- 时间复杂度：$O(qn^2)$，其中 $q$ 为 $\texttt{shortestPath}$ 的调用次数。每次求最短路的时间复杂度为 $O(n^2)$，在本题的输入下，这比堆的实现要快。
- 空间复杂度：$O(n^2)$。
