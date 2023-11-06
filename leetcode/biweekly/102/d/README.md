## 本题视频讲解

见[【双周赛 102】](https://www.bilibili.com/video/BV1Es4y1N7v1/) 第四题。

包含 Dijkstra 和 Floyd 的原理及实现。

## 方法一：Dijkstra

定义 $\textit{start}$ 为起点，$\textit{dis}[i]$ 表示从 $\textit{start}$ 到 $i$ 的最短路的长度。初始时 $\textit{dis}[\textit{start}]=0$，其余 $\textit{dis}[i]$ 为 $\infty$。

首先，从 $\textit{start}$ 出发，更新邻居的最短路。

下一步，寻找除去 $\textit{start}$ 的 $\textit{dis}$ 的最小值，设这个点为 $x$，那么 $\textit{dis}[x]$ 就已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

证明：由于除去起点的其余 $\textit{dis}[i]$ 都不低于 $\textit{dis}[x]$，且图中边权都非负，那么从另外一个点 $y$ 去更新 $\textit{dis}[x]$ 时，是无法让 $\textit{dis}[x]$ 变得更小的（因为 $\textit{dis}[x]$ 是当前最小），因此 $\textit{dis}[x]$ 已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

由于在寻找 $\textit{dis}$ 的最小值时，需要排除在前面的循环中找到的 $x$（因为已经更新 $x$ 到其它点的最短路了，无需反复更新），可以用一个 $\textit{vis}$ 数组标记这些 $x$。

以上，通过**数学归纳法**，可以证明每次找到的未被标记的 $\textit{dis}$ 的最小值就是最短路。

由于输入的图最坏情况下是**稠密图**，所以采用邻接矩阵实现。

```py [sol1-Python3]
class Graph:
    def __init__(self, n: int, edges: List[List[int]]):
        g = [[inf] * n for _ in range(n)]  # 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        for x, y, w in edges:
            g[x][y] = w  # 添加一条边（输入保证没有重边）
        self.g = g

    def addEdge(self, e: List[int]) -> None:
        self.g[e[0]][e[1]] = e[2]  # 添加一条边（输入保证这条边之前不存在）

    # 朴素 Dijkstra 算法
    def shortestPath(self, start: int, end: int) -> int:
        n = len(self.g)
        dis = [inf] * n  # 从 start 出发，到各个点的最短路，如果不存在则为无穷大
        dis[start] = 0
        vis = [False] * n
        while True:  # 至多循环 n 次
            # 找到当前最短路，去更新它的邻居的最短路
            # 根据数学归纳法，dis[x] 一定是最短路长度
            x = -1
            for i, (b, d) in enumerate(zip(vis, dis)):
                if not b and (x < 0 or d < dis[x]):
                    x = i
            if x < 0 or dis[x] == inf:  # 所有从 start 能到达的点都被更新了
                return -1
            if x == end:  # 找到终点，提前退出
                return dis[x]
            vis[x] = True  # 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for y, w in enumerate(self.g[x]):
                if dis[x] + w < dis[y]:
                    dis[y] = dis[x] + w  # 更新最短路长度
```

```java [sol1-Java]
class Graph {
    private static final int INF = Integer.MAX_VALUE / 2; // 防止更新最短路时加法溢出

    private int[][] g;

    public Graph(int n, int[][] edges) {
        g = new int[n][n]; // 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        for (int i = 0; i < n; ++i)
            Arrays.fill(g[i], INF);
        for (var e : edges)
            g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边）
    }

    public void addEdge(int[] e) {
        g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证这条边之前不存在）
    }

    // 朴素 Dijkstra 算法
    public int shortestPath(int start, int end) {
        int n = g.length;
        var dis = new int[n]; // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
        Arrays.fill(dis, INF);
        dis[start] = 0;
        var vis = new boolean[n];
        for (;;) {
            // 找到当前最短路，去更新它的邻居的最短路
            // 根据数学归纳法，dis[x] 一定是最短路长度
            int x = -1;
            for (int i = 0; i < n; ++i)
                if (!vis[i] && (x < 0 || dis[i] < dis[x]))
                    x = i;
            if (x < 0 || dis[x] == INF) // 所有从 start 能到达的点都被更新了
                return -1;
            if (x == end) // 找到终点，提前退出
                return dis[x];
            vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for (int y = 0; y < n; ++y)
                dis[y] = Math.min(dis[y], dis[x] + g[x][y]); // 更新最短路长度
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
        g = vector<vector<int>>(n, vector<int>(n, INT_MAX / 2));
        for (auto &e: edges)
            g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边）
    }

    void addEdge(vector<int> e) {
        g[e[0]][e[1]] = e[2]; // 添加一条边（输入保证这条边之前不存在）
    }

    // 朴素 Dijkstra 算法
    int shortestPath(int start, int end) {
        int n = g.size();
        vector<int> dis(n, INT_MAX / 2), vis(n);
        dis[start] = 0;
        for (;;) {
            // 找到当前最短路，去更新它的邻居的最短路
            // 根据数学归纳法，dis[x] 一定是最短路长度
            int x = -1;
            for (int i = 0; i < n; ++i)
                if (!vis[i] && (x < 0 || dis[i] < dis[x]))
                    x = i;
            if (x < 0 || dis[x] == INT_MAX / 2) // 所有从 start 能到达的点都被更新了
                return -1;
            if (x == end) // 找到终点，提前退出
                return dis[x];
            vis[x] = true; // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
            for (int y = 0; y < n; ++y)
                dis[y] = min(dis[y], dis[x] + g[x][y]); // 更新最短路长度
        }
    }
};
```

```go [sol1-Go]
const inf = math.MaxInt / 2 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
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

// 朴素 Dijkstra 算法
func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(qn^2)$，其中 $q$ 为 $\texttt{shortestPath}$ 的调用次数。每次求最短路的时间复杂度为 $O(n^2)$，在本题的输入下，这比堆的实现要快。
- 空间复杂度：$O(n^2)$。

## 方法二：Floyd

Floyd 本质是动态规划。由于这个动态规划的状态定义不是很好想出来，所以我就直接描述算法了：

定义 $d[k][i][j]$ 表示从 $i$ 到 $j$ 的最短路长度，并且从 $i$ 到 $j$ 的路径上的中间节点（不含 $i$ 和 $j$）的编号至多为 $k$。

分类讨论：

- 如果 $i$ 到 $j$ 的路径上的节点编号没有 $k$，那么按照定义 $d[k][i][j] = d[k-1][i][j]$。
- 如果 $i$ 到 $j$ 的路径上的节点编号有 $k$，那么可以视作先从 $i$ 到 $k$，再从 $k$ 到 $j$。由于 $i$ 到 $k$ 和 $k$ 到 $j$ 的中间节点都没有 $k$，所以有 $d[k][i][j] = d[k-1][i][k] + d[k-1][k][j]$。

取最小值，得 

$$
d[k][i][j] =\min(d[k-1][i][j], d[k-1][i][k] + d[k-1][k][j])
$$

初始值 $d[0][i][j]$ 为原图中 $i$ 到 $j$ 的边权，如果不存在则为 $\infty$。最终 $i$ 到 $j$ 的最短路长度为 $d[n-1][i][j]$。

代码实现时，第一个维度可以优化掉，即

$$
d[i][j] =\min(d[[i][j], d[i][k] + d[k][j])
$$

> 为什么这样做是对的？视频中讲了正确性。注意 $f[k][i][k]$ 和 $f[k-1][i][k]$ 是一样的。

对于 $\texttt{addEdge}$ 操作，记 $x=\textit{from},y=\textit{to}$。如果 $\textit{edgeCost} \ge d[x][y]$，则无法更新任何点对的最短路。否则枚举所有 $d[i][j]$，尝试看看能否更新成更小，即 $i---x-y---j$ 是否更短：

$$
d[i][j] = \min(d[i][j], d[i][x] + \textit{edgeCode} + d[y][j])
$$

由于当 $i=x$ 或 $j=y$ 时，需要用到 $d[i][i]$ 这样的值，所以初始化的时候，$d[i][i]$ 要置为 $0$。

```py [sol2-Python3]
class Graph:
    def __init__(self, n: int, edges: List[List[int]]):
        d = [[inf] * n for _ in range(n)]
        for i in range(n):
            d[i][i] = 0
        for x, y, w in edges:
            d[x][y] = w  # 添加一条边（输入保证没有重边和自环）
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    d[i][j] = min(d[i][j], d[i][k] + d[k][j])
        self.d = d

    def addEdge(self, e: List[int]) -> None:
        d = self.d
        n = len(d)
        x, y, w = e
        if w >= d[x][y]:  # 无需更新
            return
        for i in range(n):
            for j in range(n):
                d[i][j] = min(d[i][j], d[i][x] + w + d[y][j])

    def shortestPath(self, start: int, end: int) -> int:
        ans = self.d[start][end]
        return ans if ans < inf else -1
```

```java [sol2-Java]
class Graph {
    private static final int INF = Integer.MAX_VALUE / 3; // 防止更新最短路时加法溢出

    private int[][] d;

    public Graph(int n, int[][] edges) {
        d = new int[n][n]; // 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        for (int i = 0; i < n; ++i) {
            Arrays.fill(d[i], INF);
            d[i][i] = 0;
        }
        for (var e : edges)
            d[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边和自环）
        for (int k = 0; k < n; ++k)
            for (int i = 0; i < n; ++i)
                for (int j = 0; j < n; ++j)
                    d[i][j] = Math.min(d[i][j], d[i][k] + d[k][j]);
    }

    public void addEdge(int[] e) {
        int x = e[0], y = e[1], w = e[2], n = d.length;
        if (w >= d[x][y]) // 无需更新
            return;
        for (int i = 0; i < n; ++i)
            for (int j = 0; j < n; ++j)
                d[i][j] = Math.min(d[i][j], d[i][x] + w + d[y][j]);
    }

    public int shortestPath(int start, int end) {
        int ans = d[start][end];
        return ans < INF / 3 ? ans : -1;
    }
}
```

```cpp [sol2-C++]
class Graph {
    vector<vector<int>> d;
public:
    Graph(int n, vector<vector<int>> &edges) {
        // 邻接矩阵（初始化为无穷大，表示 i 到 j 没有边）
        d = vector<vector<int>>(n, vector<int>(n, INT_MAX / 3));
        for (int i = 0; i < n; ++i)
            d[i][i] = 0;
        for (auto &e: edges)
            d[e[0]][e[1]] = e[2]; // 添加一条边（输入保证没有重边和自环）
        for (int k = 0; k < n; ++k)
            for (int i = 0; i < n; ++i)
                for (int j = 0; j < n; ++j)
                    d[i][j] = min(d[i][j], d[i][k] + d[k][j]);
    }

    void addEdge(vector<int> e) {
        int x = e[0], y = e[1], w = e[2], n = d.size();
        if (w >= d[x][y]) // 无需更新
            return;
        for (int i = 0; i < n; ++i)
            for (int j = 0; j < n; ++j)
                d[i][j] = min(d[i][j], d[i][x] + w + d[y][j]);
    }

    int shortestPath(int start, int end) {
        int ans = d[start][end];
        return ans < INT_MAX / 3 ? ans : -1;
    }
};
```

```go [sol2-Go]
const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	d := make([][]int, n) // 邻接矩阵
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if j != i {
				d[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
			}
		}
	}
	for _, e := range edges {
		d[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边和自环）
	}
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	return d
}

func (d Graph) AddEdge(e []int) {
	x, y, w := e[0], e[1], e[2]
	if w >= d[x][y] { // 无需更新
		return
	}
	for i := range d {
		for j := range d {
			d[i][j] = min(d[i][j], d[i][x]+w+d[y][j])
		}
	}
}

func (d Graph) ShortestPath(start, end int) int {
	ans := d[start][end]
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n^3 + qn^2)$，其中 $q$ 为 $\texttt{addEdge}$ 的调用次数。
- 空间复杂度：$O(n^2)$。
