[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)

根据 Dijkstra 算法，同一个节点我们只会访问一次，所以「最多可使用一次开关」这个约束是多余的，我们只需把反向边的边权设置为 $2w_i$ 即可。答案为 $0$ 到 $n-1$ 的最短路长度。

```py [sol-Python3]
class Solution:
    def minCost(self, n: int, edges: List[List[int]]) -> int:
        g = [[] for _ in range(n)]  # 邻接表
        for x, y, wt in edges:
            g[x].append((y, wt))
            g[y].append((x, wt * 2))

        dis = [inf] * n
        dis[0] = 0  # 起点到自己的距离是 0
        h = [(0, 0)]  # 堆中保存 (起点到节点 x 的最短路长度，节点 x)

        while h:
            dis_x, x = heappop(h)
            if dis_x > dis[x]:  # x 之前出堆过
                continue
            if x == n - 1:  # 到达终点
                return dis_x
            for y, wt in g[x]:
                new_dis_y = dis_x + wt
                if new_dis_y < dis[y]:
                    dis[y] = new_dis_y  # 更新 x 的邻居的最短路
                    # 懒更新堆：只插入数据，不更新堆中数据
                    # 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                    heappush(h, (new_dis_y, y))

        return -1
```

```java [sol-Java]
class Solution {
    public int minCost(int n, int[][] edges) {
        List<int[]>[] g = new ArrayList[n]; // 邻接表
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            int wt = e[2];
            g[x].add(new int[]{y, wt});
            g[y].add(new int[]{x, wt * 2});
        }

        int[] dis = new int[n];
        Arrays.fill(dis, Integer.MAX_VALUE);
        // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(a -> a[0]));
        dis[0] = 0; // 起点到自己的距离是 0
        pq.offer(new int[]{0, 0});

        while (!pq.isEmpty()) {
            int[] p = pq.poll();
            int disX = p[0];
            int x = p[1];
            if (disX > dis[x]) { // x 之前出堆过
                continue;
            }
            if (x == n - 1) { // 到达终点
                return disX;
            }
            for (int[] e : g[x]) {
                int y = e[0];
                int wt = e[1];
                int newDisY = disX + wt;
                if (newDisY < dis[y]) {
                    dis[y] = newDisY; // 更新 x 的邻居的最短路
                    // 懒更新堆：只插入数据，不更新堆中数据
                    // 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
                    pq.offer(new int[]{newDisY, y});
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
    int minCost(int n, vector<vector<int>>& edges) {
        vector<vector<pair<int, int>>> g(n); // 邻接表
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].emplace_back(y, wt);
            g[y].emplace_back(x, wt * 2);
        }

        vector<int> dis(n, INT_MAX);
        // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        dis[0] = 0; // 起点到自己的距离是 0
        pq.emplace(0, 0);

        while (!pq.empty()) {
            auto [dis_x, x] = pq.top();
            pq.pop();
            if (dis_x > dis[x]) { // x 之前出堆过
                continue;
            }
            if (x == n - 1) { // 到达终点
                return dis_x;
            }
            for (auto& [y, wt] : g[x]) {
                auto new_dis_y = dis_x + wt;
                if (new_dis_y < dis[y]) {
                    dis[y] = new_dis_y; // 更新 x 的邻居的最短路
                    // 懒更新堆：只插入数据，不更新堆中数据
                    // 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                    pq.emplace(new_dis_y, y);
                }
            }
        }

        return -1;
    }
};
```

```go [sol-Go]
func minCost(n int, edges [][]int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 邻接表
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt * 2}) // 反转边
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0 // 起点到自己的距离是 0
	// 堆中保存 (起点到节点 x 的最短路长度，节点 x)
	h := &hp{{}}

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		disX, x := p.dis, p.x
		if disX > dis[x] { // x 之前出堆过
			continue
		}
		if x == n-1 { // 到达终点
			return disX
		}
		for _, e := range g[x] {
			y := e.to
			newDisY := disX + e.wt
			if newDisY < dis[y] {
				dis[y] = newDisY // 更新 x 的邻居的最短路
				// 懒更新堆：只插入数据，不更新堆中数据
				// 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
				heap.Push(h, pair{newDisY, y})
			}
		}
	}

	return -1
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m\log m)$，其中 $m$ 是 $\textit{edges}$ 的长度。
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
