## 方法一：懒删除堆

首先，建图 + DFS，把每个连通块中的节点加到各自的最小堆中。每个最小堆维护对应连通块的节点编号。

然后处理询问。

对于类型二，用一个 $\textit{offline}$ 布尔数组表示离线的电站。这一步不修改堆。

对于类型一：

- 如果电站 $x$ 在线，那么答案为 $x$。
- 否则检查 $x$ 所处堆的堆顶是否在线。若离线，则弹出堆顶，重复该过程。如果堆为不空，那么答案为堆顶，否则为 $-1$。

为了找到 $x$ 所属的堆，还需要一个数组 $\textit{belong}$ 记录每个节点在哪个堆中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1GF3qzMEni/?t=5m42s)，欢迎点赞关注~

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

## 方法二：倒序处理 + 维护最小值

倒序处理询问，离线变成在线，删除变成添加，每个连通块只需要一个 $\texttt{int}$ 变量就可以维护最小值。

注意可能存在同一个节点多次离线的情况，我们需要记录节点离线的最早时间（询问的下标）。对于倒序处理来说，离线的最早时间才是真正的在线时间。

```py [sol-Python3]
class Solution:
    def processQueries(self, c: int, connections: List[List[int]], queries: List[List[int]]) -> List[int]:
        g = [[] for _ in range(c + 1)]
        for x, y in connections:
            g[x].append(y)
            g[y].append(x)

        belong = [-1] * (c + 1)
        cc = 0  # 连通块编号

        def dfs(x: int) -> None:
            belong[x] = cc  # 记录节点 x 在哪个连通块
            for y in g[x]:
                if belong[y] < 0:
                    dfs(y)

        for i in range(1, c + 1):
            if belong[i] < 0:
                dfs(i)
                cc += 1

        # 记录每个节点的离线时间，初始为无穷大（始终在线）
        offline_time = [inf] * (c + 1)
        for i in range(len(queries) - 1, -1, -1):
            t, x = queries[i]
            if t == 2:
                offline_time[x] = i  # 记录离线时间

        # 每个连通块中仍在线的电站的最小编号
        mn = [inf] * cc
        for i in range(1, c + 1):
            if offline_time[i] == inf:  # 最终仍在线
                j = belong[i]
                mn[j] = min(mn[j], i)

        ans = []
        for i in range(len(queries) - 1, -1, -1):
            t, x = queries[i]
            j = belong[x]
            if t == 2:
                if offline_time[x] == i:
                    mn[j] = min(mn[j], x)  # 变回在线
            elif i < offline_time[x]:  # 已经在线（写 < 或者 <= 都可以）
                ans.append(x)
            elif mn[j] != inf:
                ans.append(mn[j])
            else:
                ans.append(-1)
        ans.reverse()
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
        int cc = 0;
        for (int i = 1; i <= c; i++) {
            if (belong[i] < 0) {
                dfs(i, g, belong, cc);
                cc++;
            }
        }

        int[] offlineTime = new int[c + 1];
        Arrays.fill(offlineTime, Integer.MAX_VALUE);
        int q1 = 0;
        for (int i = queries.length - 1; i >= 0; i--) {
            int[] q = queries[i];
            if (q[0] == 2) {
                offlineTime[q[1]] = i; // 记录最早离线时间
            } else {
                q1++;
            }
        }

        // 维护每个连通块的在线电站的最小编号
        int[] mn = new int[cc];
        Arrays.fill(mn, Integer.MAX_VALUE);
        for (int i = 1; i <= c; i++) {
            if (offlineTime[i] == Integer.MAX_VALUE) { // 最终仍然在线
                int j = belong[i];
                mn[j] = Math.min(mn[j], i);
            }
        }

        int[] ans = new int[q1];
        for (int i = queries.length - 1; i >= 0; i--) {
            int[] q = queries[i];
            int x = q[1];
            int j = belong[x];
            if (q[0] == 2) {
                if (offlineTime[x] == i) { // 变回在线
                    mn[j] = Math.min(mn[j], x);
                }
            } else {
                q1--;
                if (i < offlineTime[x]) { // 已经在线（写 < 或者 <= 都可以）
                    ans[q1] = x;
                } else if (mn[j] != Integer.MAX_VALUE) {
                    ans[q1] = mn[j];
                } else {
                    ans[q1] = -1;
                }
            }
        }
        return ans;
    }

    private void dfs(int x, List<Integer>[] g, int[] belong, int compId) {
        belong[x] = compId;
        for (int y : g[x]) {
            if (belong[y] < 0) {
                dfs(y, g, belong, compId);
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
        int cc = 0;
        auto dfs = [&](this auto&& dfs, int x) -> void {
            belong[x] = cc; // 记录节点 x 在哪个连通块
            for (int y : g[x]) {
                if (belong[y] < 0) {
                    dfs(y);
                }
            }
        };

        for (int i = 1; i <= c; i++) {
            if (belong[i] < 0) {
                dfs(i);
                cc++;
            }
        }

        vector<int> offline_time(c + 1, INT_MAX);
        for (int i = queries.size() - 1; i >= 0; i--) {
            auto& q = queries[i];
            if (q[0] == 2) {
                offline_time[q[1]] = i; // 记录最早离线时间
            }
        }

        // 维护每个连通块的在线电站的最小编号
        vector<int> mn(cc, INT_MAX);
        for (int i = 1; i <= c; i++) {
            if (offline_time[i] == INT_MAX) { // 最终仍然在线
                int j = belong[i];
                mn[j] = min(mn[j], i);
            }
        }

        vector<int> ans;
        for (int i = queries.size() - 1; i >= 0; i--) {
            auto& q = queries[i];
            int x = q[1];
            int j = belong[x];
            if (q[0] == 2) {
                if (offline_time[x] == i) { // 变回在线
                    mn[j] = min(mn[j], x);
                }
            } else if (i < offline_time[x]) { // 已经在线（写 < 或者 <= 都可以）
                ans.push_back(x);
            } else if (mn[j] != INT_MAX) {
                ans.push_back(mn[j]);
            } else {
                ans.push_back(-1);
            }
        }
        ranges::reverse(ans);
        return ans;
    }
};
```

