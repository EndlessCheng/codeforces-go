下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

用哈希表记录哪些节点是受限的，然后 DFS 这棵树，仅访问没有受限的子节点。

统计访问过的节点个数，即为答案。

```py [sol1-Python3]
class Solution:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)
        ans = 0
        r = set(restricted)
        def dfs(x: int, fa: int):
            nonlocal ans
            ans += 1
            for y in g[x]:
                if y != fa and y not in r:
                    dfs(y, x)
        dfs(0, -1)
        return ans
```

```go [sol1-Go]
func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	r := map[int]bool{}
	for _, x := range restricted {
		r[x] = true
	}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans++
		for _, y := range g[x] {
			if y != fa && !r[y] {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return
}
```
