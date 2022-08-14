[视频讲解](https://www.bilibili.com/video/BV1CN4y1V7uE) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

--- 
 
用哈希表记录哪些节点是受限的，建图的时候只有当两个节点都不是受限的才连边。然后 DFS 这棵树，统计从 $0$ 出发能访问到的节点数，即为答案。

```py [sol1-Python3]
class Solution:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        r = set(restricted)
        g = [[] for _ in range(n)]
        for x, y in edges:
            if x not in r and y not in r:
                g[x].append(y)
                g[y].append(x)
        ans = 0
        def dfs(x: int, fa: int):
            nonlocal ans
            ans += 1
            for y in g[x]:
                if y != fa:
                    dfs(y, x)
        dfs(0, -1)
        return ans
```

```go [sol1-Go]
func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	r := make(map[int]bool, len(restricted))
	for _, x := range restricted {
		r[x] = true
	}
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		if !r[x] && !r[y] {
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}
	}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		ans++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return
}
```
