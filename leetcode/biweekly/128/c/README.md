**请先阅读** [Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

对于本题，$\textit{answer}$ 几乎就是 $\textit{dis}$ 数组。只需要在 Dijkstra 算法的过程中，添加一个判断：

- 在更新最短路之前，如果最短路长度 $\ge \textit{disappear}[i]$，说明无法及时到达节点 $i$，不更新。

请看 [视频讲解](https://www.bilibili.com/video/BV1et42177VM/) 第三题，欢迎点赞关注！

### 答疑

**问**：我是这样做的，先不管 $\textit{disappear}$，用 Dijkstra 算法模板求出 $\textit{dis}$ 数组，然后把其中 $\textit{dis}[i]\ge \textit{disappear}[i]$ 的 $\textit{dis}[i]$ 改成 $-1$。这个做法是否正确？

**答**：这个做法是错的。考虑这样的数据：离节点 $0$ 近的节点 $x$，其 $\textit{disappear}[x]$ 很小；离节点 $0$ 远的节点 $y$，其 $\textit{disappear}[y]$ 很大。由于 $\textit{disappear}[x]$ 很小，我们无法通过节点 $x$，所以远处的节点 $y$ 实际上也无法到达。但错误做法会因为 $\textit{dis}[y] < \textit{disappear}[y]$，误认为节点 $y$ 可以到达，最终返回错误的答案。

```py [sol-Python3]
class Solution:
    def minimumTime(self, n: int, edges: List[List[int]], disappear: List[int]) -> List[int]:
        g = [[] for _ in range(n)]  # 稀疏图用邻接表
        for x, y, wt in edges:
            g[x].append((y, wt))
            g[y].append((x, wt))

        dis = [-1] * n
        dis[0] = 0
        h = [(0, 0)]
        while h:
            dx, x = heappop(h)
            if dx > dis[x]:  # x 之前出堆过
                continue
            for y, wt in g[x]:
                new_dis = dx + wt
                if new_dis < disappear[y] and (dis[y] < 0 or new_dis < dis[y]):
                    dis[y] = new_dis  # 更新 x 的邻居的最短路
                    heappush(h, (new_dis, y))
        return dis
```

```java [sol-Java]
class Solution {
    public int[] minimumTime(int n, int[][] edges, int[] disappear) {
        List<int[]>[] g = new ArrayList[n]; // 稀疏图用邻接表
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            int wt = e[2];
            g[x].add(new int[]{y, wt});
            g[y].add(new int[]{x, wt});
        }

        int[] dis = new int[n];
        Arrays.fill(dis, -1);
        dis[0] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> (a[0] - b[0]));
        pq.offer(new int[]{0, 0});
        while (!pq.isEmpty()) {
            int[] p = pq.poll();
            int dx = p[0];
            int x = p[1];
            if (dx > dis[x]) { // x 之前出堆过
                continue;
            }
            for (int[] e : g[x]) {
                int y = e[0];
                int newDis = dx + e[1];
                if (newDis < disappear[y] && (dis[y] < 0 || newDis < dis[y])) {
                    dis[y] = newDis; // 更新 x 的邻居的最短路
                    pq.offer(new int[]{newDis, y});
                }
            }
        }
        return dis;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> minimumTime(int n, vector<vector<int>>& edges, vector<int>& disappear) {
        vector<vector<pair<int, int>>> g(n); // 稀疏图用邻接表
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            g[y].emplace_back(x, wt);
        }

        vector<int> dis(n, -1);
        dis[0] = 0;
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        pq.emplace(0, 0);
        while (!pq.empty()) {
            auto [dx, x] = pq.top();
            pq.pop();
            if (dx > dis[x]) { // x 之前出堆过
                continue;
            }
            for (auto& [y, d] : g[x]) {
                int new_dis = dx + d;
                if (new_dis < disappear[y] && (dis[y] < 0 || new_dis < dis[y])) {
                    dis[y] = new_dis; // 更新 x 的邻居的最短路
                    pq.emplace(new_dis, y);
                }
            }
        }
        return dis;
    }
};
```

```go [sol-Go]
func minimumTime(n int, edges [][]int, disappear []int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 稀疏图用邻接表
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx := p.dis
		x := p.x
		if dx > dis[x] { // x 之前出堆过
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newDis := dx + e.wt
			if newDis < disappear[y] && (dis[y] < 0 || newDis < dis[y]) {
				dis[y] = newDis // 更新 x 的邻居的最短路
				heap.Push(&h, pair{newDis, y})
			}
		}
	}
	return dis
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

```js [sol-JavaScript]
var minimumTime = function(n, edges, disappear) {
    const g = Array.from({length: n}, () => []);
    for (const [x, y, wt] of edges) {
        g[x].push([y, wt]);
        g[y].push([x, wt]);
    }

    const dis = Array(n).fill(-1);
    dis[0] = 0;
    const pq = new MinPriorityQueue({priority: (p) => p[0]});
    pq.enqueue([0, 0]);
    while (!pq.isEmpty()) {
        const [dx, x] = pq.dequeue().element;
        if (dx > dis[x]) { // x 之前出堆过
            continue;
        }
        for (let [y, wt] of g[x]) {
            let new_dis = dx + wt;
            if (new_dis < disappear[y] && (dis[y] < 0 || new_dis < dis[y])) {
                dis[y] = new_dis; // 更新 x 的邻居的最短路
                pq.enqueue([new_dis, y]);
            }
        }
    }
    return dis;
};
```

```rust [sol-Rust]
use std::collections::BinaryHeap;

impl Solution {
    pub fn minimum_time(n: i32, edges: Vec<Vec<i32>>, disappear: Vec<i32>) -> Vec<i32> {
        let n = n as usize;
        let mut g = vec![vec![]; n]; // 邻接表
        for e in edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            let wt = e[2];
            g[x].push((y, wt));
            g[y].push((x, wt));
        }

        let mut dis = vec![-1; n];
        dis[0] = 0;
        let mut h = BinaryHeap::new(); // 下面取相反数，变成小根堆
        h.push((0, 0));
        while let Some((dx, x)) = h.pop() {
            if -dx > dis[x] { // x 之前出堆过
                continue;
            }
            for &(y, d) in &g[x] {
                let new_dis = -dx + d;
                if new_dis < disappear[y] && (dis[y] < 0 || new_dis < dis[y]) {
                    dis[y] = new_dis; // 更新 x 的邻居的最短路
                    h.push((-new_dis, y));
                }
            }
        }
        dis
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $m$ 为 $\textit{edges}$ 的长度。注意堆中会有重复节点，所以至多有 $\mathcal{O}(m)$ 个元素，单次操作的复杂度是 $\mathcal{O}(\log m)$，不是 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n+m)$。

更多相似题目，见下面的图论题单。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
