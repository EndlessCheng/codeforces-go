下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

把起点、终点和所有特殊路径的终点看成是图上的点。

定义 $\textit{dis}[i]$ 表示从 $\textit{start}$ 到 $i$ 的最短路的长度。初始时 $\textit{dis}[\textit{start}]=0$，其余 $\textit{dis}[i]$ 为 $\infty$。

首先，从 $\textit{start}$ 出发，更新邻居的最短路。

下一步，寻找除去 $\textit{start}$ 的 $\textit{dis}$ 的最小值，设这个点为 $x$，那么 $\textit{dis}[x]$ 就已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

证明：由于除去起点的其余 $\textit{dis}[i]$ 都不低于 $\textit{dis}[x]$，且图中边权都非负，那么从另外一个点 $y$ 去更新 $\textit{dis}[x]$ 时，是无法让 $\textit{dis}[x]$ 变得更小的（因为 $\textit{dis}[x]$ 是当前最小），因此 $\textit{dis}[x]$ 已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。

由于在寻找 $\textit{dis}$ 的最小值时，需要排除在前面的循环中找到的 $x$（因为已经更新 $x$ 到其它点的最短路了，无需反复更新），可以用一个 $\textit{vis}$ 数组标记这些 $x$。

以上，通过**数学归纳法**，可以证明每次找到的未被标记的 $\textit{dis}$ 的最小值就是最短路。

由于输入的图是**稠密图**，所以采用邻接矩阵实现。

```py [sol1-Python3]
class Solution:
    def minimumCost(self, start: List[int], target: List[int], specialRoads: List[List[int]]) -> int:
        t = tuple(target)
        dis = defaultdict(lambda: inf)
        dis[tuple(start)] = 0
        vis = set()
        while True:
            v = None
            for p, d in dis.items():
                if p not in vis and (v is None or d < dis[v]):
                    v = p
            if v == t:  # 到终点的最短路已确定
                return dis[v]
            vis.add(v)
            vx, vy = v
            dis[t] = min(dis[t], dis[v] + t[0] - vx + t[1] - vy)  # 更新到终点的最短路
            for x1, y1, x2, y2, cost in specialRoads:
                w = (x2, y2)
                # 要么直接到 (x2,y2)，要么走特殊路径到 (x2,y2)
                d = dis[v] + min(abs(x2 - vx) + abs(y2 - vy), abs(x1 - vx) + abs(y1 - vy) + cost)
                dis[w] = min(dis[w], d)
```

```go [sol1-Go]
func minimumCost(start, target []int, specialRoads [][]int) int {
	type pair struct{ x, y int }
	t := pair{target[0], target[1]}
	dis := make(map[pair]int, len(specialRoads)+2)
	dis[t] = math.MaxInt
	dis[pair{start[0], start[1]}] = 0
	vis := make(map[pair]bool, len(specialRoads)+2)
	for {
		v := pair{}
		for p, d := range dis {
			if !vis[p] && (v.x == 0 || d < dis[v]) {
				v = p
			}
		}
		if v == t { // 到终点的最短路已确定
			return dis[t]
		}
		vis[v] = true
		dis[t] = min(dis[t], dis[v]+t.x-v.x+t.y-v.y) // 更新到终点的最短路
		for _, r := range specialRoads {
			x, y := r[2], r[3]
			w := pair{x, y}
			// 要么直接到 (x,y)，要么走特殊路径到 (x,y)
			d := dis[v] + min(abs(x-v.x)+abs(y-v.y), abs(r[0]-v.x)+abs(r[1]-v.y)+r[4])
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{specialRoads}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
