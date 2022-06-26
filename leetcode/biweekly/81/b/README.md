建图后，用 DFS 可以求出每个连通块的大小。

求连通块的大小的同时，用一个变量 $\textit{tot}$ 维护前面求出的连通块的大小之和。设当前连通块的大小为 $\textit{size}$，那么它对答案的贡献就是 $\textit{size}\cdot\textit{tot}$。

累加所有贡献，即为答案。

```Python [sol1-Python3]
class Solution:
    def countPairs(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        vis, ans, tot, size = [False] * n, 0, 0, 0
        def dfs(x: int) -> None:
            nonlocal size
            vis[x] = True
            size += 1
            for y in g[x]:
                if not vis[y]:
                    dfs(y)
        for i in range(n):
            if not vis[i]:
                size = 0
                dfs(i)
                ans += size * tot
                tot += size
        return ans
```

```go [sol1-Go]
func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	tot, size := 0, 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			size = 0
			dfs(i)
			ans += int64(size) * int64(tot)
			tot += size
		}
	}
	return
}
```
