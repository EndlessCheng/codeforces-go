## 方法一：Dijkstra 算法（分层图最短路）

做法和上场周赛的 [3970. 最多 K 个连续相同字符的最短路径](https://leetcode.cn/problems/shortest-path-with-at-most-k-consecutive-identical-characters/) 是一样的，[我的题解](https://leetcode.cn/problems/shortest-path-with-at-most-k-consecutive-identical-characters/solutions/3986163/dijkstra-suan-fa-fen-ceng-tu-zui-duan-lu-magt/)。

定义 $\textit{dis}[x][\textit{rem}]$ 表示从起点 $\textit{source}$ 到节点 $x$ 的最短路长度，且在 $x$ 处的剩余电量为 $\textit{rem}$。

初始值 $\textit{dis}[\textit{source}][\textit{power}] = 0$。

堆中元素按照最短路长度为第一关键字升序，剩余电量为第二关键字降序排序。

**注**：根据 Dijkstra 算法的原理，当终点 $\textit{target}$ 首次出堆时，我们就算出了从起点到终点的最短路，且剩余电量最大，可以直接返回答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minTimeMaxPower(self, n: int, edges: list[list[int]], power: int, cost: list[int], source: int, target: int) -> list[int]:
        g = [[] for _ in range(n)]
        for x, y, t in edges:
            g[x].append((y, t))

        dis = [[inf] * (power + 1) for _ in range(n)]
        dis[source][power] = 0
        h = [(0, -power, source)]  # (最短路长度, -剩余电量, 节点编号)

        while h:
            d, rem, x = heappop(h)
            rem = -rem
            if x == target:
                return [d, rem]
            if d > dis[x][rem] or rem < cost[x]:
                continue
            rem -= cost[x]
            for y, t in g[x]:
                new_dis = d + t
                if new_dis < dis[y][rem]:
                    dis[y][rem] = new_dis
                    heappush(h, (new_dis, -rem, y))

        return [-1, -1]
```

```java [sol-Java]
class Solution {
    public long[] minTimeMaxPower(int n, int[][] edges, int power, int[] cost, int source, int target) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }

        long[][] dis = new long[n][power + 1];
        for (long[] row : dis) {
            Arrays.fill(row, Long.MAX_VALUE);
        }
        dis[source][power] = 0;

        // int[]{最短路长度, 剩余电量, 节点编号}
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> a[0] != b[0] ? Long.compare(a[0], b[0]) : (int) b[1] - (int) a[1]);
        pq.add(new long[]{0, power, source});

        while (!pq.isEmpty()) {
            long[] top = pq.poll();
            long d = top[0];
            int rem = (int) top[1];
            int x = (int) top[2];
            if (x == target) {
                return new long[]{d, rem};
            }
            if (d > dis[x][rem] || rem < cost[x]) {
                continue;
            }
            int newRem = rem - cost[x];
            for (int[] e : g[x]) {
                int y = e[0];
                long newDis = d + e[1];
                if (newDis < dis[y][newRem]) {
                    dis[y][newRem] = newDis;
                    pq.add(new long[]{newDis, newRem, y});
                }
            }
        }

        return new long[]{-1, -1};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> minTimeMaxPower(int n, vector<vector<int>>& edges, int power, vector<int>& cost, int source, int target) {
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2]);
        }

        vector dis(n, vector<long long>(power + 1, LLONG_MAX));
        dis[source][power] = 0;
        // tuple{最短路长度, -剩余电量, 节点编号}
        priority_queue<tuple<long long, int, int>, vector<tuple<long long, int, int>>, greater<>> pq;
        pq.emplace(0, -power, source);

        while (!pq.empty()) {
            auto [d, rem, x] = pq.top();
            pq.pop();
            rem = -rem;
            if (x == target) {
                return {d, rem};
            }
            if (d > dis[x][rem] || rem < cost[x]) {
                continue;
            }
            auto new_rem = rem - cost[x];
            for (auto& [y, w] : g[x]) {
                auto new_dis = d + w;
                if (new_dis < dis[y][new_rem]) {
                    dis[y][new_rem] = new_dis;
                    pq.emplace(new_dis, -new_rem, y);
                }
            }
        }

        return {-1, -1};
    }
};
```

```go [sol-Go]
func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	type edge struct{ to, t int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
	}

	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, power+1)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[source][power] = 0
	h := hp{{0, source, power}}

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d, x, rem := top.dis, top.x, top.rem
		if x == target {
			return []int64{int64(d), int64(rem)}
		}
		if d > dis[x][rem] || rem < cost[x] {
			continue
		}
		rem -= cost[x]
		for _, e := range g[x] {
			y := e.to
			newD := d + e.t
			if newD < dis[y][rem] {
				dis[y][rem] = newD
				heap.Push(&h, tuple{newD, y, rem})
			}
		}
	}

	return []int64{-1, -1}
}