```go [sol-Go]
func processQueries(c int, connections [][]int, queries [][]int) []int {
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
	cc := 0

	var dfs func(int)
	dfs = func(x int) {
		belong[x] = cc // 记录节点 x 在哪个连通块
		for _, y := range g[x] {
			if belong[y] < 0 {
				dfs(y)
			}
		}
	}
	for i := 1; i <= c; i++ {
		if belong[i] < 0 {
			dfs(i)
			cc++
		}
	}

	offlineTime := make([]int, c+1)
	for i := range offlineTime {
		offlineTime[i] = math.MaxInt
	}
	q1 := 0
	for i, q := range slices.Backward(queries) {
		if q[0] == 2 {
			offlineTime[q[1]] = i // 记录最早离线时间
		} else {
			q1++
		}
	}

	// 维护每个连通块的在线电站的最小编号
	mn := make([]int, cc)
	for i := range mn {
		mn[i] = math.MaxInt
	}
	for i := 1; i <= c; i++ {
		if offlineTime[i] == math.MaxInt { // 最终仍然在线
			j := belong[i]
			mn[j] = min(mn[j], i)
		}
	}

	ans := make([]int, q1)
	for i, q := range slices.Backward(queries) {
		x := q[1]
		j := belong[x]
		if q[0] == 2 {
			if offlineTime[x] == i { // 变回在线
				mn[j] = min(mn[j], x)
			}
		} else {
			q1--
			if i < offlineTime[x] { // 已经在线（写 < 或者 <= 都可以）
				ans[q1] = x
			} else if mn[j] != math.MaxInt {
				ans[q1] = mn[j]
			} else {
				ans[q1] = -1
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(c+n + q)$，其中 $n$ 是 $\textit{connections}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(c+n)$。返回值不计入。

## 专题训练

1. 图论题单的「**§1.1 DFS 基础**」。
2. 数据结构题单的「**§5.6 懒删除堆**」。
3. 数据结构题单的「**专题：离线算法**」。

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
