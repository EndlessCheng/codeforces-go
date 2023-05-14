下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

DFS 每个连通块，统计当前连通块的点数 $v$ 和边数 $e$。

- 每访问一个点，就把 $v$ 加一。
- $e$ 加上点 $v$ 的邻居个数。注意这样一条边会统计两次。

在完全图中，任意两点之间都有边，相当于从 $v$ 个点中选 $2$ 个点的方案数。所以有

$$
e = \dfrac{v(v-1)}{2}
$$

由于上面统计的时候，一条边统计了两次，所以代码中的判断条件是 

$$
e = v(v-1)
$$

```py [sol1-Python3]
class Solution:
    def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        vis = [False] * n
        def dfs(x: int) -> None:
            vis[x] = True
            nonlocal v, e
            v += 1
            e += len(g[x])
            for y in g[x]:
                if not vis[y]:
                    dfs(y)

        ans = 0
        for i, b in enumerate(vis):
            if not b:
                v = e = 0
                dfs(i)
                ans += e == v * (v - 1)
        return ans
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private boolean vis[];
    private int v, e;

    public int countCompleteComponents(int n, int[][] edges) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }

        int ans = 0;
        vis = new boolean[n];
        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                v = 0;
                e = 0;
                dfs(i);
                if (e == v * (v - 1))
                    ans++;
            }
        }
        return ans;
    }

    private void dfs(int x) {
        vis[x] = true;
        v++;
        e += g[x].size();
        for (var y : g[x])
            if (!vis[y])
                dfs(y);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int countCompleteComponents(int n, vector<vector<int>> &edges) {
        vector<vector<int>> g(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        vector<int> vis(n);
        int ans = 0, v, e;
        function<void(int)> dfs = [&](int x) {
            vis[x] = true;
            v++;
            e += g[x].size();
            for (int y: g[x])
                if (!vis[y])
                    dfs(y);
        };

        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                v = 0;
                e = 0;
                dfs(i);
                ans += e == v * (v - 1);
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countCompleteComponents(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	var v, e int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		v++
		e += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			v, e = 0, 0
			dfs(i)
			if e == v*(v-1) {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。