// 最短路长度, 节点编号, 剩余电量
type tuple struct{ dis, x, rem int }
type hp []tuple

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.dis < b.dis || a.dis == b.dis && a.rem > b.rem
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk + mk\log (mk))$，其中 $m$ 是 $\textit{edges}$ 的长度，$k = \textit{power}$。分层图中有 $\mathcal{O}(nk)$ 个点，$\mathcal{O}(mk)$ 条边。注意我们用的是懒更新堆，堆中有 $\mathcal{O}(mk)$ 个元素。创建 $\textit{dis}$ 数组需要 $\mathcal{O}(nk)$ 时间。
- 空间复杂度：$\mathcal{O}(nk + mk)$。

## 方法二：DAG DP

注意 $\textit{cost}[i] > 0$，每走一步，剩余电量严格递减。所以我们不可能从一个剩余电量为 $100$ 的点出发，最后又回到剩余电量为 $100$ 的点。所以分层图是一个有向无环图（DAG），我们可以直接用 DAG DP 计算最短路长度。

设当前节点为 $x$，剩余电量为 $\textit{rem}$，最短路长度为 $d$。设 $x$ 的邻居为 $y$，边权为 $t$，那么用 $d+t$ 更新节点为 $y$、剩余电量为 $\textit{rem}-\textit{cost}[x]$ 的最短路长度的最小值。

为了实现这一过程，需要从大到小枚举剩余电量。

```py [sol-Python3]
class Solution:
    def minTimeMaxPower(self, n: int, edges: list[list[int]], power: int, cost: list[int], source: int, target: int) -> list[int]:
        g = [[] for _ in range(n)]
        for x, y, t in edges:
            g[x].append((y, t))

        f = [[inf] * n for _ in range(power + 1)]
        f[power][source] = 0
        min_dis = inf
        max_rem = -1

        for rem in range(power, -1, -1):
            if f[rem][target] < min_dis:
                min_dis = f[rem][target]
                max_rem = rem
            for x, v in enumerate(f[rem]):
                if v == inf or rem < cost[x]:
                    continue
                nxt_rem = rem - cost[x]
                for y, t in g[x]:
                    f[nxt_rem][y] = min(f[nxt_rem][y], v + t)  # 刷表法

        return [-1, -1] if max_rem < 0 else [min_dis, max_rem]
```

```java [sol-Java]
class Solution {
    public long[] minTimeMaxPower(int n, int[][] edges, int power, int[] cost, int source, int target) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }

        long[][] f = new long[power + 1][n];
        for (int i = 0; i <= power; i++) {
            Arrays.fill(f[i], Long.MAX_VALUE);
        }
        f[power][source] = 0;
        long minDis = Long.MAX_VALUE;
        int maxRem = -1;

        for (int rem = power; rem >= 0; rem--) {
            if (f[rem][target] < minDis) {
                minDis = f[rem][target];
                maxRem = rem;
            }
            for (int x = 0; x < n; x++) {
                long v = f[rem][x];
                if (v == Long.MAX_VALUE || rem < cost[x]) {
                    continue;
                }
                int nxtRem = rem - cost[x];
                for (int[] e : g[x]) {
                    f[nxtRem][e[0]] = Math.min(f[nxtRem][e[0]], v + e[1]); // 刷表法
                }
            }
        }

        return maxRem < 0 ? new long[]{-1, -1} : new long[]{minDis, maxRem};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> minTimeMaxPower(int n, vector<vector<int>>& edges, int power, vector<int>& cost, int source, int target) {
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2]);
        }

        vector f(power + 1, vector<long long>(n, LLONG_MAX));
        f[power][source] = 0;
        long long min_dis = LLONG_MAX;
        int max_rem = -1;

        for (int rem = power; rem >= 0; rem--) {
            if (f[rem][target] < min_dis) {
                min_dis = f[rem][target];
                max_rem = rem;
            }
            for (int x = 0; x < n; x++) {
                auto v = f[rem][x];
                if (v == LLONG_MAX || rem < cost[x]) {
                    continue;
                }
                int nxt_rem = rem - cost[x];
                for (auto& [y, t] : g[x]) {
                    f[nxt_rem][y] = min(f[nxt_rem][y], v + t); // 刷表法
                }
            }
        }

        if (max_rem < 0) {
            return {-1, -1};
        }
        return {min_dis, max_rem};
    }
};
```

```go [sol-Go]
func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	type edge struct{ to, t int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
	}

	f := make([][]int, power+1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
	}
	f[power][source] = 0

	minDis, maxRem := math.MaxInt, -1
	for rem := power; rem >= 0; rem-- {
		if f[rem][target] < minDis {
			minDis, maxRem = f[rem][target], rem
		}
		for x, v := range f[rem] {
			if v == math.MaxInt || rem < cost[x] {
				continue
			}
			nxtRem := rem - cost[x]
			for _, e := range g[x] {
				f[nxtRem][e.to] = min(f[nxtRem][e.to], v+e.t) // 刷表法
			}
		}
	}

	if maxRem < 0 {
		return []int64{-1, -1}
	}
	return []int64{int64(minDis), int64(maxRem)}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k(n+m))$，其中 $m$ 是 $\textit{edges}$ 的长度，$k = \textit{power}$。
- 空间复杂度：$\mathcal{O}(kn)$。

## 专题训练

1. 图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。
2. 动态规划题单的「**十三、图 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
