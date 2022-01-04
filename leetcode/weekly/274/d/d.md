本文将介绍处理基环树问题的一些通用技巧。

从 $i$ 向 $\textit{favorite}[i]$ 连边，我们可以得到一张有向图。由于每个大小为 $k$ 的连通块都有 $k$ 个点和 $k$ 条边，所以每个连通块必定有且仅有一个环，且由于每个点的出度均为 $1$，这样的有向图又叫做内向**基环树 (pseudotree)**，由基环树组成的森林叫**基环树森林 (pseudoforest)**。

每一个内向基环树（连通块）都由一个**基环**和其余指向基环的**树枝**组成。例如示例 $3$ 可以得到如下内向基环树，其基环由节点 $0$、$1$、$3$ 和 $4$ 组成，节点 $2$ 为其树枝：

![1.png](https://pic.leetcode-cn.com/1641096462-IsWZUX-1.png)

特别地，我们得到的基环可能只包含两个节点。例如示例 $1$ 可以得到如下内向基环树，其基环只包含节点 $1$ 和 $2$，而节点 $0$ 和 $3$ 组成其树枝：

![2.png](https://pic.leetcode-cn.com/1641096467-KCwxMo-2.png)

对于本题来说，这两类基环树在组成圆桌时会有明显区别，下文会说明这一点。

先来看看基环大小大于 $2$ 的情况，显然基环上的节点组成了一个环，因而可以组成一个圆桌；而树枝上的点，若插入圆桌上 $v\rightarrow w$ 这两人中间，会导致节点 $v$ 无法和其喜欢的员工坐在一起，因此树枝上的点是无法插入圆桌的；此外，树枝上的点也不能单独组成圆桌，因为这样会存在一个出度为 $0$ 的节点，其无法和其喜欢的员工坐在一起。对于其余内向基环树（连通块）上的节点，和树枝同理，也无法插入该基环组成的圆桌。

因此，对于基环大小大于 $2$ 的情况，圆桌的最大员工数目即为最大的基环大小，记作 $\textit{maxRingSize}$。

下面来分析基环大小等于 $2$ 的情况。

以如下基环树为例，$0$ 和 $1$ 组成基环，其余节点组成树枝：

![3.png](https://pic.leetcode-cn.com/1641096473-JtGBgY-3.png)

我们可以先让 $0$ 和 $1$ 坐在圆桌旁（假设 $0$ 坐在 $1$ 左侧），那么 $0$ 这一侧的树枝只能坐在 $0$ 的左侧，而 $1$ 这一侧的树枝只能坐在 $1$ 的右侧。

$2$ 可以紧靠着坐在 $0$ 的左侧，而 $3$ 和 $4$ 只能选一个坐在 $2$ 的左侧（如果 $4$ 紧靠着坐在 $2$ 的左侧，那么 $3$ 是无法紧靠着坐在 $4$ 的左侧的，反之亦然）。

这意味着从 $0$ 出发倒着找树枝上的点（即沿着反图上的边），每个点只能在其反图上选择其中一个子节点，因此 $0$ 这一侧的节点必须组成一条链，那么我们可以找最长的那条链，即上图加粗的节点。

对于 $1$ 这一侧也同理，将这两条最长链拼起来即为该基环树能组成的圆桌的最大员工数。

对于多个基环大小等于 $2$ 的基环树，每个基环树所对应的链，都可以拼在其余链的末尾，因此我们可以将这些链全部拼成一个圆桌，其大小记作 $\textit{sumChainSize}$。

答案即为 $\max(\textit{maxRingSize},\textit{sumChainSize})$。

---

下面介绍基环树问题的通用写法。

我们可以通过一次拓扑排序「剪掉」所有树枝，这样就可以将基环和树枝分开，从而简化后续处理流程：

- 如果要遍历基环，可以从拓扑排序后入度大于 $0$ 的节点出发，在图上搜索；
- 如果要遍历树枝，可以以基环与树枝的连接处为起点，顺着反图来搜索树枝（树枝上的节点在拓扑排序后的入度均为 $0$），从而将问题转化成一个树形问题。

对于本题，我们可以遍历所有基环，并按基环大小分类计算：

- 对于大小大于 $2$ 的基环，我们取基环大小的最大值；
- 对于大小等于 $2$ 的基环，我们可以从基环上的点出发，在反图上找到最大的树枝节点深度。

时间复杂度和空间复杂度均为 $O(n)$。

```go [sol1-Go]
func maximumInvitations(g []int) int { // favorite 就是内向基环森林 g
	n := len(g)
	rg := make([][]int, n) // g 的反图
	deg := make([]int, n)  // g 上每个节点的入度
	for v, w := range g {
		rg[w] = append(rg[w], v)
		deg[w]++
	}

	// 拓扑排序，剪掉 g 上的所有树枝
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := g[v] // v 只有一条出边
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	// 通过反图 rg 寻找树枝上最深的链
	var rdfs func(int) int
	rdfs = func(v int) int {
		maxDepth := 1
		for _, w := range rg[v] {
			if deg[w] == 0 { // 树枝上的点在拓扑排序后，入度均为 0
				maxDepth = max(maxDepth, rdfs(w)+1)
			}
		}
		return maxDepth
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range deg {
		if d <= 0 {
			continue
		}
		// 遍历基环上的点（拓扑排序后入度大于 0）
		deg[i] = -1
		ringSize := 1
		for v := g[i]; v != i; v = g[v] {
			deg[v] = -1 // 将基环上的点的入度标记为 -1，避免重复访问
			ringSize++
		}
		if ringSize == 2 { // 基环大小为 2
			sumChainSize += rdfs(i) + rdfs(g[i]) // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
public:
    int maximumInvitations(vector<int> &g) { // favorite 就是内向基环森林 g
        int n = g.size();
        vector<vector<int>> rg(n); // g 的反图
        vector<int> deg(n); // g 上每个节点的入度
        for (int v = 0; v < n; ++v) {
            int w = g[v];
            rg[w].emplace_back(v);
            ++deg[w];
        }

        // 拓扑排序，剪掉 g 上的所有树枝
        queue<int> q;
        for (int i = 0; i < n; ++i) {
            if (deg[i] == 0) {
                q.emplace(i);
            }
        }
        while (!q.empty()) {
            int v = q.front();
            q.pop();
            int w = g[v]; // v 只有一条出边
            if (--deg[w] == 0) {
                q.emplace(w);
            }
        }

        // 通过反图 rg 寻找树枝上最深的链
        function<int(int)> rdfs = [&](int v) -> int {
            int max_depth = 1;
            for (int w: rg[v]) {
                if (deg[w] == 0) { // 树枝上的点在拓扑排序后，入度均为 0
                    max_depth = max(max_depth, rdfs(w) + 1);
                }
            }
            return max_depth;
        };

        int max_ring_size = 0, sum_chain_size = 0;
        for (int i = 0; i < n; ++i) {
            if (deg[i] <= 0) {
                continue;
            }
            // 遍历基环上的点（拓扑排序后入度大于 0）
            deg[i] = -1;
            int ring_size = 1;
            for (int v = g[i]; v != i; v = g[v]) {
                deg[v] = -1; // 将基环上的点的入度标记为 -1，避免重复访问
                ++ring_size;
            }
            if (ring_size == 2) { // 基环大小为 2
                sum_chain_size += rdfs(i) + rdfs(g[i]); // 累加两条最长链的长度
            } else {
                max_ring_size = max(max_ring_size, ring_size); // 取所有基环的最大值
            }
        }
        return max(max_ring_size, sum_chain_size);
    }
};
```

```python [sol1-Python3]
class Solution:
    def maximumInvitations(self, g: List[int]) -> int:  # favorite 就是内向基环森林 g
        n = len(g)
        rg = [[] for _ in range(n)]  # g 的反图
        deg = [0] * n  # g 上每个节点的入度
        for v, w in enumerate(g):
            rg[w].append(v)
            deg[w] += 1

        # 拓扑排序，剪掉 g 上的所有树枝
        q = deque(i for i, d in enumerate(deg) if d == 0)
        while q:
            v = q.popleft()
            w = g[v]  # v 只有一条出边
            deg[w] -= 1
            if deg[w] == 0:
                q.append(w)

        # 通过反图 rg 寻找树枝上最深的链
        def rdfs(v: int) -> int:
            max_depth = 1
            for w in rg[v]:
                if deg[w] == 0:  # 树枝上的点在拓扑排序后，入度均为 0
                    max_depth = max(max_depth, rdfs(w) + 1)
            return max_depth

        max_ring_size, sum_chain_size = 0, 0
        for i, d in enumerate(deg):
            if d <= 0:
                continue
            # 遍历基环上的点（拓扑排序后入度大于 0）
            deg[i] = -1
            ring_size = 1
            v = g[i]
            while v != i:
                deg[v] = -1  # 将基环上的点的入度标记为 -1，避免重复访问
                ring_size += 1
                v = g[v]
            if ring_size == 2:  # 基环大小为 2
                sum_chain_size += rdfs(i) + rdfs(g[i])  # 累加两条最长链的长度
            else:
                max_ring_size = max(max_ring_size, ring_size)  # 取所有基环的最大值
        return max(max_ring_size, sum_chain_size)
```

[@Class_](/u/class_/) 指出可以在拓扑排序的同时计算出最长链的长度，这样就不需要建反图和在反图上找最长链了，从而节省不少时间、空间和代码量：

```go [sol2-Go]
func maximumInvitations(g []int) int { // favorite 就是内向基环森林 g
	n := len(g)
	deg := make([]int, n) // g 上每个节点的入度
	for _, w := range g {
		deg[w]++
	}

	maxDepth := make([]int, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 拓扑排序，剪掉 g 上的所有树枝
		v := q[0]
		q = q[1:]
		maxDepth[v]++
		w := g[v] // v 只有一条出边
		maxDepth[w] = max(maxDepth[w], maxDepth[v])
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range deg {
		if d == 0 {
			continue
		}
		// 遍历基环上的点（拓扑排序后入度大于 0）
		deg[i] = 0
		ringSize := 1
		for v := g[i]; v != i; v = g[v] {
			deg[v] = 0 // 将基环上的点的入度标记为 0，避免重复访问
			ringSize++
		}
		if ringSize == 2 { // 基环大小为 2
			sumChainSize += maxDepth[i] + maxDepth[g[i]] + 2 // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol2-C++]
class Solution {
public:
    int maximumInvitations(vector<int> &g) { // favorite 就是内向基环森林 g
        int n = g.size();
        vector<int> deg(n); // g 上每个节点的入度
        for (int w: g) {
            ++deg[w];
        }

        vector<int> max_depth(n);
        queue<int> q;
        for (int i = 0; i < n; ++i) {
            if (deg[i] == 0) {
                q.emplace(i);
            }
        }
        while (!q.empty()) {  // 拓扑排序，剪掉 g 上的所有树枝
            int v = q.front();
            q.pop();
            ++max_depth[v];
            int w = g[v]; // v 只有一条出边
            max_depth[w] = max(max_depth[w], max_depth[v]);
            if (--deg[w] == 0) {
                q.emplace(w);
            }
        }

        int max_ring_size = 0, sum_chain_size = 0;
        for (int i = 0; i < n; ++i) {
            if (deg[i] == 0) {
                continue;
            }
            // 遍历基环上的点（拓扑排序后入度大于 0）
            deg[i] = 0;
            int ring_size = 1;
            for (int v = g[i]; v != i; v = g[v]) {
                deg[v] = 0; // 将基环上的点的入度标记为 0，避免重复访问
                ++ring_size;
            }
            if (ring_size == 2) { // 基环大小为 2
                sum_chain_size += max_depth[i] + max_depth[g[i]] + 2; // 累加两条最长链的长度
            } else {
                max_ring_size = max(max_ring_size, ring_size); // 取所有基环的最大值
            }
        }
        return max(max_ring_size, sum_chain_size);
    }
};
```

```Python [sol2-Python3]
class Solution:
    def maximumInvitations(self, g: List[int]) -> int:  # favorite 就是内向基环森林 g
        n = len(g)
        deg = [0] * n  # g 上每个节点的入度
        for w in g:
            deg[w] += 1

        max_depth = [0] * n
        q = deque(i for i, d in enumerate(deg) if d == 0)
        while q:  # 拓扑排序，剪掉 g 上的所有树枝
            v = q.popleft()
            max_depth[v] += 1
            w = g[v]  # v 只有一条出边
            max_depth[w] = max(max_depth[w], max_depth[v])
            deg[w] -= 1
            if deg[w] == 0:
                q.append(w)

        max_ring_size, sum_chain_size = 0, 0
        for i, d in enumerate(deg):
            if d == 0:
                continue
            # 遍历基环上的点（拓扑排序后入度大于 0）
            deg[i] = 0
            ring_size = 1
            v = g[i]
            while v != i:
                deg[v] = 0  # 将基环上的点的入度标记为 0，避免重复访问
                ring_size += 1
                v = g[v]
            if ring_size == 2:  # 基环大小为 2
                sum_chain_size += max_depth[i] + max_depth[g[i]] + 2  # 累加两条最长链的长度
            else:
                max_ring_size = max(max_ring_size, ring_size)  # 取所有基环的最大值
        return max(max_ring_size, sum_chain_size)
```
