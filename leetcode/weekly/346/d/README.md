## 视频讲解

见[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第四题，欢迎点赞投币！

## 什么时候无解？

题目要求把边权为 $-1$ 的至少修改为 $1$。如果都修改成 $1$，跑最短路（Dijkstra），发现从起点到终点的最短路长度大于 $\textit{target}$，那么由于边权变大，最短路不可能变小，所以此时无解。

另外一种无解的情况是，如果都修改成无穷大（本题限制为 $2\cdot 10^9$），发现从起点到终点的最短路长度小于 $\textit{target}$，那么由于边权变小，最短路不可能变大，所以此时也无解。

## 一个错误的思路

先把 $-1$ 都修改成 $1$，然后跑 Dijkstra。设从起点到终点的最短路长度为 $d$。

如果 $d<\textit{target}$，**看上去**把其中一个 $1$ 加上 $\textit{target}-d$，就可以使从起点到终点的最短路长度恰好为 $\textit{target}$ 了。

这是不对的，因为在增加边权后，最短路可能就走别的边了，不走刚才修改的这条边了。（具体见视频中画的例子。）

只修改一条边不行，那么修改两条边呢？也是不行的，最短路仍然可以走别的边。

## 正确思路

先把 $-1$ 都修改成 $1$，然后跑第一遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,0}$。

如果从起点到终点的最短路长度不超过 $\textit{target}$（否则无解，返回空数组），那我们就来尝试修改某些边权，使得从起点到终点的最短路长度恰好等于 $\textit{target}$。

这会引出两个问题：

1. 要按照什么样的**顺序**修改这些边？
2. 修改成多少合适？

正所谓「牵一发而动全身」，从上面对错误思路的分析可知，仅仅修改一个边权，就可能影响很多最短路的值。

那不妨再跑一遍 Dijkstra，由于 Dijkstra 算法保证每次拿到的点的最短路就是最终的最短路，所以按照 Dijkstra 算法遍历点/边的顺序去修改，就不会对**已确定的**最短路产生影响。

对于第二遍 Dijkstra，设从 $\textit{source}$ 到点 $i$ 的最短路长度为 $d_{i,1}$。

对于一条可以修改的边 $x-y$，假设要把它的边权改为 $w$，那么 $\textit{source}-x-y-\textit{destination}$ 这条路径由三部分组成：

1. 从 $\textit{source}$ 到 $x$ 的最短路，这是第二遍 Dijkstra 算出来的，即 $d_{x,1}$。
2. 从 $x$ 到 $y$，即 $w$。
3. 从 $y$ 到 $\textit{destination}$ 的最短路，由于后面的边还没有修改，这个最短路是第一遍 Dijkstra 算出来的，即 $d_{\textit{destination},0} - d_{y,0}$。

这三部分之和需要等于 $\textit{target}$，所以有

$$
d_{x,1} + w + d_{\textit{destination},0} - d_{y,0} = \textit{target}
$$

解得

$$
w = \textit{target} - d_{\textit{destination},0} + d_{y,0} - d_{x,1}
$$

> 注意上式中的 $\textit{target} - d_{\textit{destination},0}$ 是一个定值，代码中用 $\textit{delta}$ 表示。

根据「什么时候无解」中的分析，如果第二遍 Dijkstra 跑完后，从起点到终点的最短路仍然小于 $\textit{target}$，那么就说明无法修改，返回空数组。

否则，答案就是我们在第二遍 Dijkstra 中作出的修改。注意第二遍 Dijkstra 跑完后可能还有些边是 $-1$（因为在 $w=1$ 的时候没有修改，或者有些边不影响最短路），把这些边都改成 $1$ 就行。

代码实现时，为了修改边权，需要在邻接表中额外记录边的编号。

此外，由于输入最坏是稠密图，可以用朴素版的 Dijkstra 算法。

```py [sol1-Python3]
class Solution:
    def modifiedGraphEdges(self, n: int, edges: List[List[int]], source: int, destination: int, target: int) -> List[List[int]]:
        g = [[] for _ in range(n)]
        for i, (x, y, _) in enumerate(edges):
            g[x].append((y, i))
            g[y].append((x, i))  # 建图，额外保存边的编号

        dis = [[inf, inf] for _ in range(n)]
        dis[source] = [0, 0]

        def dijkstra(k: int) -> None:
            vis = [False] * n
            while True:
                # 找一个最短路最小并且还没有 vis 的点（第一轮循环找的是起点 source）
                x = -1
                for y, (b, d) in enumerate(zip(vis, dis)):
                    if not b and (x < 0 or d[k] < dis[x][k]):
                        x = y
                if x == destination:  # 起点 source 到终点 destination 的最短路已确定
                    return
                vis[x] = True  # 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
                for y, eid in g[x]:
                    wt = edges[eid][2]
                    if wt == -1:
                        wt = 1  # -1 改成 1
                    if k == 1 and edges[eid][2] == -1:
                        # 第二次 Dijkstra，改成 w
                        w = delta + dis[y][0] - dis[x][1]
                        if w > wt:
                            edges[eid][2] = wt = w  # 直接在 edges 上修改
                    # 更新最短路
                    dis[y][k] = min(dis[y][k], dis[x][k] + wt)

        dijkstra(0)
        delta = target - dis[destination][0]
        if delta < 0:  # -1 全改为 1 时，最短路比 target 还大
            return []

        dijkstra(1)
        if dis[destination][1] < target:  # 最短路无法再变大，无法达到 target
            return []

        for e in edges:
            if e[2] == -1:  # 剩余没修改的边全部改成 1
                e[2] = 1
        return edges
```

```go [sol1-Go]
func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i}) // 建图，额外记录边的编号
	}

	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{math.MaxInt, math.MaxInt}
	}
	dis[source] = [2]int{}
	dijkstra := func(k int) {
		vis := make([]bool, n)
		for {
			// 找一个最短路最小并且还没有 vis 的点（第一轮循环找的是起点 source）
			x := -1
			for y, b := range vis {
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination { // 起点 source 到终点 destination 的最短路已确定
				return
			}
			vis[x] = true
			for _, e := range g[x] {
				y, wt := e.to, edges[e.eid][2]
				if wt == -1 {
					wt = 1 // -1 改成 1
				}
				if k == 1 && edges[e.eid][2] == -1 {
					// 第二次 Dijkstra，改成 w
					w := delta + dis[y][0] - dis[x][1]
					if w > wt {
						wt = w
						edges[e.eid][2] = w // 直接在 edges 上修改
					}
				}
				// 更新最短路
				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
			}
		}
	}

	dijkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大
		return nil
	}

	dijkstra(1)
	if dis[destination][1] < target { // 最短路无法再变大，无法达到 target
		return nil
	}

	for _, e := range edges {
		if e[2] == -1 { // 剩余没修改的边全部改成 1
			e[2] = 1
		}
	}
	return edges
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{edges}$ 的长度。注意输入是连通图，$m$ 至少为 $n-1$，所以 $\mathcal{O}(n+m)=\mathcal{O}(m)$
