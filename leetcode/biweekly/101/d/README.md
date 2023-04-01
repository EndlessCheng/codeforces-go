下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

![b101_t4_cut.png](https://pic.leetcode.cn/1680363054-UnoCDM-b101_t4_cut.png)

```py [sol1-Python3]
class Solution:
    def findShortestCycle(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图

        def bfs(start: int) -> int:
            dis = [-1] * n  # dis[i] 表示从 start 到 i 的最短路长度
            dis[start] = 0
            q = deque([(start, -1)])
            while q:
                x, fa = q.popleft()
                for y in g[x]:
                    if dis[y] < 0:  # 第一次遇到
                        dis[y] = dis[x] + 1
                        q.append((y, x))
                    elif y != fa:  # 第二次遇到
                        # 由于是 BFS，后面不会遇到更短的环，直接返回
                        return dis[x] + dis[y] + 1
            return inf  # 该连通块无环

        ans = min(bfs(i) for i in range(n))
        return ans if ans < inf else -1
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private int[] dis; // dis[i] 表示从 start 到 i 的最短路长度

    public int findShortestCycle(int n, int[][] edges) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }
        dis = new int[n];

        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < n; ++i) // 枚举每个起点跑 BFS
            ans = Math.min(ans, bfs(i));
        return ans < Integer.MAX_VALUE ? ans : -1;
    }

    private int bfs(int start) {
        Arrays.fill(dis, -1);
        dis[start] = 0;
        var q = new ArrayDeque<int[]>();
        q.add(new int[]{start, -1});
        while (!q.isEmpty()) {
            var p = q.poll();
            int x = p[0], fa = p[1];
            for (int y : g[x])
                if (dis[y] < 0) { // 第一次遇到
                    dis[y] = dis[x] + 1;
                    q.add(new int[]{y, x});
                } else if (y != fa) // 第二次遇到
                    // 由于是 BFS，后面不会遇到更短的环了
                    return dis[x] + dis[y] + 1;
        }
        return Integer.MAX_VALUE; // 该连通块无环
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int findShortestCycle(int n, vector<vector<int>> &edges) {
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        int dis[n]; // dis[i] 表示从 start 到 i 的最短路长度
        auto bfs = [&](int start) -> int {
            memset(dis, -1, sizeof(dis));
            dis[start] = 0;
            queue<pair<int, int>> q;
            q.emplace(start, -1);
            while (!q.empty()) {
                auto [x, fa] = q.front();
                q.pop();
                for (int y: g[x])
                    if (dis[y] < 0) { // 第一次遇到
                        dis[y] = dis[x] + 1;
                        q.emplace(y, x);
                    } else if (y != fa) // 第二次遇到
                        // 由于是 BFS，后面不会遇到更短的环了
                        return dis[x] + dis[y] + 1;
            }
            return INT_MAX; // 该连通块无环
        };
        int ans = INT_MAX;
        for (int i = 0; i < n; ++i) // 枚举每个起点跑 BFS
            ans = min(ans, bfs(i));
        return ans < INT_MAX ? ans : -1;
    }
};
```

```go [sol1-Go]
func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	ans := math.MaxInt
	dis := make([]int, n) // dis[i] 表示从 start 到 i 的最短路长度
next:
	for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
		for j := range dis {
			dis[j] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		q := []pair{{start, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, fa := p.x, p.fa
			for _, y := range g[x] {
				if dis[y] < 0 { // 第一次遇到
					dis[y] = dis[x] + 1
					q = append(q, pair{y, x})
				} else if y != fa { // 第二次遇到
					ans = min(ans, dis[x]+dis[y]+1)
					continue next // 由于是 BFS，后面不会遇到更短的环了，直接枚举下一个 start
				}
			}
		}
	}
	if ans == math.MaxInt { // 无环图
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{edges}$ 的长度。每次 BFS 需要 $O(m)$ 的时间。
- 空间复杂度：$O(n+m)$。
