阅读理解题。

由于可以折返，取连通块中的 $\textit{distance}_i$ 最小的那条边，即为答案。

```py [sol1-Python3]
class Solution:
    def minScore(self, n: int, roads: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y, d in roads:
            g[x - 1].append((y - 1, d))
            g[y - 1].append((x - 1, d))
        ans = inf
        vis = [False] * n
        def dfs(x: int) -> None:
            nonlocal ans
            vis[x] = True
            for y, d in g[x]:
                ans = min(ans, d)
                if not vis[y]:
                    dfs(y)
        dfs(0)
        return ans
```

```go [sol1-Go]
func minScore(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n)
	for _, e := range roads {
		x, y, d := e[0]-1, e[1]-1, e[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}
	ans := math.MaxInt32
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.d)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(0)
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $m$ 为 $\textit{roads}$ 的长度。
- 空间复杂度：$O(n+m)$。
