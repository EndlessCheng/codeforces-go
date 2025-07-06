套模板就行。讲解见 [Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

设起点到 $x$ 的最短时间为 $d_x$。我们从 $x$ 出发，移动到其邻居 $y$。出发时间为 $\max(d_x, \textit{start}_i)$，必须 $\le \textit{end}_i$。到达 $y$ 的时间为 $\max(d_x, \textit{start}_i)+1$。

```py [sol-Python3]
class Solution:
    def minTime(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]
        for x, y, start, end in edges:
            g[x].append((y, start, end))

        dis = [inf] * n
        dis[0] = 0
        h = [(0, 0)]

        while h:
            dx, x = heappop(h)
            if dx > dis[x]:
                continue
            if x == n - 1:
                return dx
            for y, start, end in g[x]:
                new_dis = max(dx, start) + 1
                if new_dis - 1 <= end and new_dis < dis[y]:
                    dis[y] = new_dis
                    heappush(h, (new_dis, y))
        return -1
```

```java [sol-Java]
class Solution {
    public int minTime(int n, int[][] edges) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2], e[3]});
        }

        int[] dis = new int[n];
        Arrays.fill(dis, Integer.MAX_VALUE);
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> (a[0] - b[0]));
        dis[0] = 0;
        pq.offer(new int[]{0, 0});

        while (!pq.isEmpty()) {
            int[] p = pq.poll();
            int dx = p[0];
            int x = p[1];
            if (dx > dis[x]) {
                continue;
            }
            if (x == n - 1) {
                return dx;
            }
            for (int[] e : g[x]) {
                int y = e[0];
                int newDis = Math.max(dx, e[1]) + 1;
                if (newDis - 1 <= e[2] && newDis < dis[y]) {
                    dis[y] = newDis;
                    pq.offer(new int[]{newDis, y});
                }
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTime(int n, vector<vector<int>>& edges) {
        vector<vector<tuple<int, int, int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2], e[3]);
        }

        vector<int> dis(n, INT_MAX);
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        dis[0] = 0;
        pq.emplace(0, 0);

        while (!pq.empty()) {
            auto [dx, x] = pq.top();
            pq.pop();
            if (dx > dis[x]) {
                continue;
            }
            if (x == n - 1) {
                return dx;
            }
            for (auto& [y, start, end] : g[x]) {
                int new_dis = max(dx, start) + 1;
                if (new_dis - 1 <= end && new_dis < dis[y]) {
                    dis[y] = new_dis;
                    pq.emplace(new_dis, y);
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minTime(n int, edges [][]int) int {
	type edge struct{ to, start, end int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, e[2], e[3]})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	h := hp{{}}

	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		d := p.d
		x := p.x
		if d > dis[x] {
			continue
		}
		if x == n-1 {
			return d
		}
		for _, e := range g[x] {
			y := e.to
			newD := max(d, e.start) + 1
			if newD-1 <= e.end && newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}
	return -1
}

type pair struct{ d, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m)$，其中 $m$ 是 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 专题训练

见下面图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
