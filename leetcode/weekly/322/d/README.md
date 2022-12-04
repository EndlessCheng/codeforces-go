[视频讲解](https://www.bilibili.com/video/BV15d4y147YF) 已出炉，包含如何判定二分图的讲解，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

题目所说的 $|y-x|=1$ 这个要求，其实就是 BFS 这张图，让同一层的节点分到同一个编号。

由于数据范围比较小，枚举每个点作为起点，跑 BFS，求出最大编号（最大深度）。

每个连通块的最大深度相加，即为答案。

进一步地，$|y-x|=1$ 这个要求，也可以看成是从某一点出发，通过两条**不同**路径到达另一个点时，这两条路径长度的奇偶性是相同的。

换句话说，图中所有的环都必须有偶数个顶点，这等价于图是二分图，判定方式见 [785. 判断二分图](https://leetcode.cn/problems/is-graph-bipartite/)，原理见视频讲解。

```py [sol1-Python3]
class Solution:
    def magnificentSets(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x - 1].append(y - 1)
            g[y - 1].append(x - 1)

        time = [0] * n  # 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
        clock = 0
        def bfs(start: int) -> int:  # 返回从 start 出发的最大深度
            depth = 0
            nonlocal clock
            clock += 1
            time[start] = clock
            q = [start]
            while q:
                tmp = q
                q = []
                for x in tmp:
                    for y in g[x]:
                        if time[y] != clock:  # 没有在同一次 BFS 中访问过
                            time[y] = clock
                            q.append(y)
                depth += 1
            return depth

        color = [0] * n
        def is_bipartite(x: int, c: int) -> bool:  # 二分图判定，原理见视频讲解
            nodes.append(x)
            color[x] = c
            for y in g[x]:
                if color[y] == c or color[y] == 0 and not is_bipartite(y, -c):
                    return False
            return True

        ans = 0
        for i, c in enumerate(color):
            if c: continue
            nodes = []
            if not is_bipartite(i, 1): return -1  # 如果不是二分图（有奇环），则无法分组
            # 否则一定可以分组
            ans += max(bfs(x) for x in nodes)  # 枚举连通块的每个点，作为起点 BFS，求最大深度
        return ans
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private final List<Integer> nodes = new ArrayList<>();
    private int[] time, color; // time 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
    private int clock;

    public int magnificentSets(int n, int[][] edges) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0] - 1, y = e[1] - 1;
            g[x].add(y);
            g[y].add(x);
        }

        time = new int[n];
        color = new int[n];
        var ans = 0;
        for (var i = 0; i < n; i++) {
            if (color[i] != 0) continue;
            nodes.clear();
            if (!isBipartite(i, 1)) return -1; // 如果不是二分图（有奇环），则无法分组
            // 否则一定可以分组
            var maxDepth = 0;
            for (var x : nodes) // 枚举连通块的每个点，作为起点 BFS，求最大深度
                maxDepth = Math.max(maxDepth, bfs(x));
            ans += maxDepth;
        }
        return ans;
    }

    // 二分图判定，原理见视频讲解
    private boolean isBipartite(int x, int c) {
        nodes.add(x);
        color[x] = c;
        for (var y : g[x])
            if (color[y] == c || color[y] == 0 && !isBipartite(y, -c))
                return false;
        return true;
    }

    // 返回从 start 出发的最大深度
    private int bfs(int start) {
        var depth = 0;
        time[start] = ++clock;
        var q = new ArrayList<Integer>();
        q.add(start);
        while (!q.isEmpty()) {
            var tmp = q;
            q = new ArrayList<>();
            for (var x : tmp)
                for (var y : g[x])
                    if (time[y] != clock) { // 没有在同一次 BFS 中访问过
                        time[y] = clock;
                        q.add(y);
                    }
            ++depth;
        }
        return depth;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int magnificentSets(int n, vector<vector<int>> &edges) {
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0] - 1, y = e[1] - 1;
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int time[n], clock = 0; // time 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
        memset(time, 0, sizeof(time));
        auto bfs = [&](int start) -> int { // 返回从 start 出发的最大深度
            int depth = 0;
            time[start] = ++clock;
            vector<int> q = {start};
            while (!q.empty()) {
                vector<int> nxt;
                for (int x : q)
                    for (int y : g[x])
                        if (time[y] != clock) { // 没有在同一次 BFS 中访问过
                            time[y] = clock;
                            nxt.push_back(y);
                        }
                q = move(nxt);
                ++depth;
            }
            return depth;
        };

        int8_t color[n]; memset(color, 0, sizeof(color));
        vector<int> nodes;
        function<bool(int, int8_t)> is_bipartite = [&](int x, int8_t c) -> bool { // 二分图判定，原理见视频讲解
            nodes.push_back(x);
            color[x] = c;
            for (int y : g[x])
                if (color[y] == c || color[y] == 0 && !is_bipartite(y, -c))
                    return false;
            return true;
        };

        int ans = 0;
        for (int i = 0; i < n; ++i) {
            if (color[i]) continue;
            nodes.clear();
            if (!is_bipartite(i, 1)) return -1; // 如果不是二分图（有奇环），则无法分组
            // 否则一定可以分组
            int max_depth = 0;
            for (int x : nodes) // 枚举连通块的每个点，作为起点 BFS，求最大深度
                max_depth = max(max_depth, bfs(x));
            ans += max_depth;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func magnificentSets(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	time := make([]int, n) // 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
	clock := 0
	bfs := func(start int) (depth int) { // 返回从 start 出发的最大深度
		clock++
		time[start] = clock
		for q := []int{start}; len(q) > 0; depth++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if time[y] != clock { // 没有在同一次 BFS 中访问过
						time[y] = clock
						q = append(q, y)
					}
				}
			}
		}
		return
	}

	colors := make([]int8, n)
	var nodes []int
	var isBipartite func(int, int8) bool
	isBipartite = func(x int, c int8) bool { // 二分图判定，原理见视频讲解
		nodes = append(nodes, x)
		colors[x] = c
		for _, y := range g[x] {
			if colors[y] == c || colors[y] == 0 && !isBipartite(y, -c) {
				return false
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 {
			nodes = nil
			if !isBipartite(i, 1) { // 如果不是二分图（有奇环），则无法分组
				return -1
			}
			// 否则一定可以分组
			maxDepth := 0
			for _, x := range nodes { // 枚举连通块的每个点，作为起点 BFS，求最大深度
				maxDepth = max(maxDepth, bfs(x))
			}
			ans += maxDepth
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$O(n+m)$。
