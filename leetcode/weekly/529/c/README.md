**注**：本题不是 Dijkstra 算法，请假设你没学过 Dijkstra，可能更好理解。

当节点 $x$ 在 $t$ 时刻激活时，我们在一个数据结构中记录：

- $x$ 的邻居 $y$ 在 $t + \textit{delay}$ 时刻收到了一个信号。其中 $\textit{delay}$ 是信号从 $x$ 到 $y$ 的时间。

只要我们能按照时间顺序依次处理每个事件，就能**模拟**这道题。

这要求数据结构支持：插入时刻、找到数据结构中的最小时刻、删除数据结构中的最小时刻。

适合的数据结构是**最小堆**。

一开始把时刻 $0$ 和节点 $0$ 入堆（堆中保存的是 pair）。把 $\textit{requiredSignals}[0]$ 改成 $1$。初始化 $\textit{ans}[i] = -1$。

当节点 $x$ 出堆时，先把 $\textit{requiredSignals}[x]$ 减少 $1$，如果 $\textit{requiredSignals}[x]$ 恰好等于 $0$，说明 $x$ 恰好在此刻激活，把当前时刻记录到 $\textit{ans}[x]$ 中，然后把 $x$ 的邻居 $y$ 的收到信号的时刻以及节点 $y$ 入堆。 

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def earliestActivationTimes(self, n: int, edges: list[list[int]], requiredSignals: list[int]) -> list[int]:
        g = [[] for _ in range(n)]
        for x, y, delay in edges:
            g[x].append((y, delay))

        ans = [-1] * n
        requiredSignals[0] = 1
        h = [(0, 0)]

        while h:
            t, x = heappop(h)
            requiredSignals[x] -= 1
            if requiredSignals[x] != 0:  # x 尚未激活或者已经激活
                continue
            ans[x] = t
            for y, delay in g[x]:
                heappush(h, (t + delay, y))

        return ans
```

```java [sol-Java]
class Solution {
    public long[] earliestActivationTimes(int n, int[][] edges, int[] requiredSignals) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }

        long[] ans = new long[n];
        Arrays.fill(ans, -1);
        requiredSignals[0] = 1;
        PriorityQueue<long[]> pq = new PriorityQueue<>(Comparator.comparingLong(p -> p[0]));
        pq.offer(new long[]{0, 0});

        while (!pq.isEmpty()) {
            long[] top = pq.poll();
            long t = top[0];
            int x = (int) top[1];
            requiredSignals[x]--;
            if (requiredSignals[x] != 0) { // x 尚未激活或者已经激活
                continue;
            }
            ans[x] = t;
            for (int[] e : g[x]) {
                pq.offer(new long[]{t + e[1], e[0]});
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> earliestActivationTimes(int n, vector<vector<int>>& edges, vector<int>& requiredSignals) {
        vector<vector<pair<int, int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2]);
        }

        vector<long long> ans(n, -1);
        requiredSignals[0] = 1;
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        pq.emplace(0, 0);

        while (!pq.empty()) {
            auto [t, x] = pq.top();
            pq.pop();
            requiredSignals[x]--;
            if (requiredSignals[x] != 0) { // x 尚未激活或者已经激活
                continue;
            }
            ans[x] = t;
            for (auto& [y, delay] : g[x]) {
                pq.emplace(t + delay, y);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func earliestActivationTimes(n int, edges [][]int, requiredSignals []int) []int64 {
	type edge struct{ to, delay int }
	g := make([][]edge, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], edge{e[1], e[2]})
	}

	ans := make([]int64, n)
	for i := range ans {
		ans[i] = -1
	}
	requiredSignals[0] = 1
	h := &hp{{}}

	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		t := top.t
		x := top.x
		requiredSignals[x]--
		if requiredSignals[x] != 0 { // x 尚未激活或者已经激活
			continue
		}
		ans[x] = t
		for _, e := range g[x] {
			heap.Push(h, pair{t + int64(e.delay), e.to})
		}
	}

	return ans
}

type pair struct {
	t int64
	x int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].t < h[j].t }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m)$，其中 $m$ 是 $\textit{edges}$ 的长度。最坏情况下，堆中有 $m$ 条数据，操作堆需要 $\mathcal{O}(\log m)$ 的时间。
- 空间复杂度：$\mathcal{O}(n + m)$。

## 专题训练

见下面数据结构题单的「**五、堆（优先队列）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
