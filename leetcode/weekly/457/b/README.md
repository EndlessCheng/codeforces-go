首先，建图 + DFS，把每个连通块中的节点加到各自的最小堆中。每个最小堆维护对应连通块的节点编号。

然后处理询问。

对于类型二，用一个 $\textit{offline}$ 布尔数组表示离线的电站。这一步不修改堆。

对于类型一：

- 如果电站 $x$ 在线，那么答案为 $x$。
- 否则检查 $x$ 所处堆的堆顶是否在线。若离线，则弹出堆顶，重复该过程。如果堆为不空，那么答案为堆顶，否则为 $-1$。

为了找到 $x$ 所属的堆，还需要一个数组 $\textit{belong}$ 记录每个节点在哪个堆中。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def processQueries(self, c: int, connections: List[List[int]], queries: List[List[int]]) -> List[int]:
        g = [[] for _ in range(c + 1)]
        for x, y in connections:
            g[x].append(y)
            g[y].append(x)

        belong = [-1] * (c + 1)
        heaps = []

        def dfs(x: int) -> None:
            belong[x] = len(heaps)  # 记录节点 x 在哪个堆
            h.append(x)
            for y in g[x]:
                if belong[y] < 0:
                    dfs(y)

        for i in range(1, c + 1):
            if belong[i] >= 0:
                continue
            h = []
            dfs(i)
            heapify(h)
            heaps.append(h)

        ans = []
        offline = [False] * (c + 1)
        for op, x in queries:
            if op == 2:
                offline[x] = True
                continue
            if not offline[x]:
                ans.append(x)
                continue
            h = heaps[belong[x]]
            # 懒删除：取堆顶的时候，如果离线，才删除
            while h and offline[h[0]]:
                heappop(h)
            ans.append(h[0] if h else -1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] processQueries(int c, int[][] connections, int[][] queries) {
        List<Integer>[] g = new ArrayList[c + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : connections) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        int[] belong = new int[c + 1];
        Arrays.fill(belong, -1);
        List<PriorityQueue<Integer>> heaps = new ArrayList<>();
        PriorityQueue<Integer> pq;
        for (int i = 1; i <= c; i++) {
            if (belong[i] >= 0) {
                continue;
            }
            pq = new PriorityQueue<>();
            dfs(i, g, belong, heaps.size(), pq);
            heaps.add(pq);
        }

        int ansSize = 0;
        for (int[] q : queries) {
            if (q[0] == 1) {
                ansSize++;
            }
        }

        int[] ans = new int[ansSize];
        int idx = 0;
        boolean[] offline = new boolean[c + 1];
        for (int[] q : queries) {
            int x = q[1];
            if (q[0] == 2) {
                offline[x] = true;
                continue;
            }
            if (!offline[x]) {
                ans[idx++] = x;
                continue;
            }
            pq = heaps.get(belong[x]);
            // 懒删除：取堆顶的时候，如果离线，才删除
            while (!pq.isEmpty() && offline[pq.peek()]) {
                pq.poll();
            }
            ans[idx++] = pq.isEmpty() ? -1 : pq.peek();
        }
        return ans;
    }

    private void dfs(int x, List<Integer>[] g, int[] belong, int compId, PriorityQueue<Integer> pq) {
        belong[x] = compId; // 记录节点 x 在哪个堆
        pq.offer(x);
        for (int y : g[x]) {
            if (belong[y] < 0) {
                dfs(y, g, belong, compId, pq);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> processQueries(int c, vector<vector<int>>& connections, vector<vector<int>>& queries) {
        vector<vector<int>> g(c + 1);
        for (auto& e : connections) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        vector<int> belong(c + 1, -1);
        vector<priority_queue<int, vector<int>, greater<>>> heaps;
        priority_queue<int, vector<int>, greater<>> pq;

        auto dfs = [&](this auto&& dfs, int x) -> void {
            belong[x] = heaps.size(); // 记录节点 x 在哪个堆
            pq.push(x);
            for (int y : g[x]) {
                if (belong[y] < 0) {
                    dfs(y);
                }
            }
        };

        for (int i = 1; i <= c; i++) {
            if (belong[i] < 0) {
                dfs(i);
                heaps.emplace_back(move(pq));
            }
        }

        vector<int> ans;
        vector<int8_t> offline(c + 1);
        for (auto& q : queries) {
            int x = q[1];
            if (q[0] == 2) {
                offline[x] = true;
                continue;
            }
            if (!offline[x]) {
                ans.push_back(x);
                continue;
            }
            auto& h = heaps[belong[x]];
            // 懒删除：取堆顶的时候，如果离线，才删除
            while (!h.empty() && offline[h.top()]) {
                h.pop();
            }
            ans.push_back(h.empty() ? -1 : h.top());
        }
        return ans;
    }
};
```

```go [sol-Go]
func processQueries(c int, connections [][]int, queries [][]int) (ans []int) {
	g := make([][]int, c+1)
	for _, e := range connections {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	belong := make([]int, c+1)
	for i := range belong {
		belong[i] = -1
	}
	heaps := []hp{}
	var h hp

	var dfs func(int)
	dfs = func(x int) {
		belong[x] = len(heaps) // 记录节点 x 在哪个堆
		h.IntSlice = append(h.IntSlice, x)
		for _, y := range g[x] {
			if belong[y] < 0 {
				dfs(y)
			}
		}
	}
	for i := 1; i <= c; i++ {
		if belong[i] >= 0 {
			continue
		}
		h = hp{}
		dfs(i)
		heap.Init(&h)
		heaps = append(heaps, h)
	}

	offline := make([]bool, c+1)
	for _, q := range queries {
		x := q[1]
		if q[0] == 2 {
			offline[x] = true
			continue
		}
		if !offline[x] {
			ans = append(ans, x)
			continue
		}
		// 懒删除：取堆顶的时候，如果离线，才删除
		h := &heaps[belong[x]]
		for h.Len() > 0 && offline[h.IntSlice[0]] {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			ans = append(ans, h.IntSlice[0])
		} else {
			ans = append(ans, -1)
		}
	}
	return
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(c\log c+n + q\log c)$ 或者 $\mathcal{O}(c+n + q\log c)$，取决于实现，其中 $n$ 是 $\textit{connections}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(c+n)$。返回值不计入。

## 专题训练

1. 图论题单的「**§1.1 DFS 基础**」。
2. 数据结构题单的「**§5.6 懒删除堆**」。

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
