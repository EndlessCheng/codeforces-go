#### 提示 1

答案（最小带权子图）会是个什么形状呢？

#### 提示 2

如果只有两个点，那就是在算两点间的最短路。三个点的话，也可以考虑最短路。

#### 提示 3

枚举三条最短路的交点。

---

![111.png](https://pic.leetcode-cn.com/1647142194-ulupZX-111.png)

答案会是一个「三岔口」的形状（也有可能是一条链，仍可视为三岔口），我们可以枚举三岔口的交点 $x$，然后求 $\textit{src}_1$ 和 $\textit{src}_2$ 到 $x$ 的最短路，以及 $x$ 到 $\textit{dest}$ 的最短路，这可以通过在反图（即所有边反向后的图）上求从 $\textit{dest}$ 出发的最短路得出。

累加这三条最短路的和，即为三岔口在 $x$ 处的子图的边权和。

枚举所有 $x$，最小的子图的边权和即为答案。

```Python [sol1-Python3]
class Solution:
    def minimumWeight(self, n: int, edges: List[List[int]], src1: int, src2: int, dest: int) -> int:
        def dijkstra(g: List[List[tuple]], start: int) -> List[int]:
            dis = [inf] * n
            dis[start] = 0
            pq = [(0, start)]
            while pq:
                d, x = heappop(pq)
                if d > dis[x]:
                    continue
                for y, wt in g[x]:
                    new_d = dis[x] + wt
                    if new_d < dis[y]:
                        dis[y] = new_d
                        heappush(pq, (new_d, y))
            return dis

        g = [[] for _ in range(n)]
        rg = [[] for _ in range(n)]
        for x, y, wt in edges:
            g[x].append((y, wt))
            rg[y].append((x, wt))

        d1 = dijkstra(g, src1)
        d2 = dijkstra(g, src2)
        d3 = dijkstra(rg, dest)

        ans = min(sum(d) for d in zip(d1, d2, d3))
        return ans if ans < inf else -1
```

```go [sol1-Go]
type edge struct{ to, wt int }
func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt64 / 3
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if p.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dis[v] + e.wt; newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, pair{w, newD})
			}
		}
	}
	return dis
}

func minimumWeight(n int, edges [][]int, src1, src2, dest int) int64 {
	g := make([][]edge, n)
	rg := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], edge{w, wt})
		rg[w] = append(rg[w], edge{v, wt})
	}

	d1 := dijkstra(g, src1)
	d2 := dijkstra(g, src2)
	d3 := dijkstra(rg, dest)

	ans := math.MaxInt64 / 3
	for x := 0; x < n; x++ {
		ans = min(ans, d1[x]+d2[x]+d3[x])
	}
	if ans < math.MaxInt64 / 3 {
		return int64(ans)
	}
	return -1
}

type pair struct{ v, dis int }
type hp []pair
func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func min(a, b int) int { if a > b { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
    vector<long> dijkstra(vector<vector<pair<int, int>>> &g, int start) {
        vector<long> dis(g.size(), LONG_MAX / 3);
        dis[start] = 0;
        priority_queue<pair<long, int>, vector<pair<long, int>>, greater<>> pq;
        pq.emplace(0, start);
        while (!pq.empty()) {
            auto [d, x] = pq.top();
            pq.pop();
            if (d > dis[x]) continue;
            for (auto [y, wt]: g[x]) {
                long newD = dis[x] + wt;
                if (newD < dis[y]) {
                    dis[y] = newD;
                    pq.emplace(newD, y);
                }
            }
        }
        return dis;
    }

public:
    long long minimumWeight(int n, vector<vector<int>> &edges, int src1, int src2, int dest) {
        vector<vector<pair<int, int>>> g(n), rg(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            rg[y].emplace_back(x, wt);
        }

        auto d1 = dijkstra(g, src1);
        auto d2 = dijkstra(g, src2);
        auto d3 = dijkstra(rg, dest);

        long ans = LONG_MAX / 3;
        for (int x = 0; x < n; ++x)
            ans = min(ans, d1[x] + d2[x] + d3[x]);
        return ans < LONG_MAX / 3 ? ans : -1;
    }
};
```

```java [sol1-Java]
class Solution {
    long[] dijkstra(List<Pair<Integer, Integer>>[] g, int start) {
        var dis = new long[g.length];
        Arrays.fill(dis, Long.MAX_VALUE / 3);
        dis[start] = 0;
        // 比较器也可以用 Comparator.comparingLong(Pair::getValue)，但是性能不够好
        var pq = new PriorityQueue<Pair<Integer, Long>>((a, b) -> (int) (a.getValue() - b.getValue()));
        pq.offer(new Pair<>(start, 0L));
        while (!pq.isEmpty()) {
            var p = pq.poll();
            var x = p.getKey();
            var wt = p.getValue();
            if (wt > dis[x]) continue;
            for (var e : g[x]) {
                var y = e.getKey();
                var newD = wt + e.getValue();
                if (newD < dis[y]) {
                    dis[y] = newD;
                    pq.offer(new Pair<>(y, newD));
                }
            }
        }
        return dis;
    }

    public long minimumWeight(int n, int[][] edges, int src1, int src2, int dest) {
        List<Pair<Integer, Integer>>[] g = new ArrayList[n], rg = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        Arrays.setAll(rg, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].add(new Pair<>(y, wt));
            rg[y].add(new Pair<>(x, wt));
        }

        var d1 = dijkstra(g, src1);
        var d2 = dijkstra(g, src2);
        var d3 = dijkstra(rg, dest);

        var ans = Long.MAX_VALUE / 3;
        for (var x = 0; x < n; x++) {
            ans = Math.min(ans, d1[x] + d2[x] + d3[x]);
        }
        return ans < Long.MAX_VALUE / 3 ? ans : -1;
    }
}
```
