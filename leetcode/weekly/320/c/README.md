考虑每条边上至少需要多少辆车。

以 $0$ 为根，设子树 $x$ 的大小为 $\textit{size}$，那么它到它父节点这条边就至少需要 $\left\lceil\dfrac{\textit{size}}{\textit{seats}}\right\rceil$ 辆车。

累加除了 $x=0$ 以外的值，就是答案。

```py [sol1-Python3]
class Solution:
    def minimumFuelCost(self, roads: List[List[int]], seats: int) -> int:
        ans = 0
        g = [[] for _ in range(len(roads) + 1)]
        for x, y in roads:
            g[x].append(y)
            g[y].append(x)
        def dfs(x: int, fa: int) -> int:
            size = 1
            for y in g[x]:
                if y != fa:
                    size += dfs(y, x)
            if x:
                nonlocal ans
                ans += (size + seats - 1) // seats
            return size
        dfs(0, -1)
        return ans
```

```go [sol1-Go]
func minimumFuelCost(roads [][]int, seats int) int64 {
	ans := 0
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa {
				size += dfs(y, x)
			}
		}
		if x > 0 {
			ans += (size + seats - 1) / seats
		}
		return size
	}
	dfs(0, -1)
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{roads}$ 的长度。
- 空间复杂度：$O(n)$。
